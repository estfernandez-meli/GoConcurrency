package main

import (
	"fmt"
	"time"
)

func run(ch chan int) {
	for {
		s := <-ch
		fmt.Println("Got", s, "from channel")
		time.Sleep(time.Second * 1)
	}
}
func main() {
	ch := make(chan int, 50)

	go run(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("sending", i, "to channel")
		ch <- i
	}
}
