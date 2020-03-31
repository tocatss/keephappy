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
			r := newLRU(5)
			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_addOne(t *testing.T) {
	tests := []struct {
		name string
		f    func(r *lru)
		want []string
	}{
		{
			name: "Add 1,2,3 => 3,2,1",
			f: func(r *lru) {
				r.addOne("1")
				r.addOne("2")
				r.addOne("3")
			},
			want: []string{"3", "2", "1"},
		},
		{
			name: "Add 1 => 1",
			f: func(r *lru) {
				r.addOne("1")
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newLRU(5)
			tt.f(r)

			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_removesOne(t *testing.T) {
	tests := []struct {
		name string
		f    func(r *lru)
		want []string
	}{
		{
			name: "remove from 3,2,1 => 3,2",
			f: func(r *lru) {
				r.addOne("1")
				r.addOne("2")
				r.addOne("3")
				r.removeOne()
			},
			want: []string{"3", "2"},
		},
		{
			name: "remove from 1 => []",
			f: func(r *lru) {
				r.addOne("1")
				r.removeOne()
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newLRU(5)
			tt.f(r)

			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_move2First(t *testing.T) {
	tests := []struct {
		name string
		f    func(r *lru)
		want []string
	}{
		{
			name: "move 1 from 3,2,1 to first=> 1,3,2",
			f: func(r *lru) {
				n1 := r.addOne("1")
				r.addOne("2")
				r.addOne("3")
				r.move2First(n1, "A")
			},
			want: []string{"A", "3", "2"},
		},
		{
			name: "remove from 1 => []",
			f: func(r *lru) {
				n1 := r.addOne("1")
				r.move2First(n1, "A")
			},
			want: []string{"A"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newLRU(5)
			tt.f(r)

			assert.Equal(t, tt.want, r.Dump())
		})
	}
}

func TestLRU_Put(t *testing.T) {
	tests := []struct {
		name string
		cap  int
		f    func(r *lru)
		want []string
	}{
		{
			name: "Creat cap:3, PUT 1,2,3,1 => 1 3 2",
			cap:  3,
			f: func(r *lru) {
				r.Put("k1", "1")
				r.Put("k2", "2")
				r.Put("k3", "3")
				r.Put("k1", "1")
			},
			want: []string{"1", "3", "2"},
		},
		{
			name: "Creat cap:3, PUT 1,2,3,2 => 2 3 2",
			cap:  3,
			f: func(r *lru) {
				r.Put("k1", "1")
				r.Put("k2", "2")
				r.Put("k3", "3")
				r.Put("k4", "2")
			},
			want: []string{"2", "3", "2"},
		},
		{
			name: "Creat cap:3, PUT 1,2,3,1,1,1 => 1 3 2",
			cap:  3,
			f: func(r *lru) {
				r.Put("k1", "1")
				r.Put("k2", "2")
				r.Put("k3", "3")
				r.Put("k1", "1")
				r.Put("k1", "1")
				r.Put("k1", "1")
			},
			want: []string{"1", "3", "2"},
		},
		{
			name: "Creat cap:3, PUT 1,2,3,2 => 2 3 1",
			cap:  3,
			f: func(r *lru) {
				r.Put("k1", "1")
				r.Put("k2", "2")
				r.Put("k3", "3")
				r.Put("k2", "2")
			},
			want: []string{"2", "3", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newLRU(tt.cap)
			tt.f(r)
			assert.Equal(t, tt.want, r.Dump())
		})
	}
}
