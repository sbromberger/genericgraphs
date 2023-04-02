package main

import (
	"fmt"

	"github.com/sbromberger/gographs"
)

type logvisit struct {
	visited []bool
}

func (*logvisit) OpenVertex(v uint32) error {
	fmt.Printf("opening %d\n", v)
	return nil
}

func (*logvisit) CloseVertex(v uint32) error {
	fmt.Printf("closing %d\n", v)
	return nil
}

func (l *logvisit) VisitVertex(u, v uint32) error {
	l.visited[int(v)] = true
	fmt.Printf("visiting %d -> %d\n", u, v)
	return nil
}

func (*logvisit) RevisitVertex(u, v uint32) error {
	fmt.Printf("revisiting %d -> %d\n", u, v)
	return nil
}

func (l *logvisit) Visited(u uint32) bool {
	return l.visited[int(u)]
}

type Levels[V gographs.Vertex] struct {
	currLevel int
	levels    map[V]int
}

func (l *Levels[V]) OpenVertex(v V) error {
	l.currLevel += 1
	return nil
}

func (*Levels[V]) CloseVertex(V) error {
	return nil
}

func (l *Levels[V]) VisitVertex(_, v V) error {
	l.levels[v] = l.currLevel
	return nil
}

func (*Levels[V]) RevisitVertex(uint32, uint32) error {
	return nil
}

func (l *Levels[V]) Visited(v V) bool {
	_, found := l.levels[v]
	return found
}

func (l *Levels[V]) Levels() map[V]int {
	return l.levels
}

func main() {
	s := []uint32{0, 1, 2, 3}
	d := []uint32{1, 2, 3, 1}

	sg := gographs.New(s, d)
	// l := logvisit{ make([]bool, 0, nv(sg))}
	lev := Levels[uint32]{0, make(map[uint32]int, sg.Nv())}
	fmt.Printf("sg = %v\n", sg)

	err := gographs.BFS[uint32](&sg, 0, &lev)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("levels: %v\n", lev)
}
