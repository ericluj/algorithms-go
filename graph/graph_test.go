package graph

import (
	"fmt"
	"testing"

	"github.com/ericluj/algorithms-go/graph/digraph"
	"github.com/ericluj/algorithms-go/graph/mst"
	"github.com/ericluj/algorithms-go/graph/sp"
	"github.com/ericluj/algorithms-go/graph/undigraph"
	"github.com/ericluj/algorithms-go/lib"
)

// 图
func TestGraph(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyG.txt")
	fmt.Println(g)
}

// 深度优先搜索
func TestDepthFirstSearch(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyG.txt")
	d := undigraph.NewDepthFirstSearch(g, 0)
	fmt.Println(d.Count())
}

// 深度优先搜索查找图中路径
func TestDepthFirstPaths(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyCG.txt")
	s := 0
	d := undigraph.NewDepthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.PathTo(v))
	}
}

// 广度优先搜索查找图中路径
func TestBreadthFirstPaths(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyCG.txt")
	s := 0
	d := undigraph.NewBreadthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d: %v\n", s, v, d.PathTo(v))
	}
}

// 深度优先搜索找出图中所有连通分量
func TestCC(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyG.txt")
	cc := undigraph.NewCC(g)
	fmt.Printf("%d components\n", cc.Count)

	components := make([]*lib.Bag[int], cc.Count)
	for i := 0; i < cc.Count; i++ {
		components[i] = lib.NewBag[int]()
	}
	for v := 0; v < g.V; v++ {
		components[cc.Id[v]].Add(v)
	}
	for i := 0; i < cc.Count; i++ {
		fmt.Println(components[i])
	}
	fmt.Println(cc.Connected(6, 0))
	fmt.Println(cc.Connected(6, 8))
}

// 检测环
func TestCycle(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyG.txt")
	c := undigraph.NewCycle(g)
	fmt.Println(c.HasCycle)
}

// 双色问题
func TestTwoColor(t *testing.T) {
	g := undigraph.NewGraphByFile("./data/tinyG.txt")
	c := undigraph.NewTwoColor(g)
	fmt.Println(c.IsTwoColor)
}

// 符号图
func TestSymbolGraph(t *testing.T) {
	sg := undigraph.NewSymbolGraphByFile("./data/routes.txt", " ")
	fmt.Println(sg)
	fmt.Println(sg.G())
}

// 有向图
func TestDigraph(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	fmt.Println(g)
}

// 有向图的可达性
func TestDirectedDFS(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	d := digraph.NewDirectedDFS(g, 1)
	fmt.Println(d)
	d2 := digraph.NewDirectedDFS(g, 2)
	fmt.Println(d2)
	d3 := digraph.NewDirectedDFSSources(g, []int{1, 2, 6})
	fmt.Println(d3)
}

// 寻找有向环
func TestDirectedCycle(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	d := digraph.NewDirectedCycle(g)
	fmt.Println(d.HasCycle())
	fmt.Println(d.Cycle())
}

// 有向图中基于深度优先搜索的顶点排序
func TestDepthFirstOrder(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	d := digraph.NewDepthFirstOrder(g)
	fmt.Println(d.Pre)
	fmt.Println(d.Post)
	fmt.Println(d.ReversePost)
}

// 拓扑排序
func TestTopological(t *testing.T) {
	sg := digraph.NewSymbolDigraphByFile("./data/jobs.txt", "/")
	topo := digraph.NewTopological(sg.G())

	for _, v := range topo.Order.Data() {
		fmt.Println(sg.Name(v))
	}
}

// 计算强连通分量的Kosaraju算法
func TestKosarajuSCC(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	cc := digraph.NewKosarajuSCC(g)
	fmt.Printf("%d components\n", cc.Count)

	components := make([]*lib.Bag[int], cc.Count)
	for i := 0; i < cc.Count; i++ {
		components[i] = lib.NewBag[int]()
	}
	for v := 0; v < g.V; v++ {
		components[cc.Id[v]].Add(v)
	}
	for i := 0; i < cc.Count; i++ {
		fmt.Println(components[i])
	}
	fmt.Println(cc.StronglyConnected(5, 2))
	fmt.Println(cc.StronglyConnected(5, 11))
}

