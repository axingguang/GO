package schedule

import (
	"time"
	//logu "../../common/log/util"
	"fmt"
	"strconv"
	"strings"

	cfg "../../common/property/util"
	mt "../../monitorManager/monitor"
)

var serverMonitorT = initSerMT()

type zkmT struct {
	maxT int
	minT int
}

var zkmP = initZkm()
var aa = StartSchedule()

func initSerMT() int {
	serverMonitorTStr := cfg.GetValue("schedule", "server.monitor.ping")
	serMT := 10
	if !strings.EqualFold(serverMonitorTStr, "") {
		serverMonitor, err := strconv.Atoi(serverMonitorTStr)
		if err == nil {
			serMT = serverMonitor
		}
	}
	return serMT
}

func initZkm() *zkmT {
	zkmMaxTStr := cfg.GetValue("schedule", "zk.monitor.ping.time")
	fmt.Println("zk.monitor.ping.time", zkmMaxTStr)
	zkmMaxT := 60
	if !strings.EqualFold(zkmMaxTStr, "") {
		zkmMax, err := strconv.Atoi(zkmMaxTStr)
		if err == nil {
			zkmMaxT = zkmMax
		}
	}

	zkmMinTStr := cfg.GetValue("schedule", "zk.monitor.reconnect.time")
	fmt.Println("zk.monitor.reconnect.time", zkmMinTStr)
	zkmMinT := 10
	if !strings.EqualFold(zkmMinTStr, "") {
		zkmMin, err := strconv.Atoi(zkmMinTStr)
		if err == nil {
			zkmMinT = zkmMin
		}
	}
	return &zkmT{zkmMaxT, zkmMinT}
}

func StartSchedule() string {
	time.AfterFunc(time.Duration(zkmP.maxT)*time.Second, monitorZk)
	time.AfterFunc(time.Duration(serverMonitorT)*time.Second, monitorServer)
	return ""
}

func monitorZk() {
	mt.MonitorZk()
	zkm := mt.Getzkm()
	runT := zkmP.maxT
	if zkm.Status == 1 {
		runT = zkmP.maxT
	} else {
		runT = zkmP.minT
	}
	time.AfterFunc(time.Duration(runT)*time.Second, monitorZk)
}

func monitorServer() {
	mt.MonitorServer()
	time.AfterFunc(time.Duration(serverMonitorT)*time.Second, monitorServer)
}
