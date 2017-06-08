package main

import (
	"fmt"
	"strconv"
	"time"

	"net/http"
	_ "net/http/pprof"

	"../util"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)

	}()

	key := "test"
	val := "sfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdfsdsfdfsdfsdfsdfsdfdfd"
	num := 100000
	starttimestamp := time.Now().Unix()
	charr := make([]chan int, num)
	for i := 0; i < num; i++ {
		ch := make(chan int)
		charr[i] = ch
		a := strconv.Itoa(i)
		key += a
		go func(key string, val string, ch chan int, i int) {
			util.Set(key, ch, 10)
			ch <- i
		}(key, val, ch, i)

	}
	for _, v := range charr {
		<-v
	}
	endtimestamp := time.Now().Unix()
	fmt.Println(endtimestamp - starttimestamp)
}
