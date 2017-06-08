package main

import (
	"time"

	logUtil "../util"
)

func main() {
	for i := 0; i < 1000; i++ {
		logUtil.Info("1111111")
		logUtil.Warn("22222")
		logUtil.Error("33333")
	}
	time.Sleep(60 * time.Second)

}
