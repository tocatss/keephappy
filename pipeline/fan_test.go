package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProduceConsume(t *testing.T) {
	tests := []struct {
		name string
		fs   []stageFunc
		want []int
	}{
		{
			name: "+ 1 then x 2",
			fs: []stageFunc{
				stageFunc(func(input int) int {
					return input + 1
				}),
				stageFunc(func(input int) int {
					return input * 2
				}),
			},
			want: []int{4, 6, 8, 10},
		},
		{
			name: "x 2 then + 1",
			fs: []stageFunc{
				stageFunc(func(input int) int {
					return input * 2
				}),
				stageFunc(func(input int) int {
					return input + 1
				}),
			},
			want: []int{3, 5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := producer()
			out := worker(in, tt.fs...)
			var res []int
			for v := range out {
				res = append(res, v)
			}
			assert.ElementsMatch(t, res, tt.want)
		})
	}
}

func Test_FanIn(t *testing.T) {
	tests := []struct {
		name        string
		fs          []stageFunc
		want        []int
		workerCount int
	}{
		{
			name: "2 workers: + 1 then x 2",
			fs: []stageFunc{
				stageFunc(func(input int) int {
					return input + 1
				}),
				stageFunc(func(input int) int {
					return input * 2
				}),
			},
			want:        []int{4, 6, 8, 10},
			workerCount: 2,
		},
		{
			name: "10 workers: + 1 then x 2",
			fs: []stageFunc{
				stageFunc(func(input int) int {
					return input + 1
				}),
				stageFunc(func(input int) int {
					return input * 2
				}),
			},
			want:        []int{4, 6, 8, 10},
			workerCount: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := producer()
			c := make([]<-chan int, tt.workerCount)
			for i := 0; i < tt.workerCount; i++ {
				c[i] = worker(in, tt.fs...)
			}
			out := fanIn(c...)
			var res []int
			for v := range out {
				res = append(res, v)
			}
			assert.ElementsMatch(t, res, tt.want)
		})
	}
}
