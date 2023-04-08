package problems

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
)

type LogUint32Traversal struct {
	visited []bool
}

func (*LogUint32Traversal) OpenVertex(v uint32) bool {
	fmt.Printf("opening %d\n", v)
	return true
}

func (*LogUint32Traversal) CloseVertex(v uint32) bool {
	fmt.Printf("closing %d\n", v)
	return true
}

func (vis *LogUint32Traversal) VisitVertex(u, v uint32) bool {
	vis.visited[int(v)] = true
	fmt.Printf("visiting %d -> %d\n", u, v)
	return true
}

func (*LogUint32Traversal) RevisitVertex(u, v uint32) bool {
	fmt.Printf("revisiting %d -> %d\n", u, v)
	return true
}

func (vis *LogUint32Traversal) Visited(u uint32) bool {
	return vis.visited[int(u)]
}

type LogTraversal[V genericgraphs.Vertex] struct {
	visited map[V]struct{}
}

func NewLogTraversal[V genericgraphs.Vertex]() LogTraversal[V] {
	return LogTraversal[V]{make(map[V]struct{})}
}

func (*LogTraversal[V]) OpenVertex(v V) bool {
	fmt.Printf("opening %v\n", v)
	return true
}

func (*LogTraversal[V]) CloseVertex(v V) bool {
	fmt.Printf("closing %v\n", v)
	return true
}

func (vis *LogTraversal[V]) VisitVertex(u, v V) bool {
	vis.visited[v] = struct{}{}
	fmt.Printf("visiting %v -> %v\n", u, v)
	return true
}

func (*LogTraversal[V]) RevisitVertex(u, v V) bool {
	fmt.Printf("revisiting %v -> %v\n", u, v)
	return true
}

func (vis *LogTraversal[V]) Visited(u V) bool {
	_, found := vis.visited[u]
	return found
}

type LevelTraversal[V genericgraphs.Vertex] struct {
	currLevel int
	levels    map[V]int
}

func NewLevelTraversal[V genericgraphs.Vertex](g genericgraphs.Graph[V]) LevelTraversal[V] {
	return LevelTraversal[V]{currLevel: 0, levels: make(map[V]int, g.Nv())}
}

func (vis *LevelTraversal[V]) OpenVertex(v V) bool {
	vis.currLevel += 1
	return true
}

func (*LevelTraversal[V]) CloseVertex(V) bool {
	return true
}

func (vis *LevelTraversal[V]) VisitVertex(_, v V) bool {
	vis.levels[v] = vis.currLevel
	return true
}

func (*LevelTraversal[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (vis *LevelTraversal[V]) Visited(v V) bool {
	_, found := vis.levels[v]
	return found
}

func (vis *LevelTraversal[V]) Levels() map[V]int {
	return vis.levels
}

type EgoNetTraversal[V genericgraphs.Vertex] struct {
	currLevel int
	maxLevel  int
	levels    map[V]int
}

func NewEgoNetTraversal[V genericgraphs.Vertex](maxLevel int) EgoNetTraversal[V] {
	return EgoNetTraversal[V]{currLevel: 0, maxLevel: maxLevel, levels: make(map[V]int)}
}

func (vis *EgoNetTraversal[V]) OpenVertex(V) bool {
	if vis.currLevel > vis.maxLevel {
		return false
	}
	vis.currLevel += 1
	return true
}

func (*EgoNetTraversal[V]) CloseVertex(V) bool {
	return true
}

func (vis *EgoNetTraversal[V]) VisitVertex(_, v V) bool {
	vis.levels[v] = vis.currLevel
	return true
}

func (*EgoNetTraversal[V]) RevisitVertex(uint32, uint32) bool {
	return true
}

func (vis *EgoNetTraversal[V]) Visited(v V) bool {
	_, found := vis.levels[v]
	return found
}
