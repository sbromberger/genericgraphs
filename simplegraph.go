package genericgraphs

import (
	"sort"

	"github.com/sbromberger/graphmatrix"
)

type SimpleGraph struct {
	ne  int
	adj graphmatrix.GraphMatrix
}

func (g *SimpleGraph) Nv() int {
	return int(g.adj.N())
}

func (g *SimpleGraph) Ne() int {
	return g.ne
}

func (g *SimpleGraph) Neighbors(v uint32) ([]uint32, error) {
	return g.adj.GetRow(v)
}

type edge struct {
	s, d uint32
}

func New(srcs, dsts []uint32) SimpleGraph {
	edges := make([]edge, len(srcs))
	for i, s := range srcs {
		d := dsts[i]
		e := edge{s, d}
		edges[i] = e
	}

	sort.Slice(edges, func(i, j int) bool {
		if edges[i].s == edges[i].d {
			return edges[i].d < edges[j].d
		}
		return edges[i].s < edges[j].s
	})
	for i := range srcs {
		srcs[i] = edges[i].s
		dsts[i] = edges[i].d
	}
	adj, err := graphmatrix.NewFromSortedIJ(srcs, dsts)
	if err != nil {
		panic(err)
	}
	ne := len(srcs)
	return SimpleGraph{adj: adj, ne: ne}
}
