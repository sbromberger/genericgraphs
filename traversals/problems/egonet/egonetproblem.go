package egonet

import (
	"github.com/sbromberger/genericgraphs"
)

type TraversalProblem[V genericgraphs.Vertex] struct {
	currLevel int
	maxLevel  int
	levels    map[V]int
}

func New[V genericgraphs.Vertex](maxLevel int) TraversalProblem[V] {
	return TraversalProblem[V]{currLevel: 0, maxLevel: maxLevel, levels: make(map[V]int)}
}

func (tp *TraversalProblem[V]) OpenVertex(V) bool {
	if tp.currLevel > tp.maxLevel {
		return false
	}
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
