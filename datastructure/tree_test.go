package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataSource_shift(t *testing.T) {
	tests := []struct {
		name    string
		source  dataSource
		f       func(source *dataSource)
		expect  string
		wantErr bool
	}{
		{
			`shift from ["a","b","c"]`,
			dataSource{data: []string{"a", "b", "c"}},
			func(source *dataSource) {
				_, _ = source.shift()
				_, _ = source.shift()
			},
			"c",
			false,
		},
		{
			`shift from ["a","b"]`,
			dataSource{data: []string{"a", "b"}},
			func(source *dataSource) {
				_, _ = source.shift()
				_, _ = source.shift()
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f(&tt.source)
			got, err := tt.source.shift()
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expect, got)
		})
	}
}

var (
	tree1 = dataSource{data: []string{"a", "b", emptyMark, "c", emptyMark, emptyMark, "d", emptyMark, emptyMark}}
	tree2 = dataSource{data: []string{"a", "b", "c", emptyMark, emptyMark, emptyMark, "e", "f", emptyMark, emptyMark, "g", "h", emptyMark}}
	tree3 = dataSource{data: []string{"a", "b", "d", emptyMark, emptyMark, "e", "f", emptyMark, emptyMark, "g", emptyMark, emptyMark, "c", emptyMark, emptyMark}}
)

func TestDumpByHead(t *testing.T) {
	tests := []struct {
		name   string
		source dataSource
		expect []string
	}{
		{
			`generate tree and headDump`,
			tree1,
			[]string{"a", "b", "c", "d"},
		},
		{
			`generate tree and headDump`,
			tree2,
			[]string{"a", "b", "c", "e", "f", "g", "h"},
		},
		{
			`generate tree and headDump`,
			tree3,
			[]string{"a", "b", "d", "e", "f", "g", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binTree := CreateNodeByHead(&tt.source)()
			assert.Equal(t, tt.expect, binTree.DumpByHead())
		})
	}
}

func TestDumpByMid(t *testing.T) {
	tests := []struct {
		name   string
		source dataSource
		expect []string
	}{
		{
			`generate tree and midDump`,
			tree1,
			[]string{"b", "c", "a", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binTree := CreateNodeByHead(&tt.source)()
			assert.Equal(t, tt.expect, binTree.DumpByMid())
		})
	}
}

func TestDumpByTail(t *testing.T) {
	tests := []struct {
		name   string
		source dataSource
		expect []string
	}{
		{
			`generate tree and midDump`,
			tree1,
			[]string{"c", "b", "d", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binTree := CreateNodeByHead(&tt.source)()
			assert.Equal(t, tt.expect, binTree.DumpTail())
		})
	}
}
