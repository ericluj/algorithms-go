package graph

// 是无环图吗？
type Cycle struct {
	marked   []bool
	hasCycle bool
}

func NewCycle(g *Graph) *Cycle {
	c := &Cycle{
		marked: make([]bool, g.V),
	}
	for s := 0; s < g.V; s++ {
		if !c.marked[s] {
			c.dfs(g, s, s)
		}
	}
	return c
}

func (c *Cycle) dfs(g *Graph, v, u int) {
	c.marked[v] = true
	for _, w := range g.Adj(v).Data() {
		if !c.marked[w] {
			c.dfs(g, w, v)
		} else if w != u {
			c.hasCycle = true
		}
	}
}
