package graph

type Graph struct {
	V   int    // 顶点数目
	E   int    // 边的数目
	adj []*Bag // 邻接表
}

func NewGraph(v int) *Graph {
	g := &Graph{
		V:   v,
		E:   0,
		adj: make([]*Bag, v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = NewBag()
	}
	return g
}

func (g *Graph) GetV() int {
	return g.V
}

func (g *Graph) GetE() int {
	return g.E
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.E++
}

func (g *Graph) Adj(v int) *Bag {
	return g.adj[v]
}
