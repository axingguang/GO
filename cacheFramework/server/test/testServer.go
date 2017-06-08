package main

import (
	"net/http"

	_ "net/http/pprof"

	sht "../serverHttp"
)

func main() {

	go func() {
		http.ListenAndServe(":6060", nil)

	}()

	sht.SetPath("/aaaa", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("aaaa"))
	})

	sht.SetPath("/bbbb", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bbbb"))
	})

	sht.StartServer()
}
