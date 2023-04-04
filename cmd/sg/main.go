package main

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
)

func main() {
	s := []uint32{0, 1, 2, 3}
	d := []uint32{1, 2, 3, 1}

	sg := genericgraphs.New(s, d)
	// l := logvisit{make([]bool, 0, sg.Nv())}
	lev := genericgraphs.NewLevelVisitor(genericgraphs.Graph[uint32](&sg))
	fmt.Printf("sg = %v\n", sg)

	finished, err := genericgraphs.BFS[uint32](&sg, 0, &lev)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("finished prematurely: %v\n", !finished)
	fmt.Printf("levels: %v\n", lev)
}
