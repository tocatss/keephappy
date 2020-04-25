package main

import (
	"log"
	"sync"
	"testing"
	"time"
)

func Test_allowAction(t *testing.T) {
	tests := []struct {
		name    string
		payload *funnel
		routine int
		isAllow bool
	}{
		{
			name:    "10 times in 5s => 10/s",
			payload: newFunnel(10, 10),
			isAllow: true,
			routine: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wg sync.WaitGroup
			for i := tt.routine; i > 0; i-- {
				wg.Add(1)
				go func(i int) {
					isAllow := allowAction(tt.payload)
					log.Print(i, isAllow)
					time.Sleep(1 * time.Second)

					isAllow = allowAction(tt.payload)
					log.Print(i, isAllow)

					time.Sleep(1 * time.Second)

					isAllow = allowAction(tt.payload)
					log.Print(i, isAllow)

					wg.Done()
				}(i)
			}
			wg.Wait()
		})
	}
}
