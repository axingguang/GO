package main

import (
	"fmt"

	"../httpclient"
)

func main() {
	testclicent()
}

func testclicent() {
	client := httpclient.GetHttpclient()

	res, err := client.Get("http://10.87.14.216:8880/ping")
	fmt.Print(res, err)
}
