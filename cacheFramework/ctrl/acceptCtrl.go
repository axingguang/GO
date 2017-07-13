package ctrl

import (
	"net/http"

	"database/sql"
	"encoding/json"

	m "../dataModel"
	"../util"

	"strconv"

	db "../database"
)

/*

msgHeader={
	msgTypeNo:"1000",
	sysCode:"11111",
	msgId:10,
	msgStatus:1

}
msgBody={
		dataNo:"skuNo",
		dataType:"1",
		operateType:"1"
	}

*/

func SendMsg(w http.ResponseWriter, r *http.Request) {

	mHeader := r.FormValue("msgHeader")
	msgBody := r.FormValue("msgBody")
	if util.StrIsEmpty(mHeader) {
		w.Write([]byte("{status:0,msg:'msgHeader is nil'}"))
		return
	}
	if util.StrIsEmpty(msgBody) {
		w.Write([]byte("{status:0,msg:'msgBody is nil'}"))
		return
	}

	var headerData = &m.MsgHeader{}

	e := json.Unmarshal([]byte(mHeader), headerData)
	if e != nil {
		w.Write([]byte("msgHeader error"))
		return
	}
	var id int
	db.SaveData(func(t *sql.Tx) error {
		insertSql := "insert into msg_record(msg_type_no,msg_body,msg_process,msg_node,msg_status,sys_code,create_time) values(?,?,'0001_0','0001',0,?,now())"

		result, er := t.Exec(insertSql, headerData.MsgTypeNo, msgBody, headerData.SysCode)
		if er != nil {
			return er
		}

		idv, _ := result.LastInsertId()
		id = int(idv)
		return er
	})

	w.Write([]byte(strconv.Itoa(id)))
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
