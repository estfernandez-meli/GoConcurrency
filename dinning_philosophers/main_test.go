package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second

	orderFinished = []string{}
	dine()

	require.Len(t, orderFinished, len(philosophers))
}

func Test_Main_With_Delay(t *testing.T) {
	tests := []struct {
		name  string
		delay time.Duration
	}{
		{
			name:  "without delay",
			delay: time.Second * 0,
		},
		{
			name:  "quarter second delay",
			delay: time.Millisecond * 250,
		},
		{
			name:  "half second delay",
			delay: time.Millisecond * 500,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			eatTime = test.delay
			thinkTime = test.delay
			sleepTime = test.delay

			orderFinished = []string{}
			dine()

			require.Len(t, orderFinished, len(philosophers))
		})
	}
}
