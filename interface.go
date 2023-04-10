package genericgraphs

// Vertex describes the interface for a generic graph vertex.
type Vertex interface {
	comparable
}

// Graph describes a generic graph interface.
type Graph[T Vertex] interface {
	Ne() int
	Nv() int
	Neighbors(v T) ([]T, error)
}
