package backtrack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permute(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    [][]int
	}{
		{
			name:    "1,2,3 => 3!",
			payload: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 3, 1},
				{2, 1, 3},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, Permute(tt.payload))
		})
	}
}

func Test_PermuteUnique(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    [][]int
	}{
		{
			name:    "unique",
			payload: []int{1, 1, 2},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			name:    "unique",
			payload: []int{1, 2, 1},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, PermuteUnique(tt.payload))
		})
	}
}

func Test_CombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "candidates = [2,3,5], target = 8",
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, CombinationSum(tt.candidates, tt.target))
		})
	}
}

func Test_CombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "candidates = [2,3,5], target = 8",
			candidates: []int{2, 3, 5},
			target:     8,
			want: [][]int{
				{3, 5},
			},
		},
		{
			name:       "candidates = [2,2,3,3,5], target = 8",
			candidates: []int{2, 2, 3, 3, 5},
			target:     8,
			want: [][]int{
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name:       "candidates = [10,1,2,7,6,1,5], target = 8,",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			want: [][]int{
				{1, 7},
				{1, 2, 5},
				{2, 6},
				{1, 1, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, CombinationSum2(tt.candidates, tt.target))
		})
	}
}
