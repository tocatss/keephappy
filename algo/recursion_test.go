package algo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RecursionSum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "[1,2,3,4] => 10",
			data: []int{1, 2, 3, 4},
			want: 10,
		},
		{
			name: "[1] => 1",
			data: []int{1},
			want: 1,
		},
		{
			name: "[] => 0",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RecursionSum(tt.data))
		})
	}
}

func Test_RecursionMax(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "[1,4,3,2] => 4",
			data: []int{1, 4, 3, 2},
			want: 4,
		},
		{
			name: "[1,3,5,10,7] => 10",
			data: []int{1, 3, 5, 10, 7},
			want: 10,
		},
		{
			name: "[1] => 1",
			data: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, RecursionMax(tt.data))
		})
	}
}

func Test_RecursionBinSearch(t *testing.T) {
	tests := []struct {
		name     string
		data     []int
		target   int
		want     int
		notFound bool
	}{
		{
			name:   "find 2 in [1,2,3,4] => 1",
			data:   []int{1, 2, 3, 4},
			target: 2,
			want:   1,
		},
		{
			name:   "find 1 in [1,2,3,4] => 0",
			data:   []int{1, 2, 3, 4},
			target: 1,
			want:   0,
		},
		{
			name:   "find 3 in [1,2,3,4] => 2",
			data:   []int{1, 2, 3, 4},
			target: 3,
			want:   2,
		},
		{
			name:   "find 4 in [1,2,3,4] => 3",
			data:   []int{1, 2, 3, 4},
			target: 4,
			want:   3,
		},
		{
			name:   "find 9 in [1,3,5,7,9] => 4",
			data:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
		{
			name:   "find 1 in [1,3,5,7,9] => 0",
			data:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:     "notFound: 2 in [1,3,5,7,9] => 4",
			data:     []int{1, 3, 5, 7, 9},
			target:   2,
			notFound: true,
		},
		{
			name:     "notFound: 4 in [1,3,5,7,9] => 4",
			data:     []int{1, 3, 5, 7, 9},
			target:   4,
			notFound: true,
		},
		{
			name:     "notFound: 6 in [1,3,5,7,9] => 4",
			data:     []int{1, 3, 5, 7, 9},
			target:   6,
			notFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.notFound {
				assert.Less(t, RecursionBinSearch(tt.data, tt.target), 0)
				return
			}

			assert.Equal(t, tt.want, RecursionBinSearch(tt.data, tt.target))
		})
	}
}
