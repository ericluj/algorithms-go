package graph

// 最小生成树的Prim算法的延时实现
type LazyPrimMST struct {
	marked []bool        // 最小生成树的顶点
	mst    *Queue[*Edge] // 最小生成树的边
	pq     *MinQueue     // 横切边（包括失效的边）
}

func NewLazyPrimMST(g *EdgeWeightedGraph) *LazyPrimMST {
	p := &LazyPrimMST{
		marked: make([]bool, g.V),
		mst:    NewQueue[*Edge](),
		pq:     NewMinQueue(),
	}
	p.visit(g, 0) //假设g是连通的

	for !p.pq.IsEmpty() {
		// 从pq中得到权重最小的边
		e := p.pq.DelMin()
		// 跳过失效的边
		v := e.eigher()
		w := e.other(v)
		if p.marked[v] && p.marked[w] {
			continue
		}
		// 将边添加到树中
		p.mst.Enqueue(e)
		// 将顶点v或w添加到树中
		if !p.marked[v] {
			p.visit(g, v)
		}
		if !p.marked[w] {
			p.visit(g, w)
		}
	}

	return p
}

func (p *LazyPrimMST) visit(g *EdgeWeightedGraph, v int) {
	// 标记顶点v并将所有连接v和未被标记顶点加入pq
	p.marked[v] = true
	for _, e := range g.Adj(v).Data() {
		if !p.marked[e.other(v)] {
			p.pq.Enqueue(e)
		}
	}
}
