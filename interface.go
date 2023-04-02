package gographs

type Vertex interface {
	comparable
}

type Graph[T Vertex] interface {
	Ne() int
	Nv() int
	Neighbors(v T) ([]T, error)
}

type Visitors[T Vertex] interface {
	OpenVertex(v T) error
	VisitVertex(u, v T) error
	RevisitVertex(u, v T) error
	CloseVertex(v T) error
	Visited(v T) bool
}

func BFS[T Vertex](g Graph[T], seed T, visitors Visitors[T]) error {
	visited := make(map[T]struct{}, g.Nv())

	currLevel := make([]T, 0, 4)
	nextLevel := make([]T, 0, 4)

	currLevel = append(currLevel, seed)
	if err := visitors.VisitVertex(seed, seed); err != nil {
		return err
	}

	var u T
	for len(currLevel) > 0 {
		u, currLevel = currLevel[0], currLevel[1:]
		if err := visitors.OpenVertex(u); err != nil {
			return err
		}
		neighs, err := g.Neighbors(u)
		if err != nil {
			return err
		}
		for _, v := range neighs {
			if !visitors.Visited(v) {
				if err := visitors.VisitVertex(u, v); err != nil {
					return err
				}
				nextLevel = append(nextLevel, v)
				visited[v] = struct{}{}
			} else {
				if err := visitors.RevisitVertex(u, v); err != nil {
					return err
				}
			}
		}

		if err := visitors.CloseVertex(u); err != nil {
			return err
		}
		nextLevel, currLevel = currLevel, nextLevel
		nextLevel = nextLevel[:0]

	}
	return nil
}
