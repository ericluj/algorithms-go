package graph

import (
	"fmt"
	"testing"

	"github.com/ericluj/algorithms-go/lib"
)

// 图
func TestGraph(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	fmt.Println(g)
}

// 深度优先搜索
func TestDepthFirstSearch(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	d := NewDepthFirstSearch(g, 0)
	fmt.Println(d.Count())
}

// 深度优先搜索查找图中路径
func TestDepthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./tinyCG.txt")
	s := 0
	d := NewDepthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}

// 广度优先搜索查找图中路径
func TestBreadthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./tinyCG.txt")
	s := 0
	d := NewBreadthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}

// 深度优先搜索找出图中所有连通分量
func TestCC(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	cc := NewCC(g)
	fmt.Printf("%d components\n", cc.count)

	components := make([]*lib.Bag, cc.count)
	for i := 0; i < cc.count; i++ {
		components[i] = lib.NewBag()
	}
	for v := 0; v < g.V; v++ {
		components[cc.id[v]].Add(v)
	}
	for i := 0; i < cc.count; i++ {
		fmt.Println(components[i])
	}
}

// 检测环
func TestCycle(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	c := NewCycle(g)
	fmt.Println(c.hasCycle)
}

// 双色问题
func TestTwoColor(t *testing.T) {
	g := NewGraphByFile("./tinyG.txt")
	c := NewTwoColor(g)
	fmt.Println(c.isTwoColor)
}

// 符号图
func TestSymbolGraph(t *testing.T) {
	sg := NewSymbolGraphByFile("./routes.txt")
	fmt.Println(sg)
	fmt.Println(sg.g)
}

// 有向图
func TestDigraph(t *testing.T) {
	g := NewDigraphByFile("./tinyDG.txt")
	fmt.Println(g)
}
