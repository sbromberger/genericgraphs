package traversals

import (
	"log"

	"github.com/sbromberger/genericgraphs"
)

// Logger provides a generic `TraversalProblem` that logs each
// vertex access during traversal.
type Logger[V genericgraphs.Vertex] struct {
	visited map[V]struct{}
	logger  *log.Logger
}

// NewLogger creates a Logger `TraversalProblem` with the
// specified logger.
func NewLogger[V genericgraphs.Vertex](logger *log.Logger) Logger[V] {
	return Logger[V]{visited: make(map[V]struct{}), logger: logger}
}

func (l *Logger[V]) OpenVertex(v V) bool {
	l.logger.Printf("opening %v\n", v)
	return true
}

func (l *Logger[V]) CloseVertex(v V) bool {
	l.logger.Printf("closing %v\n", v)
	return true
}

func (l *Logger[V]) VisitVertex(u, v V) bool {
	l.visited[v] = struct{}{}
	l.logger.Printf("visiting %v -> %v\n", u, v)
	return true
}

func (l *Logger[V]) RevisitVertex(u, v V) bool {
	l.logger.Printf("revisiting %v -> %v\n", u, v)
	return true
}

func (l *Logger[V]) Visited(u V) bool {
	_, found := l.visited[u]
	return found
}

// Uint32Logger provides a specialized `TraversalProblem` that logs each
// vertex access during traversal. This logger is designed for simplegraphs
// with uint32 vertices.
type Uint32Logger struct {
	visited []bool
	logger  *log.Logger
}

// NewUint32Logger creates a Uint32Logger `TraversalProblem` with the
// specified logger.
func NewUint32Logger(g genericgraphs.Graph[uint32], logger *log.Logger) Uint32Logger {
	m := make([]bool, g.Nv())
	return Uint32Logger{visited: m, logger: logger}
}

func (l *Uint32Logger) OpenVertex(v uint32) bool {
	l.logger.Printf("opening %d\n", v)
	return true
}

func (l *Uint32Logger) CloseVertex(v uint32) bool {
	l.logger.Printf("closing %d\n", v)
	return true
}

func (l *Uint32Logger) VisitVertex(u, v uint32) bool {
	l.visited[int(v)] = true
	l.logger.Printf("visiting %d -> %d\n", u, v)
	return true
}

func (l *Uint32Logger) RevisitVertex(u, v uint32) bool {
	l.logger.Printf("revisiting %d -> %d\n", u, v)
	return true
}

func (l *Uint32Logger) Visited(u uint32) bool {
	return l.visited[int(u)]
}
