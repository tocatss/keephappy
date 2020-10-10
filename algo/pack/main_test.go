package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroOnePack(t *testing.T) {
	tests := []struct {
		name     string
		packCaps int
		vs       map[string]int
		cs       map[string]int
		want     int
	}{
		{
			name:     "select c",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 2,
				"c": 5,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 6,
			},
			want: 5,
		},
		{
			name:     "select a + b",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 5,
				"c": 5,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 6,
			},
			want: 6,
		},
		{
			name:     "also select a + b",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 5,
				"c": 100000,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 7,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ZeroOnePack(tt.packCaps, tt.vs, tt.cs))
		})
	}
}

func TestZeroOnePackOptimization(t *testing.T) {
	tests := []struct {
		name     string
		packCaps int
		vs       map[string]int
		cs       map[string]int
		want     int
	}{
		{
			name:     "select c",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 2,
				"c": 5,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 6,
			},
			want: 5,
		},
		{
			name:     "select a + b",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 5,
				"c": 5,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 6,
			},
			want: 6,
		},
		{
			name:     "also select a + b",
			packCaps: 6,
			vs: map[string]int{
				"a": 1,
				"b": 5,
				"c": 100000,
			},
			cs: map[string]int{
				"a": 2,
				"b": 1,
				"c": 7,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ZeroOnePackOptimization(tt.packCaps, tt.vs, tt.cs))
		})
	}
}
