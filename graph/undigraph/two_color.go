package undigraph

// 双色问题（能够用两种颜色将图的所有顶点着色，使得任意一条边两个端点颜色不相同）
// 是二分图吗？
type TwoColor struct {
	marked     []bool
	color      []bool
	IsTwoColor bool
}

func NewTwoColor(g *Graph) *TwoColor {
	t := &TwoColor{
		marked:     make([]bool, g.V),
		color:      make([]bool, g.V),
		IsTwoColor: true,
	}
	for s := 0; s < g.V; s++ {
		if !t.marked[s] {
			t.dfs(g, s)
		}
	}
	return t
}

func (t *TwoColor) dfs(g *Graph, v int) {
	t.marked[v] = true
	for _, w := range g.Adj(v).Data() {
		if !t.marked[w] {
			t.color[w] = !t.color[v]
			t.dfs(g, w)
		} else if t.color[w] == t.color[v] {
			t.IsTwoColor = false
		}
	}
}
