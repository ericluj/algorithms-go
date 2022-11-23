package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 加权无向图
type EdgeWeightedGraph struct {
	V   int           // 顶点数目
	E   int           // 边的数目
	adj []*Bag[*Edge] // 邻接表
}

func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	g := &EdgeWeightedGraph{
		V:   v,
		E:   0,
		adj: make([]*Bag[*Edge], v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = NewBag[*Edge]()
	}
	return g
}

// 从文件中读入一幅加权无向图
func NewEdgeWeightedGraphByFile(fileName string) *EdgeWeightedGraph {

	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 不做错误处理，认为提供的文件一定符合格式
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text()) // 顶点数目
	scanner.Scan()
	e, _ := strconv.Atoi(scanner.Text()) // 边的数目

	g := NewEdgeWeightedGraph(v)
	// 添加边
	for i := 0; (i < e) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		w, _ := strconv.Atoi(arr[1])
		weight, _ := strconv.ParseFloat(arr[2], 64)
		g.AddEdge(NewEdge(v, w, weight))
	}

	return g
}

// 添加边
func (g *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.eigher()
	w := e.other(v)
	g.adj[v].Add(e)
	g.adj[w].Add(e)
	g.E++
}

// 与v相邻的所有顶点
func (g *EdgeWeightedGraph) Adj(v int) *Bag[*Edge] {
	return g.adj[v]
}

// 对象的字符串表示
func (g *EdgeWeightedGraph) String() string {
	var res string
	for key, val := range g.adj {
		res += fmt.Sprintf("%d:%v\n", key, val.Data())
	}
	return res
}

func (g *EdgeWeightedGraph) Edges() []*Edge {
	res := make([]*Edge, 0)
	for v := 0; v < g.V; v++ {
		for _, e := range g.Adj(v).Data() {
			// 7-1 在 1-7已经被加入过了
			if e.other(v) > v {
				res = append(res, e)
			}
		}
	}
	return res
}
