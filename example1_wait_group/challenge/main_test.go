package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_UpdateMessage(t *testing.T) {
	//Arrange
	expectedMessage := "Hello, updated!"
	var wg sync.WaitGroup
	//Act
	wg.Add(1)
	go updateMessage("updated", &wg)
	wg.Wait()
	//Assert
	require.Contains(t, msg, expectedMessage)
}

func Test_PrintMessage(t *testing.T) {
	//Arrange
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	msg = "Hello, world :D !"
	//Act
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	//Assert
	require.Contains(t, output, msg)
}

func Test_Main(t *testing.T) {
	//Arrange
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	expectedMessage := "Hello, universe!\nHello, cosmos!\nHello, world!"
	//Act
	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	//Assert
	require.Contains(t, output, expectedMessage)
}
