package graph

import "github.com/ericluj/algorithms-go/lib"

type BreadthFirstPaths struct {
	marked []bool // 这个顶点上调用过dfs了吗？
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewBreadthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	b := &DepthFirstPaths{
		marked: make([]bool, g.GetV()),
		edgeTo: make([]int, g.GetV()),
		s:      s,
	}
	b.dfs(g, s)
	return b
}

func (b *BreadthFirstPaths) bfs() {

}

func (b *BreadthFirstPaths) hasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPaths) pathTo(v int) *lib.Stack {
	path := lib.NewStack()
	if !b.hasPathTo(v) {
		return path
	}

	for x := v; x != b.s; x = b.edgeTo[x] {
		path.Push(x)
	}
	path.Push(b.s)
	return path
}
