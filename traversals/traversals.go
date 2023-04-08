package traversals

import (
	"github.com/sbromberger/genericgraphs"
	"github.com/sbromberger/genericgraphs/errors"
)

func BFS[T genericgraphs.Vertex](g genericgraphs.Graph[T], seed T, tp TraversalProblem[T]) error {
	currLevel := make([]T, 0, 4)
	nextLevel := make([]T, 0, 4)

	currLevel = append(currLevel, seed)
	if !tp.VisitVertex(seed, seed) {
		return errors.EarlyTerminationError("visit vertex")
	}

	var u T
	for len(currLevel) > 0 {
		u, currLevel = currLevel[0], currLevel[1:]
		if !tp.OpenVertex(u) {
			return errors.EarlyTerminationError("open vertex")
		}
		neighs, err := g.Neighbors(u)
		if err != nil {
			return err
		}
		for _, v := range neighs {
			if !tp.Visited(v) {
				if !tp.VisitVertex(u, v) {
					return errors.EarlyTerminationError("visit vertex")
				}
				nextLevel = append(nextLevel, v)
			} else {
				if !tp.RevisitVertex(u, v) {
					return errors.EarlyTerminationError("revisit vertex")
				}
			}
		}

		if !tp.CloseVertex(u) {
			return errors.EarlyTerminationError("close vertex")
		}
		nextLevel, currLevel = currLevel, nextLevel
		nextLevel = nextLevel[:0]

	}
	return nil
}

func DFS[T genericgraphs.Vertex](g genericgraphs.Graph[T], seed T, tp TraversalProblem[T]) error {
	stack := make([]T, 0, 4)
	stack = append(stack, seed)

	var u T
	for len(stack) > 0 {
		stack, u = stack[0:len(stack)-1], stack[len(stack)-1]
		if !tp.OpenVertex(u) {
			return errors.EarlyTerminationError("open vertex")
		}
		neighbors, err := g.Neighbors(u)
		if err != nil {
			return err
		}
		for _, v := range neighbors {
			visited := tp.Visited(v)
			if visited {
				if !tp.RevisitVertex(u, v) {
					return errors.EarlyTerminationError("revisited vertex")
				}
			} else {
				if !tp.VisitVertex(u, v) {
					return errors.EarlyTerminationError("visit vertex")
				}

				stack = append(stack, v)
			}
			if !tp.CloseVertex(u) {
				return errors.EarlyTerminationError("close vertex")
			}
		}
	}
	return nil
}
