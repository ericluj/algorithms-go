package lib

import "fmt"

// 加权边
type Edge struct {
	v      int     // 顶点之一
	w      int     // 另一个顶点
	Weight float64 // 边的权重
}

func NewEdge(v, w int, weight float64) *Edge {
	return &Edge{
		v:      v,
		w:      w,
		Weight: weight,
	}
}

func (e *Edge) Eigher() int {
	return e.v
}

func (e *Edge) Other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	return e.v
}

func (e *Edge) compareTo(that *Edge) int {
	if e.Weight < that.Weight {
		return -1
	} else if e.Weight > that.Weight {
		return 1
	} else {
		return 0
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.Weight)
}
