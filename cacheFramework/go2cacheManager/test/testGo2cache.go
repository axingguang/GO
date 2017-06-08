package main

import (
	"log"

	cache "../go2cache"
)

func main() {
	key := "ttt1"
	//val := "tttt"
	//ex := 1000
	//cache.Set(key, val, ex)
	var a interface{}
	cache.Get(key, &a)
	log.Fatal("1----->", a)

}
