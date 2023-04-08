package logger

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
)

type TraversalProblem[V genericgraphs.Vertex] struct {
	visited map[V]struct{}
}

func New[V genericgraphs.Vertex]() TraversalProblem[V] {
	return TraversalProblem[V]{make(map[V]struct{})}
}

func (*TraversalProblem[V]) OpenVertex(v V) bool {
	fmt.Printf("opening %v\n", v)
	return true
}

func (*TraversalProblem[V]) CloseVertex(v V) bool {
	fmt.Printf("closing %v\n", v)
	return true
}

func (tp *TraversalProblem[V]) VisitVertex(u, v V) bool {
	tp.visited[v] = struct{}{}
	fmt.Printf("visiting %v -> %v\n", u, v)
	return true
}

func (*TraversalProblem[V]) RevisitVertex(u, v V) bool {
	fmt.Printf("revisiting %v -> %v\n", u, v)
	return true
}

func (tp *TraversalProblem[V]) Visited(u V) bool {
	_, found := tp.visited[u]
	return found
}
