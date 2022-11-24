package lib

import "fmt"

// 加权有向边
type DirectedEdge struct {
	v      int     // 边的起点
	w      int     // 边的终点
	Weight float64 // 边的权重
}

func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	return &DirectedEdge{
		v:      v,
		w:      w,
		Weight: weight,
	}
}

func (e *DirectedEdge) From() int {
	return e.v
}

func (e *DirectedEdge) To() int {
	return e.w
}

func (e *DirectedEdge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.Weight)
}
