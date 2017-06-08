package main

import (
	"fmt"

	"../../cache"
)

func main() {
	t1()
}

func t1() {
	key := "test1"
	v := cache.Get(key)
	fmt.Println("ceshijieguo:", v)
}
