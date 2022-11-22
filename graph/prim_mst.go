package graph

// 最小生成树的Prim算法的即时实现
type PrimMST struct {
	edgeTo []*Edge   // 距离树最近的边
	distTo []float64 // distTo[w]=edgeTo[w].weight
	marked []bool    // 如果v在树中则为true
	pq     *MinQueue // 有效的横切边
}

func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	p := &PrimMST{
		edgeTo: make([]*Edge, g.V),
		distTo: make([]float64, g.V),
		marked: make([]bool, g.V),
		pq:     NewMinQueue(),
	}
	p.distTo[0] = 0.0
	return p
}

func (p *EdgeWeightedGraph) visit(g *EdgeWeightedGraph, v int) {

}
