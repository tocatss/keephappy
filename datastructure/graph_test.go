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

func Test_Dijkstra(t *testing.T) {
	tests := []struct {
		name  string
		graph map[string]map[string]int
		from  string
		to    string
		want  int
	}{
		{
			name: "A-C-E-F-G",
			//        B
			//   10/     \20
			//   1   2   3   4
			// A - C - E - F - G
			//   \  5        60/
			//          D
			graph: map[string]map[string]int{
				"A": {"B": 10, "C": 1, "D": 5},
				"B": {"F": 20},
				"C": {"E": 2},
				"D": {"G": 60},
				"E": {"F": 3},
				"F": {"G": 4},
			},
			from: "A",
			to:   "G",
			want: 10,
		},
		{
			name: "A-D-G",
			//        B
			//   2/     \7
			//   4   3   2   1
			// A - C - E - F - G
			//   \  8        1/
			//          D
			graph: map[string]map[string]int{
				"A": {"B": 2, "C": 4, "D": 8},
				"B": {"F": 7},
				"C": {"E": 3},
				"D": {"G": 1},
				"E": {"F": 2},
				"F": {"G": 1},
			},
			from: "A",
			to:   "G",
			want: 9,
		},
		{
			name: "A-B-F-G",
			//        B
			//   2/     \6
			//   4   3   2   1
			// A - C - E - F - G
			//   \  8        3/
			//          D
			graph: map[string]map[string]int{
				"A": {"B": 2, "C": 4, "D": 8},
				"B": {"F": 6},
				"C": {"E": 3},
				"D": {"G": 3},
				"E": {"F": 2},
				"F": {"G": 1},
			},
			from: "A",
			to:   "G",
			want: 9,
		},
		{
			name: "un exist",
			//        B
			//   2/     \7
			//   4   3   2   1
			// A - C - E - F - G
			//   \  8        1/
			//          D
			graph: map[string]map[string]int{
				"A": {"B": 2, "C": 4, "D": 8},
				"B": {"F": 7},
				"C": {"E": 3},
				"D": {"G": 1},
				"E": {"F": 2},
				"F": {"G": 1},
			},
			from: "A",
			to:   "?",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dijkstra(tt.graph, tt.from, tt.to)
			assert.Equal(t, tt.want, got)
		})
	}
}
