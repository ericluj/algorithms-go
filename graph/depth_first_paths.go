package graph

import "github.com/ericluj/algorithms-go/lib"

type DepthFirstPaths struct {
	marked []bool // 这个顶点上调用过dfs了吗？
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewDepthFirstPaths(g *Graph, s int) *DepthFirstPaths {
	d := &DepthFirstPaths{
		marked: make([]bool, g.GetV()),
		edgeTo: make([]int, g.GetV()),
		s:      s,
	}
	d.dfs(g, s)
	return d
}

func (d *DepthFirstPaths) dfs(g *Graph, v int) {
	d.marked[v] = true

	for _, w := range g.Adj(v).Data() {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstPaths) hasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DepthFirstPaths) pathTo(v int) *lib.Stack {
	path := lib.NewStack()
	if !d.hasPathTo(v) {
		return path
	}

	for x := v; x != d.s; x = d.edgeTo[x] {
		path.Push(x)
	}
	path.Push(d.s)
	return path
}
