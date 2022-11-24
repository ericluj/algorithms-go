package undigraph

// 检测环
// 是无环图吗？
type Cycle struct {
	marked   []bool
	HasCycle bool
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

// u为上一步访问过的结点
func (c *Cycle) dfs(g *Graph, v, u int) {
	c.marked[v] = true
	for _, w := range g.Adj(v).Data() {
		// 发现有一个已经被访问过的结点，且这个结点不是上一步访问的结点，则说明有环
		if !c.marked[w] {
			c.dfs(g, w, v)
		} else if w != u {
			c.HasCycle = true
		}
	}
}
