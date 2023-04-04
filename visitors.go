package genericgraphs

import "fmt"

type LogUint32Visitor struct {
	visited []bool
}

func (*LogUint32Visitor) OpenVertex(v uint32) bool {
	fmt.Printf("opening %d\n", v)
	return true
}

func (*LogUint32Visitor) CloseVertex(v uint32) bool {
	fmt.Printf("closing %d\n", v)
	return true
}

func (l *LogUint32Visitor) VisitVertex(u, v uint32) bool {
	l.visited[int(v)] = true
	fmt.Printf("visiting %d -> %d\n", u, v)
	return true
}

func (*LogUint32Visitor) RevisitVertex(u, v uint32) bool {
	fmt.Printf("revisiting %d -> %d\n", u, v)
	return true
}

func (l *LogUint32Visitor) Visited(u uint32) bool {
	return l.visited[int(u)]
}

type LevelVisitor[V Vertex] struct {
	currLevel int
	levels    map[V]int
}

func NewLevelVisitor[V Vertex](g Graph[V]) LevelVisitor[V] {
	return LevelVisitor[V]{currLevel: 0, levels: make(map[V]int, g.Nv())}
}

func (l *LevelVisitor[V]) OpenVertex(v V) bool {
	l.currLevel += 1
	return true
}

func (*LevelVisitor[V]) CloseVertex(V) bool {
	return true
}

func (l *LevelVisitor[V]) VisitVertex(_, v V) bool {
	l.levels[v] = l.currLevel
	return true
}

func (*LevelVisitor[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (l *LevelVisitor[V]) Visited(v V) bool {
	_, found := l.levels[v]
	return found
}

func (l *LevelVisitor[V]) Levels() map[V]int {
	return l.levels
}
