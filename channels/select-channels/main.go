package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func server1(ch chan string) {
	for {
		fmt.Println("Starting server 1 ...")
		time.Sleep(time.Second * 6)
		ch <- "Server 1 started!!"
	}
}

func server2(ch chan string) {
	for {
		fmt.Println("Starting server 2 ...")
		time.Sleep(time.Second * 3)
		ch <- "Server 2 started!!"

	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	fmt.Println("Channels with select")
	fmt.Println("--------------------")

	go server1(ch1)
	go server2(ch2)

	for {
		select {
		case s1 := <-ch1:
			color.Green("Case 1 : %s", s1)
		case s2 := <-ch1:
			color.Cyan("Case 2 : %s", s2)
		case s3 := <-ch2:
			color.Red("Case 3 : %s", s3)
		case s4 := <-ch2:
			color.Yellow("Case 4 : %s", s4)

		}
	}

}
