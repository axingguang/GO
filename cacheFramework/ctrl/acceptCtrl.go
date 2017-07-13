package ctrl

import (
	"net/http"
)

func SendMsg(w http.ResponseWriter, r *http.Request) {

	
	w.Write([]byte("SendMsg")
}

func ReceiveMsg(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("ReceiveMsg"))
}

func RefrashMsgTypeAndSysRef(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("RefrashMsgTypeAndSysRef"))
}

func RefrashMsgTypeAndNodesRef(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("RefrashMsgTypeAndNodesRef"))
}
