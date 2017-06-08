package main

/**
客户端doc地址：github.com/samuel/go-zookeeper/zk
10.154.82.106:2181,10.154.82.107:2181,10.154.82.108:2181
**/
import (
	"fmt"

	"time"

	//zkutil "../util"
	"../zk"
)

/**
 * 获取一个zk连接
 * @return {[type]}
 */
func getConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

/**
 * 测试连接
 * @return
 */
func test1() {
	zkList := []string{"10.154.82.106:2181", "10.154.82.107:2181", "10.154.82.108:2181"}
	conn := getConnect(zkList)

	defer conn.Close()

	co, err := conn.Create("/go_serversyyyy", nil, 0, zk.WorldACL(zk.PermAll))
	fmt.Printf("------->co %v \n", co)
	fmt.Printf("------->err %v \n", err)

	//cc, ercc := conn.Create("/go_servers/aaa1", []byte("nodeTest1"), 0, zk.WorldACL(zk.PermAll))
	//fmt.Printf("------->cc %v \n", cc)
	//fmt.Printf("------->ercc %v \n", ercc)
	//time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 * @return {[type]}
 */
func test2() {
	zkList := []string{"10.154.82.106:2181", "10.154.82.107:2181", "10.154.82.108:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	conn.Create("/testadaadsasdsaw", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

/**
 * 获取所有节点
 */
func test3() {
	//zkList := []string{"zk01.n.lemall.com:2181", "zk02.n.lemall.com:2181", "zk03.n.lemall.com:2181",
	//	"zk04.n.lemall.com:2181", "zk05.n.lemall.com:2181"}
	zkList := []string{"10.154.82.106:2181", "10.154.82.107:2181", "10.154.82.108:2181"}
	//zkList := []string{"10.154.80.205:2181", "10.154.80.211:2181", "10.154.80.211:2181"}
	conn := getConnect(zkList)

	defer conn.Close()

	part := "/cacheFramework/cacheFrame/10.180.219.123:8881"

	children, _, err := conn.Children(part)
	if err != nil {
		fmt.Println("-------->get ccc%v \n", err)
	}

	for _, v := range children {
		strt := part + "/" + v
		bt, _, _ := conn.Get(strt)
		fmt.Printf("-------->children data"+strt+"---> %v \n", string(bt))
		//fmt.Printf("-------->children erbt %v \n", erbt)
	}

	fmt.Printf("children%v \n", children)
}

func setval() {
	//zkutil.SetNoteVal("10.87.14.216:8888", "fail")
}
func deletNode() {
	// /cacheFramework/cacheFrame/10.87.14.216:8880/status
	// /cacheFramework/cacheFrame/10.87.14.216:8880/config

	zkList := []string{"10.154.82.106:2181", "10.154.82.107:2181", "10.154.82.108:2181"}
	//zkList := []string{"10.154.80.205:2181", "10.154.80.211:2181", "10.154.80.211:2181"}
	conn := getConnect(zkList)

	defer conn.Close()
	e := conn.Delete("/cacheFramework", -1)
	//e2 := conn.Delete("/cacheFramework/cacheFrame/10.87.14.216:8882/status", -1)
	//e3 := conn.Delete("/cacheFramework/cacheFrame/10.87.14.216:8880", -1)
	fmt.Println("sdfsdfsdfsdf", e)
}

func main() {
	//test1()
	//test2()
	//setval()
	test3()
	//deletNode()

	//fmt.Println(zkutil.GetNoteAll())
}
