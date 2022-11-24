package sp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ericluj/algorithms-go/lib"
)

// 加权有向图
type EdgeWeightedDigraph struct {
	V   int                           // 顶点数目
	E   int                           // 边的数目
	adj []*lib.Bag[*lib.DirectedEdge] // 邻接表
}

func NewEdgeWeightedDigraph(v int) *EdgeWeightedDigraph {
	g := &EdgeWeightedDigraph{
		V:   v,
		E:   0,
		adj: make([]*lib.Bag[*lib.DirectedEdge], v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = lib.NewBag[*lib.DirectedEdge]()
	}
	return g
}

// 从文件中读入一幅加权有向图
func NewEdgeWeightedDigraphByFile(fileName string) *EdgeWeightedDigraph {

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

	g := NewEdgeWeightedDigraph(v)
	// 添加边
	for i := 0; (i < e) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		w, _ := strconv.Atoi(arr[1])
		weight, _ := strconv.ParseFloat(arr[2], 64)
		g.AddEdge(lib.NewDirectedEdge(v, w, weight))
	}

	return g
}

// 添加边
func (g *EdgeWeightedDigraph) AddEdge(e *lib.DirectedEdge) {
	g.adj[e.From()].Add(e)
	g.E++
}

// 与v相邻的所有顶点
func (g *EdgeWeightedDigraph) Adj(v int) *lib.Bag[*lib.DirectedEdge] {
	return g.adj[v]
}

// 对象的字符串表示
func (g *EdgeWeightedDigraph) String() string {
	var res string
	for key, val := range g.adj {
		res += fmt.Sprintf("%d:%v\n", key, val.Data())
	}
	return res
}

func (g *EdgeWeightedDigraph) Edges() []*lib.DirectedEdge {
	res := make([]*lib.DirectedEdge, 0)
	for v := 0; v < g.V; v++ {
		for _, e := range g.Adj(v).Data() {
			res = append(res, e)
		}
	}
	return res
}
