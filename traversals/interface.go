package traversals

import (
	"github.com/sbromberger/genericgraphs"
)

// TraveralProblem is an interface that describes a set of
// visitor functions along with a visit function and whatever
// associated state is required for the specific problem.
type TraversalProblem[T genericgraphs.Vertex] interface {
	OpenVertex(v T) bool
	VisitVertex(u, v T) bool
	RevisitVertex(u, v T) bool
	CloseVertex(v T) bool
	Visited(v T) bool
}
