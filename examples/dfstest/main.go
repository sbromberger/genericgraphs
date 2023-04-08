package main

import "github.com/sbromberger/genericgraphs"

type intgraph struct {
	adjmap map[int][]int
}

func (g *intgraph) Nv() int {
	return len(g.adjmap)
}

func (g *intgraph) Ne() int {
	var acc int
	for _, v := range g.adjmap {
		acc += len(v)
	}
	return acc
}

func (g *intgraph) Neighbors(v int) ([]int, error) {
	n, found := g.adjmap[v]
	if !found {
		return nil, genericgraphs.InvalidVertexError("vertex not found")
	}
	return n, nil
}

func main() {
	adj := make(map[int][]int, 10)
	adj[1] = []int{2, 7, 8}
	adj[2] = []int{3, 6}
	adj[3] = []int{4, 5}
	adj[4] = []int{}
	adj[5] = []int{}
	adj[6] = []int{}
	adj[7] = []int{}
	adj[8] = []int{9, 12}
	adj[9] = []int{10, 11}
	adj[10] = []int{}
	adj[11] = []int{}
	adj[12] = []int{}
	g := intgraph{adjmap: adj}
	l := genericgraphs.NewLogVisitor[int]()
	genericgraphs.DFS[int](&g, 1, &l)
}
