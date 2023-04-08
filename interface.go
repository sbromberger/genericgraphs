package genericgraphs

type Vertex interface {
	comparable
}

type Graph[T Vertex] interface {
	Ne() int
	Nv() int
	Neighbors(v T) ([]T, error)
}
