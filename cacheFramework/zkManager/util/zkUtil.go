package util

import (
	"sync"

	"errors"

	"net"
	"strconv"
	"strings"
	"time"

	logu "../../common/log/util"
	cfg "../../common/property/util"
	"../zk"
)

const (
	ZkErr = iota
	ZkSuc
)

type ZkC struct {
	Zkhost    []string
	IpNode    map[string]string
	Conn      *zk.Conn
	mu        *sync.RWMutex
	LocalAddr string
	Status    int
}

var (
	zkU = New()
)

func Init() {
	New()
}

func GetZkClient() *ZkC {
	return zkU
}

func getRootNote() string {
	namespace := cfg.GetValue("zk", "zk.namespace")
	if strings.EqualFold(namespace, "") {
		logu.Error("读取配置zk.namespace失败，zk.namespace:", namespace)
		return ""
	}

	projectname := cfg.GetValue("zk", "zk.projectname")
	if strings.EqualFold(projectname, "") {
		logu.Error("读取配置zk.projectname失败，zk.projectname:", projectname)
		return ""
	}

	rootNote := "/" + namespace + "/" + projectname
	return rootNote
}
func New() *ZkC {
	//配置文件
	zkstr := cfg.GetValue("zk", "zk.host")
	if strings.EqualFold(zkstr, "") {
		logu.Error("读取配置zk.host失败，zk.host:", zkstr)
		return nil
	}
	zkList := strings.Split(zkstr, ",")
	host, _ := getLoacalAddrs(zkList)
	port := cfg.GetValue("server", "server.port")
	if strings.EqualFold(port, "") {
		logu.Error("读取配置server.port失败，server.port:", port)
		return nil
	}
	timestr := cfg.GetValue("server", "zk.timeout")
	timeout := 10
	if !strings.EqualFold(timestr, "") {
		timeoutt, atoiErr := strconv.Atoi(timestr)
		if atoiErr != nil {
			timeout = 10
		} else {
			timeout = timeoutt
		}
	}

	conn, _, err := zk.Connect(zkList, time.Duration(timeout)*time.Second)

	if err != nil {
		logu.Error("连接zk失败:", zkList)
		return nil
	}
	//配置文件

	namespace := cfg.GetValue("zk", "zk.namespace")
	if strings.EqualFold(namespace, "") {
		logu.Error("读取配置zk.namespace失败，zk.namespace:", namespace)
		return nil
	}

	projectname := cfg.GetValue("zk", "zk.projectname")
	if strings.EqualFold(projectname, "") {
		logu.Error("读取配置zk.projectname失败，zk.projectname:", projectname)
		return nil
	}

	rootNote := "/" + namespace
	exist, _, _ := conn.Exists(rootNote)
	if !exist {
		_, errc := conn.Create(rootNote, nil, 0, zk.WorldACL(zk.PermAll))
		if errc != nil {
			logu.Error("创建节点失败，node:", rootNote, "err:", errc)
			return nil
		}
		logu.Info("创建节点成功，node:", rootNote)
	}

	projectNode := rootNote + "/" + projectname
	existp, _, _ := conn.Exists(projectNode)
	if !existp {
		_, errc := conn.Create(projectNode, nil, 0, zk.WorldACL(zk.PermAll))
		if errc != nil {
			logu.Error("创建节点失败，node:", projectNode, "err:", errc)
			return nil
		}
		logu.Info("创建节点成功，node:", projectNode)
	}

	localNote := host + ":" + port

	loacalFullNote := projectNode + "/" + localNote

	existl, _, _ := conn.Exists(loacalFullNote)
	if !existl {
		_, errc := conn.Create(loacalFullNote, nil, 0, zk.WorldACL(zk.PermAll))
		if errc != nil {
			logu.Error("创建节点失败，node:", loacalFullNote, "err:", errc)
			return nil
		}
		logu.Info("创建节点成功，node:", loacalFullNote)
	}

	loacalFullNoteStatus := loacalFullNote + "/status"
	loacalFullNoteConf := loacalFullNote + "/config"

	existNStatus, _, _ := conn.Exists(loacalFullNoteStatus)
	if !existNStatus {
		_, errc := conn.Create(loacalFullNoteStatus, []byte("ready"), 0, zk.WorldACL(zk.PermAll))
		if errc != nil {
			logu.Error("创建节点失败，node：", loacalFullNoteStatus, errc)
			return nil
		}
		logu.Info("创建节点成功，node:", loacalFullNoteStatus)
	} else {
		conn.Set(loacalFullNoteStatus, []byte("ready"), -1)
	}
	existNConf, _, _ := conn.Exists(loacalFullNoteConf)
	if !existNConf {
		_, errc := conn.Create(loacalFullNoteConf, nil, 0, zk.WorldACL(zk.PermAll))
		if errc != nil {
			logu.Error("创建节点失败，node：", loacalFullNoteConf, errc)
		}
		logu.Info("创建节点成功，node:", loacalFullNoteConf)
	}

	ipNode := getIpNodes(conn, projectNode)
	logu.Info("当前节点，nodes:", ipNode)
	ipNode[localNote] = "ready"
	aa := &ZkC{
		Zkhost:    zkList,
		IpNode:    ipNode,
		Conn:      conn,
		mu:        &sync.RWMutex{},
		LocalAddr: localNote,
		Status:    1,
	}

	wmap := make(map[string]string)

	zkNodeDataW(aa, conn, projectNode, wmap)

	go func(zkc *ZkC, con *zk.Conn, path string) {
		for {
			children, stat, ch, err := con.ChildrenW(path)
			e := <-ch
			logu.Info("节点发生变更，stat:", stat, "-->ch:", e, "-->err:", err)
			if len(children) > 0 {
				zkc.mu.RLock()
				zkc.IpNode = getIpNodes(con, path)
				logu.Info("最新节点输出，nodes:", zkc.IpNode)
				zkc.mu.RUnlock()
			}
			zkNodeDataW(aa, conn, projectNode, wmap)
		}

	}(aa, conn, projectNode)

	return aa
}

