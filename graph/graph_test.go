package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	fmt.Println(g)
}
