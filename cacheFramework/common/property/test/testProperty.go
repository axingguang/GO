package main

import (
	"log"

	cfg "../util"
)

func main() {

	v := cfg.GetValue("redis", "redis.host")
	log.Fatal(v)
}
