package mst

import (
	"math"

	"github.com/ericluj/algorithms-go/lib"
)

// 最小生成树的Prim算法的即时实现
type PrimMST struct {
	EdgeTo []*lib.Edge        // 距离树最近的边
	marked []bool             // 如果v在树中则为true
	pq     *lib.IndexMinQueue // 有效的横切边（顶点->权重对应的边）
}

func NewPrimMST(g *EdgeWeightedGraph) *PrimMST {
	p := &PrimMST{
		EdgeTo: make([]*lib.Edge, g.V),
		marked: make([]bool, g.V),
		pq:     lib.NewIndexMinQueue(),
	}
	// 初始给一个最大值，这样顶点第一个遍历到的边为最优边
	for v := 0; v < g.V; v++ {
		p.EdgeTo[v] = lib.NewEdge(0, 0, math.MaxFloat64)
	}
	// 默认给一个0顶点，开启for循环
	p.pq.Insert(0, 0.0)
	for !p.pq.IsEmpty() {
		p.visit(g, p.pq.DelMin())
	}
	return p
}

func (p *PrimMST) visit(g *EdgeWeightedGraph, v int) {
	// 将顶点v添加到最小生成树中，更新数据
	p.marked[v] = true
	for _, e := range g.Adj(v).Data() {
		w := e.Other(v)
		// v-w失效，边已经在最小生成树中
		if p.marked[w] {
			continue
		}
		if e.Weight < p.EdgeTo[w].Weight {
			// 连接w和树的最佳边Edge变为e
			p.EdgeTo[w] = e
			p.pq.Insert(w, e.Weight)
		}
	}
}
