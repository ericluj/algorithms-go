# 图
## 无向图
| 名称                             | 结构体            |
| -------------------------------- | ----------------- |
| 无向图                           | Graph             |
| 深度优先搜索                     | DepthFirstSearch  |
| 深度优先搜索查找图中路径         | DepthFirstPaths   |
| 广度优先搜索查找图中路径         | BreadthFirstPaths |
| 深度优先搜索找出图中所有连通分量 | CC                |
| 检测环                           | Cycle             |
| 双色问题                         | TwoColor          |
| 符号图                           | SymbolGraph       |

## 有向图
| 名称                               | 结构体            |
| ---------------------------------- | ----------------- |
| 有向图                             | Digraph           |
| 有向图的可达性                     | DirectedDFS       |
| 寻找有向环                         | DirectedCycle     |
| 有向图中基于深度优先搜索的顶点排序 | DepthFirstOrder   |
| 拓扑排序                           | Topological       |
| 计算强连通分量的Kosaraju算法       | KosarajuSCC       |
| 有向图的顶点对可达性               | TransitiveClosure |

## 最小生成树
| 名称                           | 结构体            |
| ------------------------------ | ----------------- |
| 加权无向图                     | EdgeWeightedGraph |
| 最小生成树的Prim算法的延时实现 | LazyPrimMST       |
| 最小生成树的Prim算法的即时实现 | PrimMST           |
| 最小生成树的Kruskal算法        | KruskalMST        |

## 最短路径
| 名称                         | 结构体              |
| ---------------------------- | ------------------- |
| 加权有向图                   | EdgeWeightedDigraph |
| 最短路径的Dijkstra算法       | DijkstraSP          |
| 无环加权有向图的最短路径算法 | AcyclicSP           |