package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	//Arrange

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	expectedBalance := 34320
	//Act
	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	//Assert
	require.Contains(t, output, fmt.Sprintf("You have earned $%d in 52 weeks (1 year)", expectedBalance))
}
