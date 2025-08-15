package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wp *sync.WaitGroup) {
	defer wp.Done()
	fmt.Println(s)
}
func main() {
	var wp sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
	}
	wp.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wp)
	}
	wp.Wait()

	wp.Add(1)
	printSomething("Holardas", &wp)

}
