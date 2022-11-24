package mst

import (
	"github.com/ericluj/algorithms-go/base"
	"github.com/ericluj/algorithms-go/lib"
)

// 最小生成树的Kruskal算法
// 按照边的权重顺序（从小到大）处理，将边加入最小生成树
// 加入的边不与已经加入的边构成环，直到树中含有V-1条边
type KruskalMST struct {
	Mst *lib.Queue[*lib.Edge]
}

func NewKruskalMST(g *EdgeWeightedGraph) *KruskalMST {
	k := &KruskalMST{
		Mst: lib.NewQueue[*lib.Edge](),
	}

	pq := lib.NewMinQueue()
	for _, e := range g.Edges() {
		pq.Enqueue(e)
	}

	uf := base.NewUF(g.V)

	for !pq.IsEmpty() && k.Mst.Size() < g.V-1 {
		// 从pq得到权重最小的边和它的顶点
		e := pq.DelMin()
		v := e.Eigher()
		w := e.Other(v)

		// 忽略失效的边
		if uf.Connected(v, w) {
			continue
		}

		// 合并分量
		uf.Union(v, w)
		// 将边添加到最小生成树中
		k.Mst.Enqueue(e)
	}

	return k
}
