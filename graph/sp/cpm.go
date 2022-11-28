package sp

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/ericluj/algorithms-go/lib"
)

// 优先级限制下的并行任务调度问题的关键路径方法
type CPM struct {
	N  int // 任务数
	S  int // 起点
	T  int // 终点
	LP *AcyclicLP
}

func NewCPM(fileName string) *CPM {
	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 不做错误处理，认为提供的文件一定符合格式
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // 顶点数目

	g := NewEdgeWeightedDigraph(2*n + 2)
	s := 2 * n   // 起点
	t := 2*n + 1 // 终点
	for i := 0; (i < n) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		duration, _ := strconv.ParseFloat(arr[0], 64)

		// 每个任务对应三条边
		// 从起始顶点到结束顶点权重为时耗的边
		g.AddEdge(lib.NewDirectedEdge(i, i+n, duration))
		// 从起点到起始顶点权重为零的边
		g.AddEdge(lib.NewDirectedEdge(s, i, 0.0))
		// 从结束顶点到终点权重为零的边
		g.AddEdge(lib.NewDirectedEdge(i+n, t, 0.0))

		// 优先级限制条件对应一条权重为零的边
		for j := 1; j < len(arr); j++ {
			successor, _ := strconv.Atoi(arr[j])
			g.AddEdge(lib.NewDirectedEdge(i+n, successor, 0.0))
		}
	}

	// 寻找最长路径
	return &CPM{
		N:  n,
		S:  s,
		T:  t,
		LP: NewAcyclicLP(g, s),
	}

}
