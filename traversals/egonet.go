package traversals

import (
	"github.com/sbromberger/genericgraphs"
)

// Egonet describes a traversal problem that performs a search up
// to `maxLevel` levels.
type Egonet[V genericgraphs.Vertex] struct {
	currLevel int
	maxLevel  int
	levels    map[V]int
}

// NewEgonet will create a new egonet traversal problem with a maximum depth of `maxLevel`.
func NewEgonet[V genericgraphs.Vertex](maxLevel int) Egonet[V] {
	return Egonet[V]{currLevel: 0, maxLevel: maxLevel, levels: make(map[V]int)}
}

func (tp *Egonet[V]) OpenVertex(V) bool {
	if tp.currLevel > tp.maxLevel {
		return false
	}
	tp.currLevel += 1
	return true
}

func (*Egonet[V]) CloseVertex(V) bool {
	return true
}

func (tp *Egonet[V]) VisitVertex(_, v V) bool {
	tp.levels[v] = tp.currLevel
	return true
}

func (*Egonet[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (tp *Egonet[V]) Visited(v V) bool {
	_, found := tp.levels[v]
	return found
}