func zkNodeDataW(zkc *ZkC, con *zk.Conn, path string, wmap map[string]string) {
	ipNodes := zkc.IpNode
	for k, _ := range ipNodes {
		if _, ok := wmap[k]; !ok {
			wmap[k] = "1"
			go func(ipN string) {
				for {

					bt, Stat, ch, err := con.GetW(path + "/" + ipN + "/status")
					e := <-ch
					if err != nil {

						logu.Error("节点数据变化异常，nodes:", ipN, "result:", string(bt), Stat, e, err)
					}
					bt2, _, _ := con.Get(path + "/" + ipN + "/status")
					logu.Info("节点数据更新，nodes:", ipN, "由(", string(bt), ")变为(", string(bt2), ")")
					ipNodes[ipN] = string(bt2)
				}
			}(k)

		}
	}
	//pathData := con.GetW(path)
}

func getIpNodes(conn *zk.Conn, node string) map[string]string {
	ipNodes := make(map[string]string)
	children, _, err := conn.Children(node)
	if err != nil {
		logu.Error("获取子节点失败，当前节点为：", node)
		return nil
	}
	for _, v := range children {
		strt := node + "/" + v + "/status"
		bt, _, _ := conn.Get(strt)
		ipNodes[v] = string(bt)
	}

	return ipNodes

}

func getLoacalAddrs(zkL []string) (retStr string, retErr error) {
	var conn net.Conn
	var err error
	retErr = errors.New("获取本地ip失败")
	for _, str := range zkL {
		conn = nil
		err = nil
		conn, err = net.Dial("udp", str)
		if err != nil {
			logu.Error("zk服务器网络不通，host为：", str)
			continue
		}
		retStr = strings.Split(conn.LocalAddr().String(), ":")[0]
		retErr = nil
		break
	}

	defer conn.Close()
	return
}

func SetNote(key string, val string) {
	zkU.mu.RLock()
	defer zkU.mu.RUnlock()
	zkU.IpNode[key] = val
}
func SetNoteVal(host string, val string) {
	zkU.mu.RLock()
	defer zkU.mu.RUnlock()
	loacalFullNoteStatus := getRootNote() + "/" + host + "/status"
	s, e := zkU.Conn.Set(loacalFullNoteStatus, []byte(val), -1)
	logu.Info("设置节点状态，node:", loacalFullNoteStatus, "stat:", s, "错误信息：", e)
}

func GetNoteAll() []string {
	//var hostList []string
	var str string = ""
	for k, v := range zkU.IpNode {
		if strings.EqualFold(v, "ready") {
			str += k + ","
		}
	}
	str = strings.Trim(str, ",")
	if strings.EqualFold(str, "") {
		return nil
	}

	strar := strings.Split(str, ",")
	return strar
}
func GetNoteExLocal() []string {
	//var hostList []string
	var str string = ""
	for k, v := range zkU.IpNode {
		if strings.EqualFold(v, "ready") && !strings.EqualFold(zkU.LocalAddr, k) {
			str += k + ","

		}
	}
	str = strings.Trim(str, ",")
	if strings.EqualFold(str, "") {
		return nil
	}

	strar := strings.Split(str, ",")
	return strar
}
