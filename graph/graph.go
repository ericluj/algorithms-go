package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	V   int    // 顶点数目
	E   int    // 边的数目
	adj []*Bag // 邻接表
}

// 创建一个含有v个顶点但不含有边的图
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

// 从文件中读入一幅图
func NewGraphByFile(fileName string) *Graph {

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

	g := NewGraph(v)
	// 添加边
	for i := 0; (i < e) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		w, _ := strconv.Atoi(arr[1])
		g.AddEdge(v, w)
	}

	return g
}

// 顶点数
func (g *Graph) GetV() int {
	return g.V
}

// 边数
func (g *Graph) GetE() int {
	return g.E
}

// 添加边
func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.E++
}

// 与v相邻的所有顶点
func (g *Graph) Adj(v int) *Bag {
	return g.adj[v]
}

// 对象的字符串表示
func (g *Graph) String() string {
	var res string
	for key, val := range g.adj {
		res += fmt.Sprintf("%d:%v\n", key, val)
	}
	return res
}

// 计算v的度数
func (g *Graph) Degree(v int) int {
	return g.Adj(v).Len()
}

// 计算所有顶点的最大度数
func (g *Graph) MaxDegree() int {
	max := 0
	for v := 0; v < g.GetV(); v++ {
		if g.Degree(v) > max {
			max = g.Degree(v)
		}
	}
	return max
}

// 计算所有顶点的平均度数
func (g *Graph) AvgDegree() float64 {
	return 2 * float64(g.GetE()) / float64(g.GetV())
}
