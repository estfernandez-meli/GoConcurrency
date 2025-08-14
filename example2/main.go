package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = fmt.Sprintf("Hello, %s", s)
	m.Unlock()
}

func main() {
	msg = "Hello, world!"
	var mutex sync.Mutex
	wg.Add(2)
	go updateMsg("universe", &mutex)
	go updateMsg("cosmos", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
