package main

import (
	//"fmt"
	//"time"
	"strconv"

	syn "../dataSyn"
)

func main() {
	i := 0
	for ; ; i++ {
		key := "aaa" + strconv.Itoa(i)
		val := key
		syn.SetData(&syn.Data{Key: key, Val: val, Ex: 0})
	}
}
