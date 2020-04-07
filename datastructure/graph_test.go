package datastructure

import "testing"

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
