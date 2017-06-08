package util

import (
	"encoding/json"
	"errors"
	"runtime/debug"
	"strconv"

	logu "../../common/log/util"
	cfg "../../common/property/util"

	"sync"

	"../freecache"
)

var mu = &sync.Mutex{}
var C *freecache.Cache

func GetCache() *freecache.Cache {
	if C == nil {
		C = initCache()
	}
	return C
}

func initCache() *freecache.Cache {
	mu.Lock()
	defer mu.Unlock()
	if C == nil {
		//大小为g
		sizeStr := cfg.GetValue("cache", "cache.size")
		logu.Info("cfg.GetValue(cache, cache.size):----->", sizeStr)
		if sizeStr == "" {
			logu.Error("获取设置内存大小数据失败，cfg.GetValue(cache, cache.size):----->", sizeStr)

			sizeStr = "1"
		}

		size, er := strconv.Atoi(sizeStr)

		if er != nil {
			size = 1
			logu.Error("cfg.GetValue(schedule, cache.clean)定时清理缓存无效数据参数异常:使用默认值（1）")
		}
		g := 1024 * 1024 * 1024
		//读取配置文件
		C = freecache.NewCache(g * size)
		debug.SetGCPercent(20)
	}
	return C
}

func Set(key string, val interface{}, ex int) error {
	if val == nil {
		return errors.New("val is nil ")
	}
	by, byEr := json.Marshal(val)
	if byEr != nil {
		return byEr
	}

	return SetBytes(key, by, ex)
}

func SetBytes(key string, val []byte, ex int) error {
	if val == nil {
		return errors.New("val is nil ")
	}

	return GetCache().Set([]byte(key), val, ex)
}

func Get(key string, v interface{}) (err error) {
	by, er := GetBytes(key)
	if er != nil {
		return er
	}
	return json.Unmarshal(by, v)
}

func GetBytes(key string) ([]byte, error) {
	return GetCache().Get([]byte(key))
}

func Delete(key string) bool {
	return GetCache().Del([]byte(key))
}
