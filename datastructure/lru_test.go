package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU_Dump(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "dump zero",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLRUModel(5)
			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_Put(t *testing.T) {
	tests := []struct {
		name string
		f    func(r *lruModel)
		want []string
	}{
		{
			name: "Add v3",
			f: func(r *lruModel) {
				r.Put("v3")
				r.Put("v3")
			},
			want: []string{"v3"},
		},
		{
			name: "Add v1,v2,v3",
			f: func(r *lruModel) {
				r.Put("v1")
				r.Put("v2")
				r.Put("v3")
			},
			want: []string{"v3", "v2", "v1"},
		},
		{
			name: "Add v1,v2,v3,v2",
			f: func(r *lruModel) {
				r.Put("v1")
				r.Put("v2")
				r.Put("v3")
				r.Put("v2")
			},
			want: []string{"v2", "v3", "v1"},
		},
		{
			name: "Add v1,v2,v3,v2,v4,v5,v6",
			f: func(r *lruModel) {
				r.Put("v1")
				r.Put("v2")
				r.Put("v3")
				r.Put("v2")
				r.Put("v4")
				r.Put("v5")
				r.Put("v6")
			},
			want: []string{"v6", "v5", "v4", "v2", "v3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLRUModel(5)
			tt.f(r)

			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_Get(t *testing.T) {
	tests := []struct {
		name    string
		v       string
		f       func(r *lruModel)
		want    string
		wantErr bool
	}{
		{
			name: "ok: Get v3",
			v:    "v3",
			f: func(r *lruModel) {
				r.Put("v3")
			},
			want:    "v3",
			wantErr: false,
		},
		{
			name:    "notexist: Get v3",
			v:       "v3",
			f:       func(r *lruModel) {},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewLRUModel(5)
			tt.f(r)
			got, err := r.Get(tt.v)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
