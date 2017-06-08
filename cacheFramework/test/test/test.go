package main

import (
	"fmt"
	"time"
)

type Aa struct {
	a int
	b int
}

func main() {
	testChannel()
}
func testChannel() {
	ch := make(chan int, 10)
	ch <- 2
	go func() {
		for c := range ch {
			fmt.Println(c)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for ; ; i++ {
		ch <- i
	}

}

func testNew() {
	t1 := new(Aa)
	t1 = &Aa{1, 2}

	t3 := &Aa{7, 8}
	fmt.Println("1|a:--->", t1.a, "b:-->", t1.b)
	fmt.Println("2|a:--->", t3.a, "b:-->", t3.b)

	t2 := new(Aa)
	t2 = &Aa{4, 5}
	fmt.Println("3|a:--->", t2.a, "b:-->", t2.b)
	fmt.Println("4|a:--->", t1.a, "b:-->", t1.b)
}
