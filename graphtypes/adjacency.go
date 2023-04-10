package graphtypes

import (
	"github.com/sbromberger/genericgraphs"
	"github.com/sbromberger/genericgraphs/errors"
)

// AdjGraph is a graph described by a map of maps storing adjacency information for
// a given vertex.
type AdjGraph[V genericgraphs.Vertex] struct {
	adjmap map[V]map[V]struct{}
}

// NewAdjacencyGraph creates a new adjacency graph.
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

// AddEdge adds an edge from `u` to `v` to the given adjacency graph.
// If the vertices are not yet in the graph, they will be created.
func (g *AdjGraph[V]) AddEdge(u, v V) bool {
	u_neighs, foundU := g.adjmap[u]
	if !foundU {
		g.adjmap[u] = make(map[V]struct{}, 1)
	}
	if _, foundV := g.adjmap[v]; !foundV {
		g.adjmap[v] = make(map[V]struct{}, 1)
	}
	if _, foundUV := u_neighs[v]; foundUV {
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
