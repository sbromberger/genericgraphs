package traversals

import (
	"github.com/sbromberger/genericgraphs"
)

type TraversalProblem[T genericgraphs.Vertex] interface {
	OpenVertex(v T) bool
	VisitVertex(u, v T) bool
	RevisitVertex(u, v T) bool
	CloseVertex(v T) bool
	Visited(v T) bool
}
