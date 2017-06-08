package util

import (
	"time"

	"encoding/json"

	"strconv"
	"sync"

	logu "../../common/log/util"
	cfg "../../common/property/util"
	"../radix/pool"
	"../radix/redis"
)

var p *pool.Pool
var mu = &sync.Mutex{}

func getPool() *pool.Pool {
	mu.Lock()
	defer mu.Unlock()
	if p == nil {
		p = initPool()
	}

	return p
}

func initPool() *pool.Pool {

	host := cfg.GetValue("redis", "redis.host")
	port := cfg.GetValue("redis", "redis.port")
	poolNumstr := cfg.GetValue("redis", "redis.pool")

	poolnum, atoiErr := strconv.Atoi(poolNumstr)
	if atoiErr != nil {
		logu.Error("获取redis连接池数量 error:获取的redis连接池数量为：" + poolNumstr)
		logu.Error(atoiErr)
		return nil
	}
	idelTime := 10
	p, _ := pool.NewCustom("tcp", host+":"+port, poolnum, df)

	go func() {
		for {
			p.Cmd("PING")
			time.Sleep(time.Duration(idelTime) * time.Second)
		}
	}()
	return p
}

func df(network, addr string) (*redis.Client, error) {
	//配置文件
	ex := 100
	auth := cfg.GetValue("redis", "redis.pwd")
	dbstr := cfg.GetValue("redis", "redis.db")
	db, atoiErr := strconv.Atoi(dbstr)
	if atoiErr != nil {
		logu.Error("获取redis db error:获取的db为：" + dbstr)
		logu.Error(atoiErr)
		return nil, atoiErr
	}
	client, err := redis.DialTimeout(network, addr, time.Duration(ex)*time.Second)
	if err != nil {
		return nil, err
	}

	if err = client.Cmd("AUTH", auth).Err; err != nil {
		client.Close()
		logu.Error("link Redis error:")
		logu.Error(err)
		return nil, err
	}
	if err = client.Cmd("SELECT", db).Err; err != nil {
		client.Close()
		logu.Error("link Redis error:")
		logu.Error(err)
		return nil, err
	}
	return client, nil
}

func Set(key string, value interface{}, ex int) error {

	by, er := json.Marshal(value)
	if er != nil {
		logu.Error("序列化json失败 error:")
		logu.Error(er)
		return er
	}

	return SetBytes(key, by, ex)

}
func SetBytes(key string, val []byte, ex int) error {

	p := getPool()
	r := p.Cmd("set", key, val)
	if ex > 0 {
		SetexpireTime(key, ex)
	}
	return r.Err

}

func Get(key string, v interface{}) error {

	by, byErr := GetBytes(key)
	if byErr != nil {
		return byErr
	}
	return json.Unmarshal(by, v)
}
func GetBytes(key string) ([]byte, error) {
	p := getPool()
	r := p.Cmd("get", key)
	if r.Err != nil {
		return nil, r.Err

	}
	by, byErr := r.Bytes()
	if byErr != nil {
		return nil, byErr
	}
	if by == nil {
		return nil, nil
	}
	return by, nil
}
func SetexpireTime(key string, ex int) error {
	p := getPool()
	ret := p.Cmd("expire", key, ex)
	return ret.Err
}

func GetexpireTime(key string) (int, error) {
	p := getPool()
	ret := p.Cmd("ttl", key)
	return ret.Int()
}

func Delete(key string) bool {
	p := getPool()
	ret := p.Cmd("del", key)
	i, _ := ret.Int()
	if i <= 0 {
		return false
	}
	return true
}
