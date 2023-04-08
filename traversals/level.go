package traversals

import (
	"github.com/sbromberger/genericgraphs"
)

type Level[V genericgraphs.Vertex] struct {
	currLevel int
	levels    map[V]int
}

func NewLevel[V genericgraphs.Vertex](g genericgraphs.Graph[V]) Level[V] {
	return Level[V]{currLevel: 0, levels: make(map[V]int, g.Nv())}
}

func (tp *Level[V]) OpenVertex(v V) bool {
	tp.currLevel += 1
	return true
}

func (*Level[V]) CloseVertex(V) bool {
	return true
}

func (tp *Level[V]) VisitVertex(_, v V) bool {
	tp.levels[v] = tp.currLevel
	return true
}

func (*Level[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (tp *Level[V]) Visited(v V) bool {
	_, found := tp.levels[v]
	return found
}

func (tp *Level[V]) Levels() map[V]int {
	return tp.levels
}
