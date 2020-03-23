package tree

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

func TestDumpFromHead(t *testing.T) {
	tests := []struct {
		name   string
		source dataSource
		expect []string
	}{
		{
			`generate tree and headDump`,
			dataSource{data: []string{"a", "b", emptyMark, "c", emptyMark, emptyMark, "d", emptyMark, emptyMark}},
			[]string{"a", "b", "c", "d"},
		},
		{
			`generate tree and headDump`,
			dataSource{data: []string{"a", "b", "c", emptyMark, emptyMark, emptyMark,
				"e", "f", emptyMark, emptyMark, "g", "h", emptyMark}},
			[]string{"a", "b", "c", "e", "f", "g", "h"},
		},
		{
			`generate tree and headDump`,
			dataSource{data: []string{"a", "b", "d", emptyMark, emptyMark, "e",
				"f", emptyMark, emptyMark, "g", emptyMark, emptyMark, "c", emptyMark, emptyMark}},
			[]string{"a", "b", "d", "e", "f", "g", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binTree := CreateNodeFromHead(&tt.source)()
			assert.Equal(t, tt.expect, binTree.DumpFromHead())
		})
	}
}
