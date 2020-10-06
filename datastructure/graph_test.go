package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// vo-v1
//  |  |
//   -v2
//  |  |
//   -v3
var testGraph = &graph{
	numVexes: 4,
	numEdges: 5,
	vexes:    []string{"v0", "v1", "v2", "v3"},
	edges: [][]bool{
		[]bool{false, true, true, true},
		[]bool{true, false, true, false},
		[]bool{true, true, false, true},
		[]bool{true, false, true, false},
	},
}

func TestGraph_DFSTraverseBySlice(t *testing.T) {
	testGraph.DFSTraverseBySlice()
}
func TestGraph_BFSShortestPath(t *testing.T) {
	tests := []struct {
		name  string
		graph map[string][]string
		from  string
		to    string
		want  []string
	}{
		//       B
		//   /       \
		// A - C - E - F - G
		//   \            /
		//          D
		{
			name: "from A to D",
			graph: map[string][]string{
				"A": {"B", "C", "D"},
				"B": {"F"},
				"C": {"E"},
				"D": {"G"},
				"E": {"F"},
				"F": {"G"},
			},
			from: "A",
			to:   "G",
			want: []string{"A", "D", "G"},
		},
		{
			name: "from C to G",
			graph: map[string][]string{
				"A": {"B", "C", "D"},
				"B": {"F"},
				"C": {"E"},
				"D": {"G"},
				"E": {"F"},
				"F": {"G"},
			},
			from: "C",
			to:   "G",
			want: []string{"C", "E", "F", "G"},
		},
		{
			name: "from A to ?",
			graph: map[string][]string{
				"A": {"B", "C", "D"},
				"B": {"F"},
				"C": {"E"},
				"D": {"G"},
				"E": {"F"},
				"F": {"G"},
			},
			from: "A",
			to:   "?",
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BFSShortestPath(tt.graph, tt.from, tt.to)
			assert.Equal(t, tt.want, got)
		})
	}
}
