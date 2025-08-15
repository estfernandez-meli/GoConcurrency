package main

import (
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func Test_UpdateMsg(t *testing.T) {
	//Arrange
	expectedMsg := "Hello, updated world!"
	var mutex sync.Mutex
	//Act
	wg.Add(1)
	go updateMsg("updated world!", &mutex)
	wg.Wait()
	//Assert
	require.Equal(t, expectedMsg, msg)
}
