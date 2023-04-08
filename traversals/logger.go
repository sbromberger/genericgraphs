package traversals

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
)

type Logger[V genericgraphs.Vertex] struct {
	visited map[V]struct{}
}

func NewLogger[V genericgraphs.Vertex]() Logger[V] {
	return Logger[V]{make(map[V]struct{})}
}

func (*Logger[V]) OpenVertex(v V) bool {
	fmt.Printf("opening %v\n", v)
	return true
}

func (*Logger[V]) CloseVertex(v V) bool {
	fmt.Printf("closing %v\n", v)
	return true
}

func (tp *Logger[V]) VisitVertex(u, v V) bool {
	tp.visited[v] = struct{}{}
	fmt.Printf("visiting %v -> %v\n", u, v)
	return true
}

func (*Logger[V]) RevisitVertex(u, v V) bool {
	fmt.Printf("revisiting %v -> %v\n", u, v)
	return true
}

func (tp *Logger[V]) Visited(u V) bool {
	_, found := tp.visited[u]
	return found
}

type Uint32Logger struct {
	visited []bool
}

func NewUint32Logger(g genericgraphs.Graph[uint32]) Uint32Logger {
	m := make([]bool, g.Nv())
	return Uint32Logger{m}
}

func (*Uint32Logger) OpenVertex(v uint32) bool {
	fmt.Printf("opening %d\n", v)
	return true
}

func (*Uint32Logger) CloseVertex(v uint32) bool {
	fmt.Printf("closing %d\n", v)
	return true
}

func (vis *Uint32Logger) VisitVertex(u, v uint32) bool {
	vis.visited[int(v)] = true
	fmt.Printf("visiting %d -> %d\n", u, v)
	return true
}

func (*Uint32Logger) RevisitVertex(u, v uint32) bool {
	fmt.Printf("revisiting %d -> %d\n", u, v)
	return true
}

func (vis *Uint32Logger) Visited(u uint32) bool {
	return vis.visited[int(u)]
}
