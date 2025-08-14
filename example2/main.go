package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string) {
	defer wg.Done()
	msg = fmt.Sprintf("Hello, %s", s)
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMsg("universe")
	go updateMsg("cosmos")

	wg.Wait()

	fmt.Println(msg)
}
