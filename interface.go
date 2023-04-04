package genericgraphs

import "fmt"

type Vertex interface {
	comparable
}

type Graph[T Vertex] interface {
	Ne() int
	Nv() int
	Neighbors(v T) ([]T, error)
}

type Visitors[T Vertex] interface {
	OpenVertex(v T) bool
	VisitVertex(u, v T) bool
	RevisitVertex(u, v T) bool
	CloseVertex(v T) bool
	Visited(v T) bool
}

type EarlyTerminationError string

func (e EarlyTerminationError) Error() string {
	return fmt.Sprintf("EarlyTerminationError: %s", string(e))
}

type InvalidVertexError string

func (e *InvalidVertexError) Error() string {
	return fmt.Sprintf("InvalidVertexError: %s", string(*e))
}

func BFS[T Vertex](g Graph[T], seed T, visitors Visitors[T]) (bool, error) {
	currLevel := make([]T, 0, 4)
	nextLevel := make([]T, 0, 4)

	currLevel = append(currLevel, seed)
	if !visitors.VisitVertex(seed, seed) {
		return false, EarlyTerminationError("visit vertex")
	}

	var u T
	for len(currLevel) > 0 {
		u, currLevel = currLevel[0], currLevel[1:]
		if !visitors.OpenVertex(u) {
			return false, EarlyTerminationError("open vertex")
		}
		neighs, err := g.Neighbors(u)
		if err != nil {
			return false, err
		}
		for _, v := range neighs {
			if !visitors.Visited(v) {
				if !visitors.VisitVertex(u, v) {
					return false, EarlyTerminationError("visit vertex")
				}
				nextLevel = append(nextLevel, v)
			} else {
				if !visitors.RevisitVertex(u, v) {
					return false, EarlyTerminationError("revisit vertex")
				}
			}
		}

		if !visitors.CloseVertex(u) {
			return false, EarlyTerminationError("close vertex")
		}
		nextLevel, currLevel = currLevel, nextLevel
		nextLevel = nextLevel[:0]

	}
	return true, nil
}
