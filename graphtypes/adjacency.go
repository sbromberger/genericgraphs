package graphtypes

import (
	"github.com/sbromberger/genericgraphs"
	"github.com/sbromberger/genericgraphs/errors"
)

type AdjGraph[V genericgraphs.Vertex] struct {
	adjmap map[V]map[V]struct{}
}

func NewAdjacencyGraph[V genericgraphs.Vertex]() AdjGraph[V] {
	m := make(map[V]map[V]struct{})
	return AdjGraph[V]{m}
}

func (g *AdjGraph[V]) Nv() int {
	return len(g.adjmap)
}

func (g *AdjGraph[V]) Ne() int {
	var acc int
	for _, v := range g.adjmap {
		acc += len(v)
	}
	return acc
}

func (g *AdjGraph[V]) AddEdge(u, v V) bool {
	u_neighs, foundU := g.adjmap[u]
	if !foundU {
		g.adjmap[u] = make(map[V]struct{}, 1)
	}
	if _, foundV := u_neighs[v]; foundV {
		return false
	}

	g.adjmap[u][v] = struct{}{}
	return true
}

func (g *AdjGraph[V]) Neighbors(v V) ([]V, error) {
	nbrs, found := g.adjmap[v]
	if !found {
		return nil, errors.InvalidVertexError("vertex not found")
	}

	n := make([]V, 0, len(nbrs))
	for nbr := range nbrs {
		n = append(n, nbr)
	}
	return n, nil
}
