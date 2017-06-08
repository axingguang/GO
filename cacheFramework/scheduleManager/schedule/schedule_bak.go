package schedule

/*
import (
	"time"
	//logu "../../common/log/util"
	"strconv"
	"strings"

	cfg "../../common/property/util"
	mt "../../monitorManager/monitor"
)

var aa = StartSchedule()

type zkmT struct {
	maxT int
	minT int
}

var zkmP = initZkm()

func initZkm() *zkmT {
	zkmMaxTStr := cfg.GetValue("schedule", "zk.monitor.ping.time")
	zkmMaxT := 60
	if !strings.EqualFold(zkmMaxTStr, "") {
		zkmMax, err := strconv.Atoi(zkmMaxTStr)
		if err != nil {
			zkmMaxT = zkmMax
		}
	}

	zkmMinTStr := cfg.GetValue("schedule", "zk.monitor.reconnect.time")
	zkmMinT := 10
	if !strings.EqualFold(zkmMinTStr, "") {
		zkmMin, err := strconv.Atoi(zkmMinTStr)
		if err != nil {
			zkmMinT = zkmMin
		}
	}
	return &zkmT{zkmMaxT, zkmMinT}
}

func StartSchedule() string {
	time.AfterFunc(time.Duration(zkmP.maxT)*time.Second, monitorZk)
	return ""
}

func monitorZk() {
	mt.MonitorZk()
	zkm := mt.Getzkm()
	runT := 60
	if zkm.Status == 1 {
		runT = zkmP.maxT
	} else {
		runT = zkmP.minT
	}
	time.AfterFunc(time.Duration(runT)*time.Second, monitorZk)
}
*/
