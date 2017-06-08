package serverHttp

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	//"time"

	cacheUtil "../../cacheManager/util"
	logUtil "../../common/log/util"
	cfg "../../common/property/util"
	syn "../../dataSynManager/dataSyn"
	cache "../../go2cacheManager/go2cache"
	_ "../../scheduleManager/schedule"
)

var (
	MyHander = New()
)

type pathFunc func(http.ResponseWriter, *http.Request)

func StartServer() {
	ncache()
	wcache()
	syncache()
	for k, v := range MyHander.Path {
		http.HandleFunc(k, v)
	}

	//	server.port=8888
	//server.readTimeout=10
	//server.writeTimeout=10
	readTimeOutStr := cfg.GetValue("server", "server.readTimeout")
	readTimeOut := 10
	if !strings.EqualFold(readTimeOutStr, "") {
		a, er := strconv.Atoi(readTimeOutStr)
		if er == nil {
			readTimeOut = a
		}
	}

	writeTimeOutStr := cfg.GetValue("server", "server.writeTimeout")
	writeTimeOut := 10
	if !strings.EqualFold(writeTimeOutStr, "") {
		a, er := strconv.Atoi(writeTimeOutStr)
		if er == nil {
			writeTimeOut = a
		}
	}
	port := ":" + cfg.GetValue("server", "server.port")
	if strings.EqualFold(port, ":") {
		logUtil.Error("未设置端口号")
		return
	}
	log.Println(readTimeOut, writeTimeOut)
	//修改配置
	s := &http.Server{
		Addr: port,
		//Handler:      MyHander,
		//ReadTimeout:  time.Duration(readTimeOut) * time.Second,
		//WriteTimeout: time.Duration(writeTimeOut) * time.Second,
		//IdleTimeout:  time.Duration(readTimeOut) * 10 * time.Second,
		//	MaxHeaderBytes: 1 << 20,
	}

	log.Fatal("aaaa------->", s.ListenAndServe())
}

type Myh struct {
	Path map[string]pathFunc
}

func (myh *Myh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path

	aa := myh.Path[uri]
	if aa != nil {
		aa(w, r)
	} else {
		w.Write([]byte("404"))
	}

	/*

		key := r.FormValue("key")
		val := r.FormValue("value")
		duS := r.FormValue("ex")
		ds, _ := strconv.Atoi(duS)
		switch {
		case uri == "/":
		case uri == "/setValue":
			w.Write([]byte(setValue(key, val, ds)))
		case uri == "/getValue":
			w.Write(getValue(key))
		}

	*/

}

/*
func setValue(key string, val interface{}, du int) string {
	var dui int

	fmt.Println(dui)
	t := du
	cacheU.Set(key, val, t)

	return "success"
}

func getValue(key string) []byte {
	val, _ := cacheU.GetBytes(key)
	return val
}
*/

func New() *Myh {
	m := make(map[string]pathFunc)

	return &Myh{Path: m}
}
func GetPath() map[string]pathFunc {

	return MyHander.Path
}
func SetPath(path string, fun pathFunc) {
	MyHander.Path[path] = fun
}

func wcache() {

	MyHander.Path["/cache/setValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		val := r.FormValue("value")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}
		if strings.EqualFold(val, "") {
			w.Write([]byte("{status:0,msg:'val is nil'}"))
			return
		}
		exstr := r.FormValue("ex")
		ex := 0
		if !strings.EqualFold(exstr, "") {
			a, er := strconv.Atoi(exstr)
			if er != nil {
				ex = 0
			} else {
				ex = a
			}
		}
		err := cache.Set(key, val, ex)
		if err != nil {
			logUtil.Error("设置缓存失败，key：", key, "value:", val, "ex:", ex, "error:", err)
			w.Write([]byte("{status:0,msg:'set cache error'}"))
			return
		}
		w.Write([]byte("{status:1,msg:'success'}"))
	}

	MyHander.Path["/cache/getValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}
		var reta string
		err := cache.Get(key, &reta)
		if err != nil {
			logUtil.Error("读取缓存失败，key：", key, "error:", err)
			w.Write([]byte("{status:0,msg:'get cache error'}"))
			return
		}
		ret := "{status:1,msg:'success',data:" + reta + "}"
		w.Write([]byte(ret))
	}

	MyHander.Path["/cache/delValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}

		cache.Delete(key)
		w.Write([]byte("{status:1,msg:'success'}"))
	}

}

func ncache() {
	MyHander.Path["/ping"] = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}

	MyHander.Path["/cachen/setValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		val := r.FormValue("value")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}
		if strings.EqualFold(val, "") {
			w.Write([]byte("{status:0,msg:'val is nil'}"))
			return
		}
		exstr := r.FormValue("ex")
		ex := 0
		if !strings.EqualFold(exstr, "") {
			a, er := strconv.Atoi(exstr)
			if er != nil {
				ex = 0
			} else {
				ex = a
			}
		}
		err := cacheUtil.Set(key, val, ex)
		if err != nil {
			logUtil.Error("设置本地缓存失败，key：", key, "value:", val, "ex:", ex, "error:", err)
			w.Write([]byte("{status:0,msg:'set cache error'}"))
			return
		}
		w.Write([]byte("{status:1,msg:'success'}"))
	}

	MyHander.Path["/cachen/getValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}
		var reta string
		err := cacheUtil.Get(key, &reta)
		if err != nil {
			logUtil.Error("读取缓存失败，key：", key, "error:", err)
			w.Write([]byte("{status:0,msg:'get cache error'}"))
			return
		}
		ret := "{status:1,msg:'success',data:" + reta + "}"
		w.Write([]byte(ret))
	}

	MyHander.Path["/cachen/delValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}

		cacheUtil.Delete(key)
		w.Write([]byte("{status:1,msg:'success'}"))
	}
}

func syncache() {

	MyHander.Path["/syn/setValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		val := r.FormValue("value")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}
		if strings.EqualFold(val, "") {
			w.Write([]byte("{status:0,msg:'val is nil'}"))
			return
		}
		exstr := r.FormValue("ex")
		ex := 0
		if !strings.EqualFold(exstr, "") {
			a, er := strconv.Atoi(exstr)
			if er != nil {
				ex = 0
			} else {
				ex = a
			}
		}
		syn.SetData(&syn.Data{Key: key, Val: val, Ex: ex, DataType: syn.SetType})
	}

	MyHander.Path["/syn/delValue"] = func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if strings.EqualFold(key, "") {
			w.Write([]byte("{status:0,msg:'key is nil'}"))
			return
		}

		syn.SetData(&syn.Data{Key: key, DataType: syn.DelType})
	}
}
