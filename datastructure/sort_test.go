package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort_islow2HighSorted(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		payload []int
	}{
		{
			name:    "1,2,3 is is low to high sorted",
			payload: []int{1, 2, 3},
			want:    true,
		},
		{
			name:    "1,2,3,4,1 is is not low to high sorted",
			payload: []int{1, 2, 3, 4, 1},
			want:    false,
		},
		{
			name:    "1,3,2 is is not low to high sorted",
			payload: []int{1, 3, 2},
			want:    false,
		},
		{
			name:    "3,2,1 is is not low to high sorted",
			payload: []int{3, 2, 1},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sortAble{
				data: tt.payload,
			}
			assert.Equal(t, tt.want, r.islow2HighSorted())
		})
	}
}

func TestSort_insertSort(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    []int
	}{
		{
			name:    "3,2,1",
			payload: []int{3, 2, 1},
			want:    []int{1, 2, 3},
		},
		{
			name:    "3",
			payload: []int{3},
			want:    []int{3},
		},
		{
			name:    "1,5,4,2,3",
			payload: []int{1, 5, 4, 2, 3},
			want:    []int{1, 2, 3, 4, 5},
		},
		{
			name:    "3,2",
			payload: []int{3, 2},
			want:    []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sortAble{
				data: tt.payload,
			}
			r.insertSort()
			assert.Equal(t, tt.want, r.data)
			assert.Equal(t, true, r.islow2HighSorted())
		})
	}
}

func TestSort_convert2MaxHeap(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    []int
	}{
		{
			name:    "[1,4,5] => [5,1,4]",
			payload: []int{1, 4, 5},
			want:    []int{5, 4, 1},
		},
		{
			name:    "[1,4,5,6,7,2,10] => [10,7,5,6,4,2,1]",
			payload: []int{1, 4, 5, 6, 7, 2, 10},
			want:    []int{10, 7, 5, 6, 4, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sortAble{
				data: tt.payload,
			}
			r.convert2MaxHeap()
			assert.Equal(t, tt.want, r.data)
		})
	}
}

func TestSort_heapSort(t *testing.T) {
	tests := []struct {
		name    string
		payload []int
		want    []int
	}{
		{
			name:    "3,2,1",
			payload: []int{3, 2, 1},
			want:    []int{1, 2, 3},
		},
		{
			name:    "3",
			payload: []int{3},
			want:    []int{3},
		},
		{
			name:    "1,5,4,2,3",
			payload: []int{1, 5, 4, 2, 3},
			want:    []int{1, 2, 3, 4, 5},
		},
		{
			name:    "3,2",
			payload: []int{3, 2},
			want:    []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &sortAble{
				data: tt.payload,
			}
			r.heapSort()
			assert.Equal(t, tt.want, r.data)
			assert.Equal(t, true, r.islow2HighSorted())
		})
	}
}

func TestSort_DivideAndMerge(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "want 1,2,3,5,7",
			s:    []int{1, 3, 2, 7, 5},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			name: "want 1,2,3",
			s:    []int{2, 3, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "want 2,2",
			s:    []int{2, 1, 0, 2},
			want: []int{0, 1, 2, 2},
		},
		{
			name: "want 1",
			s:    []int{1},
			want: []int{1},
		},
		{
			name: "want []",
			s:    nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, DivideAndMerge(tt.s))
		})
	}
}

func TestSort_QuickSort(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "1",
			s:    []int{1},
			want: []int{1},
		},
		{
			name: "21 => 12",
			s:    []int{2, 1},
			want: []int{1, 2},
		},
		{
			s:    []int{2, 1, 5, 3, 7},
			want: []int{1, 2, 3, 5, 7},
		},
		{
			s:    []int{2, 1, 5, 3, 7, 6},
			want: []int{1, 2, 3, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, QuickSort(tt.s))
		})
	}
}
