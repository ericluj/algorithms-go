package graph

import "math"

// 最小生成树的Prim算法的即时实现
type PrimMST struct {
	edgeTo []*Edge         // 距离树最近的边
	marked []bool          // 如果v在树中则为true
	pq     map[int]float64 // 有效的横切边（顶点->权重对应的边）
}

func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	p := &PrimMST{
		edgeTo: make([]*Edge, g.V),
		marked: make([]bool, g.V),
		pq:     make(map[int]float64),
	}
	// 初始给一个最大值，这样顶点第一个遍历到的边为最优边
	for v := 0; v < g.V; v++ {
		p.edgeTo[v] = NewEdge(0, 0, math.MaxFloat64)
	}
	// 默认给一个0顶点，开启for循环
	p.pq[0] = 0.0
	for len(p.pq) != 0 {
		p.visit(g, p.delMinPQ())
	}
	return p
}

func (p *PrimMST) visit(g *EdgeWeightedGraph, v int) {
	// 将顶点v添加到最小生成树中，更新数据
	p.marked[v] = true
	for _, e := range g.Adj(v).Data() {
		w := e.other(v)
		// v-w失效，边已经在最小生成树中
		if p.marked[w] {
			continue
		}
		if e.weight < p.edgeTo[w].weight {
			// 连接w和树的最佳边Edge变为e
			p.edgeTo[w] = e
			p.pq[w] = e.weight
		}
	}
}

// 获取最小有效横切边
func (p *PrimMST) delMinPQ() int {
	var (
		minK      int
		minWeight float64
	)
	for k, v := range p.pq {
		minK = k
		minWeight = v
		break
	}
	for k, v := range p.pq {
		if v < minWeight {
			minK = k
			minWeight = v
		}
	}
	delete(p.pq, minK)
	return minK
}