// 有向图的顶点对可达性
func TestTransitiveClosure(t *testing.T) {
	g := digraph.NewDigraphByFile("./data/tinyDG.txt")
	tc := digraph.NewTransitiveClosure(g)
	fmt.Println(tc.Reachable(1, 2))
	fmt.Println(tc.Reachable(2, 1))
}

// 加权无向图
func TestEdgeWeightedGraph(t *testing.T) {
	g := mst.NewEdgeWeightedGraphByFile("./data/tinyEWG.txt")
	fmt.Println(g)
}

// 最小生成树的Prim算法的延时实现
func TestLazyPrimMST(t *testing.T) {
	g := mst.NewEdgeWeightedGraphByFile("./data/tinyEWG.txt")
	l := mst.NewLazyPrimMST(g)
	fmt.Println(l.Mst)
}

// 最小生成树的Prim算法的即时实现
func TestPrimMST(t *testing.T) {
	g := mst.NewEdgeWeightedGraphByFile("./data/tinyEWG.txt")
	l := mst.NewPrimMST(g)
	fmt.Println(l.EdgeTo[1:])
}

// 最小生成树的Kruskal算法
func TestKruskalMST(t *testing.T) {
	g := mst.NewEdgeWeightedGraphByFile("./data/tinyEWG.txt")
	k := mst.NewKruskalMST(g)
	fmt.Println(k.Mst)
}

// 加权有向图
func TestEdgeWeightedDigraph(t *testing.T) {
	g := sp.NewEdgeWeightedDigraphByFile("./data/tinyEWD.txt")
	fmt.Println(g)
}

// 最短路径的Dijkstra算法
func TestDijkstraSP(t *testing.T) {
	g := sp.NewEdgeWeightedDigraphByFile("./data/tinyEWD.txt")
	s := 0
	d := sp.NewDijkstraSP(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d (%.2f): %v\n", s, v, d.DistTo[v], d.PathTo(v))
	}
}

// 无环加权有向图的最短路径算法
func TestAcyclicSP(t *testing.T) {
	g := sp.NewEdgeWeightedDigraphByFile("./data/tinyEWDAG.txt")
	s := 5
	d := sp.NewAcyclicSP(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d (%.2f): %v\n", s, v, d.DistTo[v], d.PathTo(v))
	}
}

// 无环加权有向图的最长路径算法
func TestAcyclicLP(t *testing.T) {
	g := sp.NewEdgeWeightedDigraphByFile("./data/tinyEWDAG.txt")
	s := 5
	d := sp.NewAcyclicLP(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d (%.2f): %v\n", s, v, d.DistTo[v], d.PathTo(v))
	}
}

// 优先级限制下的并行任务调度问题的关键路径方法
func TestCPM(t *testing.T) {
	c := sp.NewCPM("./data/jobsPC.txt")
	fmt.Printf("Start times:\n")
	for i := 0; i < c.N; i++ {
		fmt.Printf("%d : %.1f\n", i, c.LP.DistTo[i])
	}
	fmt.Printf("Finish time: %.1f\n", c.LP.DistTo[c.T])
}

// 基于队列的BellmanFord算法
func TestBellmanFordSP(t *testing.T) {
	g := sp.NewEdgeWeightedDigraphByFile("./data/tinyEWDn.txt")
	s := 0
	d := sp.NewBellmanFordSP(g, s)
	for v := 0; v < g.V; v++ {
		fmt.Printf("%d to %d (%.2f): %v\n", s, v, d.DistTo[v], d.PathTo(v))
	}
	fmt.Println(d.Cycle)
}
