package main

import (
	"fmt"

	"github.com/sbromberger/genericgraphs"
	"github.com/sbromberger/genericgraphs/graphtypes/simple"
	"github.com/sbromberger/genericgraphs/traversals"
	"github.com/sbromberger/genericgraphs/traversals/problems"
)

func main() {
	s := []uint32{0, 1, 2, 3}
	d := []uint32{1, 2, 3, 1}

	sg := simple.New(s, d)
	l := problems.NewLogTraversal[uint32]()
	lev := problems.NewLevelTraversal(genericgraphs.Graph[uint32](&sg))
	fmt.Printf("sg = %v\n", sg)

	err := traversals.BFS[uint32](&sg, 0, &lev)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("levels: %v\n", lev)

	err = traversals.DFS[uint32](&sg, 0, &l)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}
	fmt.Println("DFS done")
}
