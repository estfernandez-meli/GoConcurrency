package main

import (
	"fmt"
	"strings"
)

// ping <- chan string, pong chan <-  string
func shout(ping chan string, pong chan string) {
	for {
		receivedFromPing := <-ping // Wait for input from the ping channel

		pong <- fmt.Sprintf("%s!!", strings.ToUpper(receivedFromPing)) // Send modified response to pong channel
	}
}

func main() {
	ping := make(chan string) // Channel for sending user input
	pong := make(chan string) // Channel for receiving processed output

	go shout(ping, pong) // Launch goroutine for processing ping-pong communication

	fmt.Println("Type something and press ENTER (enter Q to quit)")
	for {
		fmt.Printf("-> ")
		var userInput string

		_, _ = fmt.Scanln(&userInput) // Read user input

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput  // Send input to the ping channel
		response := <-pong // Receive response from the pong channel

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels...")
	close(ping) // Close the ping channel to prevent further sends
	close(pong) // Close the pong channel as operations are complete
}
