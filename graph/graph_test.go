package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	fmt.Println(g)
}

func TestDepthFirstSearch(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	d := NewDepthFirstSearch(g, 0)
	fmt.Println(d.Count())
}
