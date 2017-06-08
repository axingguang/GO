package go2cache

import (
	"encoding/json"
	"errors"

	cacheUtil "../../cacheManager/util"
	logUtil "../../common/log/util"
	redisUtil "../../redisManager/util"
)

func GetBytes(key string) (by []byte, reErr error) {
	by, err := cacheUtil.GetBytes(key)

	if by == nil || err != nil {
		by, reErr = redisUtil.GetBytes(key)
		if reErr != nil {
			logUtil.Error("根据key从redis中获取数据失败，key：", key, reErr)
			return
		}
		ex, exEr := redisUtil.GetexpireTime(key)

		if by != nil && err == nil {
			if ex > 0 && exEr == nil {
				cacheUtil.SetBytes(key, by, ex)
			} else {
				cacheUtil.SetBytes(key, by, 0)
			}
		}
	}
	return

}

func Get(key string, v interface{}) error {
	by, err := GetBytes(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(by, v)

}

func SetBytes(key string, val []byte, ex int) error {
	err := cacheUtil.SetBytes(key, val, ex)
	if err != nil {
		logUtil.Error("写入内存缓存失败，key：", key, err)
		return err
	}
	redisErr := redisUtil.SetBytes(key, val, ex)
	if redisErr != nil {
		logUtil.Error("写入redis失败，key：", key, redisErr)
	}
	return redisErr
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

func Delete(key string) {
	cacheUtil.Delete(key)
	redisUtil.Delete(key)
}
