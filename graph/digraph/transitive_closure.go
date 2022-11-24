package digraph

// 有向图的顶点对可达性
type TransitiveClosure struct {
	all []*DirectedDFS
}

func NewTransitiveClosure(g *Digraph) *TransitiveClosure {
	t := &TransitiveClosure{
		all: make([]*DirectedDFS, g.V),
	}
	for v := 0; v < g.V; v++ {
		t.all[v] = NewDirectedDFS(g, v)
	}
	return t
}

func (t *TransitiveClosure) Reachable(v, w int) bool {
	return t.all[v].marked[w]
}
