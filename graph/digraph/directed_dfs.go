package digraph

import (
	"fmt"
)

// 有向图的可达性
type DirectedDFS struct {
	Marked []bool
}

// 从g中找到s可达的所有顶点
func NewDirectedDFS(g *Digraph, s int) *DirectedDFS {
	d := &DirectedDFS{
		Marked: make([]bool, g.V),
	}
	d.dfs(g, s)

	return d
}

// 在g中找到从sources中的所有顶点可达的所有顶点
func NewDirectedDFSSources(g *Digraph, sources []int) *DirectedDFS {
	d := &DirectedDFS{
		Marked: make([]bool, g.V),
	}

	for _, s := range sources {
		if !d.Marked[s] {
			d.dfs(g, s)
		}
	}

	return d
}

func (d *DirectedDFS) dfs(g *Digraph, v int) {
	d.Marked[v] = true

	for _, w := range g.Adj(v).Data() {
		if !d.Marked[w] {
			d.dfs(g, w)
		}
	}
}

// 打印从起点可达的顶点
func (d *DirectedDFS) String() string {
	var res string
	for v := 0; v < len(d.Marked); v++ {
		if d.Marked[v] {
			res += fmt.Sprintf("%d,", v)
		}
	}
	return res[:len(res)-1]
}
