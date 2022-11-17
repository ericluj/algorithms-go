package graph

import (
	"fmt"
	"testing"

	"github.com/ericluj/algorithms-go/lib"
)

// 图
func TestGraph(t *testing.T) {
	g := NewGraphByFile("./data/tinyG.txt")
	fmt.Println(g)
}

// 深度优先搜索
func TestDepthFirstSearch(t *testing.T) {
	g := NewGraphByFile("./data/tinyG.txt")
	d := NewDepthFirstSearch(g, 0)
	fmt.Println(d.Count())
}

// 深度优先搜索查找图中路径
func TestDepthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./data/tinyCG.txt")
	s := 0
	d := NewDepthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}

// 广度优先搜索查找图中路径
func TestBreadthFirstPaths(t *testing.T) {
	g := NewGraphByFile("./data/tinyCG.txt")
	s := 0
	d := NewBreadthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.pathTo(v))
	}
}

// 深度优先搜索找出图中所有连通分量
func TestCC(t *testing.T) {
	g := NewGraphByFile("./data/tinyG.txt")
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
	fmt.Println(cc.connected(6, 0))
	fmt.Println(cc.connected(6, 8))
}

// 检测环
func TestCycle(t *testing.T) {
	g := NewGraphByFile("./data/tinyG.txt")
	c := NewCycle(g)
	fmt.Println(c.hasCycle)
}

// 双色问题
func TestTwoColor(t *testing.T) {
	g := NewGraphByFile("./data/tinyG.txt")
	c := NewTwoColor(g)
	fmt.Println(c.isTwoColor)
}

// 符号图
func TestSymbolGraph(t *testing.T) {
	sg := NewSymbolGraphByFile("./data/routes.txt", " ")
	fmt.Println(sg)
	fmt.Println(sg.g)
}

// 有向图
func TestDigraph(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	fmt.Println(g)
}

// 有向图的可达性
func TestDirectedDFS(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	d := NewDirectedDFS(g, 1)
	fmt.Println(d)
	d2 := NewDirectedDFS(g, 2)
	fmt.Println(d2)
	d3 := NewDirectedDFSSources(g, []int{1, 2, 6})
	fmt.Println(d3)
}

// 寻找有向环
func TestDirectedCycle(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	d := NewDirectedCycle(g)
	fmt.Println(d.hasCycle())
	fmt.Println(d.Cycle())
}

// 有向图中基于深度优先搜索的顶点排序
func TestDepthFirstOrder(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	d := NewDepthFirstOrder(g)
	fmt.Println(d.pre)
	fmt.Println(d.post)
	fmt.Println(d.reversePost)
}

// 拓扑排序
func TestTopological(t *testing.T) {
	sg := NewSymbolDigraphByFile("./data/jobs.txt", "/")
	topo := NewTopological(sg.g)

	for _, v := range topo.order.Data() {
		fmt.Println(sg.Name(v))
	}
}

// 计算强连通分量的Kosaraju算法
func TestKosarajuSCC(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	cc := NewKosarajuSCC(g)
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
	fmt.Println(cc.stronglyConnected(5, 2))
	fmt.Println(cc.stronglyConnected(5, 11))
}

// 有向图的顶点对可达性
func TestTransitiveClosure(t *testing.T) {
	g := NewDigraphByFile("./data/tinyDG.txt")
	tc := NewTransitiveClosure(g)
	fmt.Println(tc.reachable(1, 2))
	fmt.Println(tc.reachable(2, 1))
}
