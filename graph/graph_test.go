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

func TestDepthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./tinyCG.txt")
	s := 0
	d := NewDepthFirstPaths(g, s)
	for v := 0; v < g.GetV(); v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}

func TestBreadthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./tinyCG.txt")
	s := 0
	d := NewBreadthFirstPaths(g, s)
	for v := 0; v < g.GetV(); v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}
