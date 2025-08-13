package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_PrintSomething(t *testing.T) {

	//Arrange

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	//Act
	go printSomething("theta", &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	//Assert
	require.Contains(t, output, "theta")
}
