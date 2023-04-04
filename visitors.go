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

func (vis *LogUint32Visitor) VisitVertex(u, v uint32) bool {
	vis.visited[int(v)] = true
	fmt.Printf("visiting %d -> %d\n", u, v)
	return true
}

func (*LogUint32Visitor) RevisitVertex(u, v uint32) bool {
	fmt.Printf("revisiting %d -> %d\n", u, v)
	return true
}

func (vis *LogUint32Visitor) Visited(u uint32) bool {
	return vis.visited[int(u)]
}

type LevelVisitor[V Vertex] struct {
	currLevel int
	levels    map[V]int
}

func NewLevelVisitor[V Vertex](g Graph[V]) LevelVisitor[V] {
	return LevelVisitor[V]{currLevel: 0, levels: make(map[V]int, g.Nv())}
}

func (vis *LevelVisitor[V]) OpenVertex(v V) bool {
	vis.currLevel += 1
	return true
}

func (*LevelVisitor[V]) CloseVertex(V) bool {
	return true
}

func (vis *LevelVisitor[V]) VisitVertex(_, v V) bool {
	vis.levels[v] = vis.currLevel
	return true
}

func (*LevelVisitor[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (vis *LevelVisitor[V]) Visited(v V) bool {
	_, found := vis.levels[v]
	return found
}

func (vis *LevelVisitor[V]) Levels() map[V]int {
	return vis.levels
}

type EgoNetVisitor[V Vertex] struct {
	currLevel int
	maxLevel  int
	levels    map[V]int
}

func NewEgoNetVisitor[V Vertex](maxLevel int) EgoNetVisitor[V] {
	return EgoNetVisitor[V]{currLevel: 0, maxLevel: maxLevel, levels: make(map[V]int)}
}

func (vis *EgoNetVisitor[V]) OpenVertex(V) bool {
	if vis.currLevel > vis.maxLevel {
		return false
	}
	vis.currLevel += 1
	return true
}

func (*EgoNetVisitor[V]) CloseVertex(V) bool {
	return true
}

func (vis *EgoNetVisitor[V]) VisitVertex(_, v V) bool {
	vis.levels[v] = vis.currLevel
	return true
}

func (*EgoNetVisitor[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (vis *EgoNetVisitor[V]) Visited(v V) bool {
	_, found := vis.levels[v]
	return found
}
