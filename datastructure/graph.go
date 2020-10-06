package datastructure

import (
	"log"
	"math"
)

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

// 最短路径
//       B
//   /       \
// A - C - E - F - G
//   \           /
//         D
// want : A,D,G
func BFSShortestPath(graph map[string][]string, from, to string) []string {
	if _, ok := graph[from]; !ok {
		return nil
	}

	visited := make(map[string]bool)
	ns := make([]*node, 0, 10)
	ns = append(ns, &node{
		parent: nil,
		value:  from,
	})

	found := bfs(ns, graph, visited, to)
	tail := make([]string, 0, 10)
	for found != nil {
		tail = append(tail, found.value)
		found = found.parent
	}
	res := make([]string, 0, len(tail))
	for i := len(tail) - 1; i >= 0; i-- {
		res = append(res, tail[i])
	}

	return res
}

type node struct {
	parent *node
	value  string
}

func bfs(ns []*node, graph map[string][]string, visited map[string]bool, want string) *node {
	for i := 0; i < len(ns); i++ {
		n := ns[i]
		next, ok := graph[n.value]
		if !ok {
			continue
		}
		for _, v := range next {
			if _, ok := visited[v]; ok {
				continue
			}
			visited[v] = true

			if v == want {
				return &node{
					parent: n,
					value:  v,
				}
			}
			ns = append(ns, &node{
				parent: n,
				value:  v,
			})
		}
	}
	return nil
}

// 狄杰斯特拉计算加权图的最短路径
// 1. 找到当前可访问到的cost最小的点。
// 2. 从该节点出发，找到其相邻节点并更新开销
// 3. 重复1，2直至检查到所有节点。
func Dijkstra(graph map[string]map[string]int, from, to string) int {
	if _, ok := graph[from]; !ok {
		return -1
	}
	if from == to {
		return 0
	}

	var (
		parents = make(map[string]string)
		costs   = make(map[string]int)
		visited = make(map[string]bool)
	)

	// init
	for k, v := range graph[from] {
		costs[k] = v
		parents[k] = from
	}
	weight := -1

	for n := findNextNode(costs, visited, weight); n != ""; n = findNextNode(costs, visited, weight) {
		log.Print("next", n)
		visited[n] = true
		weight = costs[n]

		base := costs[n]
		g, ok := graph[n]
		if !ok {
			continue
		}
		for k, v := range g {
			nv := base + v

			origin, ok := costs[k]
			if !ok || nv < origin {
				// 更新权重。
				costs[k] = nv
				// 更新父亲。
				parents[k] = n
			}
		}
	}

	cost, ok := costs[to]
	if !ok {
		return -1
	}
	return cost
}

func findNextNode(costs map[string]int, visited map[string]bool, weight int) string {
	min := math.MaxInt64
	res := ""

	for k, v := range costs {
		if visited[k] || v < weight {
			continue
		}

		if v < min {
			min = v
			res = k
		}
	}

	return res
}
