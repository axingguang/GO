package monitor

/*
import (
	"time"

	"strings"

	"strconv"

	"fmt"

	logu "../../common/log/util"
	cfg "../../common/property/util"
	zku "../../zkManager/util"
	"../../zkManager/zk"
)

type ZkM struct {
	Status int
}

var zkm = New()

func Getzkm() *ZkM {
	return &zkm
}
func New() *ZkM {
	zkc := zku.GetZkClient()
	if zkc == nil {
		return &ZkM{Status: 0}
	} else {
		return &ZkM{Status: zkc.Status}
	}
}

func MonitorZk() {
	//配置文件
	zkstr := cfg.GetValue("zk", "zk.host")

	//zkstr := "10.154.82.106:218,10.154.82.107:218,10.154.82.108:218"
	if strings.EqualFold(zkstr, "") {
		logu.Error("读取配置zk.host失败，zk.host:", zkstr)
	}
	zkList := strings.Split(zkstr, ",")
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
	zkc := zku.GetZkClient()
	fmt.Println("zkstatus------------->", conn.State())
	fmt.Println("zkstatusaaaaaaaaaanil------------->", err)
	if err == nil {
		conn.Close()
		zkm.Status = zku.ZkSuc
		zkc.Status = zku.ZkSuc
		logu.Info("zk连接成功")
	} else {
		zkm.Status = zku.ZkErr
		zkc.Status = zku.ZkErr
		logu.Error("zk连接失败，zk.host:", zkstr)
	}
}
*/
