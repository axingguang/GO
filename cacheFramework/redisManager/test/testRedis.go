package main

import (
	"fmt"

	redisUtil "../util"
)

func main() {
	setEx()

}
func setEx() {

	//redisUtil.Set("ttt1", "tttttt", 0)

	var v interface{}
	e := redisUtil.Get("ttt1", &v)
	fmt.Println("1-----", v)
	fmt.Println("2-----", e)
}
