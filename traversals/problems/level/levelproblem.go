package level

import (
	"github.com/sbromberger/genericgraphs"
)

type TraversalProblem[V genericgraphs.Vertex] struct {
	currLevel int
	levels    map[V]int
}

func New[V genericgraphs.Vertex](g genericgraphs.Graph[V]) TraversalProblem[V] {
	return TraversalProblem[V]{currLevel: 0, levels: make(map[V]int, g.Nv())}
}

func (tp *TraversalProblem[V]) OpenVertex(v V) bool {
	tp.currLevel += 1
	return true
}

func (*TraversalProblem[V]) CloseVertex(V) bool {
	return true
}

func (tp *TraversalProblem[V]) VisitVertex(_, v V) bool {
	tp.levels[v] = tp.currLevel
	return true
}

func (*TraversalProblem[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (tp *TraversalProblem[V]) Visited(v V) bool {
	_, found := tp.levels[v]
	return found
}

func (tp *TraversalProblem[V]) Levels() map[V]int {
	return tp.levels
}
