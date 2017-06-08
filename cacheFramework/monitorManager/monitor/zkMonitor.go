package monitor

import (
	//"fmt"
	"strings"

	zku "../../zkManager/util"
)

type ZkM struct {
	Status int
}

var zkm = New()

func New() *ZkM {
	zkc := zku.GetZkClient()
	if zkc == nil {
		return &ZkM{Status: 0}
	} else {
		return &ZkM{Status: zkc.Status}
	}
}
func Getzkm() *ZkM {
	return zkm
}
func MonitorZk() {

	zkc := zku.GetZkClient()
	state := zkc.Conn.State().String()

	if strings.EqualFold(state, "StateUnknown") || strings.EqualFold(state, "StateDisconnected") {
		zkm.Status = 0
		zkc.Conn.Close()
		zku.Init()
	} else {
		zkm.Status = 1
	}
}
