package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibrecursion(t *testing.T) {
	// 0,1,1,2,3,5
	assert.Equal(t, 5, fibrecursion(5))
	assert.Equal(t, 5, fibdynamic(5))
}

func Test_ClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=6 => want 13",
			n:    6,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ClimbingStairsRecursion(tt.n))
			assert.Equal(t, tt.want, ClimbingStairsDynamic(tt.n))
			assert.Equal(t, tt.want, ClimbingStairsRecWithMemo(tt.n))
		})
	}
}

func Test_MaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[-2,1,-3,4,-1,2,1,-5,4] => [4,-1,2,1] is 6",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MaxSubArray(tt.nums))
		})
	}
}

func Benchmark_ClimbingStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClimbingStairsDynamic(i)
	}
}

func Test_LengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "[10,9,2,5,3,7,101,18] => [2,3,7,101] => 4",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LengthOfLIS(tt.nums))
		})
	}
}

func Test_UniquePaths(t *testing.T) {
	tests := []struct {
		name string
		m    int
		n    int
		want int
	}{
		{
			name: "m = 3, n = 2 => 3",
			m:    3,
			n:    2,
			want: 3,
		},
		{
			name: "m = 7, n = 3 => 28",
			m:    7,
			n:    3,
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UniquePaths(tt.m, tt.n))
		})
	}
}

func Test_UniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name    string
		payload [][]int
		want    int
	}{
		{
			name: "m = 3, n = 2 => 3",
			payload: [][]int{
				[]int{0, 0, 0},
				[]int{0, 0, 0},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UniquePathsWithObstacles(tt.payload))
		})
	}
}
