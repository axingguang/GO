package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"

	_ "../schedule"
)

func main() {

	go func() {
		http.ListenAndServe(":6060", nil)

	}()
	fmt.Println("100000")
	time.Sleep(3 * time.Minute)
}
