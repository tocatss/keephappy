package datastructure

import "log"

type graph struct {
	numVexes int
	numEdges int
	vexes    []string
	edges    [][]bool
}

func NewGraph(vexes []string, edges [][]bool) *graph {
	return &graph{
		numVexes: len(vexes),
		numEdges: len(edges),
		vexes:    vexes,
		edges:    edges,
	}
}

func (g *graph) DFSTraverseBySlice() {
	visited := make([]bool, g.numVexes)
	for i := 0; i < g.numVexes; i++ {
		if visited[i] {
			continue
		}
		var dfs func(i int)
		dfs = func(i int) {
			visited[i] = true
			for j, v := range g.edges[i] {
				if v && !visited[j] {
					log.Printf("%s-%s", g.vexes[i], g.vexes[j])
					dfs(j)
				}
			}
		}
		dfs(i)
	}
}
