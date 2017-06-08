package monitor

import (
	"strings"

	"../../common/httpclient"
	logUtil "../../common/log/util"
	zku "../../zkManager/util"
)

var client = httpclient.GetHttpclient()

func MonitorServer() {
	zkc := zku.GetZkClient()
	ipNode := zkc.IpNode
	logUtil.Info("监控节点状态：nodes：", ipNode)
	if ipNode != nil {
		for k, v := range ipNode {
			go ping(k, v, zkc)
		}
	}
}

func ping(host, status string, zkc *zku.ZkC) {

	if strings.EqualFold(status, "ready") {
		res, err := client.Get("http://" + host + "/ping")
		logUtil.Info("ping", "http://"+host+"/ping", "|result:", res)
		if err != nil || res.StatusCode != 200 {
			logUtil.Error("ping", "http://"+host+"/ping", "|err:", err)
			zku.SetNote(host, "fail")
			zku.SetNoteVal(host, "fail")
		}
	}
	/*
		if strings.EqualFold(status, "ready") {
			if err != nil || res.StatusCode != 200 {
				zku.SetNote(host, "fail")
				zku.SetNoteVal(host, "fail")
			}
		} else {
			if err == nil && res.StatusCode == 200 {
				zku.SetNote(host, "ready")
				zku.SetNoteVal(host, "ready")
			}
		}*/

}
