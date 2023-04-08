package logger

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
)

type LogUint32Traversal struct {
	visited []bool
}

func NewUint32(g genericgraphs.Graph[uint32]) LogUint32Traversal {
	m := make([]bool, 0, g.Nv())
	return LogUint32Traversal{m}
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
