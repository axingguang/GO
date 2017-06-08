package util

import (
	"sync"

	log "../../log/logrus"
	cfg "../goconfig"
)

var C *cfg.ConfigFile
var mu = &sync.Mutex{}

func GetValue(section, key string) string {
	str, _ := GetConfig().GetValue(section, key)
	return str
}

func GetConfig() *cfg.ConfigFile {
	if C == nil {
		C = initConfig()
	}
	return C
}

func initConfig() *cfg.ConfigFile {
	mu.Lock()
	defer mu.Unlock()
	var e error
	if C == nil {
		C, e = cfg.LoadConfigFile("../../config/conf.ini", "../../config/init.ini")
		if e != nil {
			log.Error("找不到配置文件:", e)
		}

	}

	return C
}
