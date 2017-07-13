package router

import (
	"../ctrl"
	sht "../server/serverHttp"
)

func init() {
	sht.SetPath("/sendMsg", ctrl.SendMsg)
	sht.SetPath("/ReceiveMsg", ctrl.ReceiveMsg)
	sht.SetPath("/RefrashMsgTypeAndSysRef", ctrl.RefrashMsgTypeAndSysRef)
	sht.SetPath("/ReceiveMsg", ctrl.RefrashMsgTypeAndNodesRef)
}
