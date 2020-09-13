package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SimpleSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "find 6 from [1,3,5,7]",
			nums:   []int{1, 3, 5, 6, 7},
			target: 6,
			want:   3,
		},
		{
			name:   "find 2 from [1,2,2,2,3]",
			nums:   []int{1, 2, 2, 2, 3},
			target: 2,
			want:   2,
		},
		{
			name:   "find 2 from [1,2,2,2,3,5]",
			nums:   []int{1, 2, 2, 2, 3, 5},
			target: 4,
			want:   -1,
		},
		{
			name:   "find 1 from [1]",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "find 2 from [1]",
			nums:   []int{1},
			target: 2,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SimpleSearch(tt.nums, tt.target))
		})
	}
}

func Test_LeftSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "find 2 from [1,2,2,2,3]",
			nums:   []int{1, 2, 2, 2, 3},
			target: 2,
			want:   1,
		},
		{
			name:   "find 2 from [1,2,2,2,3,5]",
			nums:   []int{1, 2, 2, 2, 3, 5},
			target: 4,
			want:   -1,
		},
		{
			name:   "find 2 from [1,2,2,2,3,5]",
			nums:   []int{1, 2, 2, 2, 3, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "find 1 from [1]",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "find 2 from [1]",
			nums:   []int{1},
			target: 2,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, LeftSearch(tt.nums, tt.target))
		})
	}
}

func Test_RightSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "find 2 from [1,2,2,2,3]",
			nums:   []int{1, 2, 2, 2, 3},
			target: 2,
			want:   3,
		},
		{
			name:   "find 2 from [1,2,2,2,3,5]",
			nums:   []int{1, 2, 2, 2, 3, 5},
			target: 4,
			want:   -1,
		},
		{
			name:   "find 2 from [1,2,2,2,3,5]",
			nums:   []int{1, 2, 2, 2, 3, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "find 1 from [1]",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "find 2 from [1]",
			nums:   []int{1},
			target: 2,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RightSearch(tt.nums, tt.target))
		})
	}
}
