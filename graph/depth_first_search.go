package graph

type DepthFirstSearch struct {
	marked []bool
	count  int
}

func NewDepthFirstSearch(g *Graph, s int) *DepthFirstSearch {
	d := &DepthFirstSearch{
		marked: make([]bool, g.GetV()),
	}
	d.dfs(g, s)
	return d
}

func (d *DepthFirstSearch) dfs(g *Graph, v int) {
	d.marked[v] = true
	d.count++

	for _, w := range g.Adj(v).data {
		if !d.isMarked(w) {
			d.dfs(g, w)
		}
	}
}

func (d *DepthFirstSearch) isMarked(w int) bool {
	return d.marked[w]
}

func (d *DepthFirstSearch) Count() int {
	return d.count
}
