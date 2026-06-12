package graph

import (
	"container/heap"
	"math"
)

// Edge 表示图中的一条边
type Edge struct {
	To     int // 目标节点
	Weight int // 边权重
}

// Graph 使用邻接表表示的带权有向图
type Graph struct {
	n   int
	adj [][]Edge
}

// NewGraph 创建一个包含 n 个节点（编号 0 ~ n-1）的空图
func NewGraph(n int) *Graph {
	return &Graph{
		n:   n,
		adj: make([][]Edge, n),
	}
}

// AddEdge 添加一条从 u 到 v、权重为 w 的有向边
func (g *Graph) AddEdge(u, v, w int) {
	if u < 0 || u >= g.n || v < 0 || v >= g.n {
		return
	}
	g.adj[u] = append(g.adj[u], Edge{To: v, Weight: w})
}

// AddUndirectedEdge 添加一条连接 u 和 v、权重为 w 的无向边
// 无向边等价于两条方向相反的有向边
func (g *Graph) AddUndirectedEdge(u, v, w int) {
	g.AddEdge(u, v, w)
	g.AddEdge(v, u, w)
}

// Size 返回图的节点数量
func (g *Graph) Size() int {
	return g.n
}

// BFS 广度优先搜索
// 从起点 start 出发，按 BFS 顺序返回访问到的节点序列
// 时间复杂度: O(V + E)  空间复杂度: O(V)
func (g *Graph) BFS(start int) []int {
	if start < 0 || start >= g.n {
		return nil
	}
	visited := make([]bool, g.n)
	queue := make([]int, 0, g.n)
	result := make([]int, 0, g.n)

	visited[start] = true
	queue = append(queue, start)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		result = append(result, u)

		for _, e := range g.adj[u] {
			if !visited[e.To] {
				visited[e.To] = true
				queue = append(queue, e.To)
			}
		}
	}
	return result
}

// DFS 深度优先搜索
// 从起点 start 出发，按 DFS 顺序返回访问到的节点序列
// 时间复杂度: O(V + E)  空间复杂度: O(V)
func (g *Graph) DFS(start int) []int {
	if start < 0 || start >= g.n {
		return nil
	}
	visited := make([]bool, g.n)
	result := make([]int, 0, g.n)
	var dfs func(int)
	dfs = func(u int) {
		visited[u] = true
		result = append(result, u)
		for _, e := range g.adj[u] {
			if !visited[e.To] {
				dfs(e.To)
			}
		}
	}
	dfs(start)
	return result
}

// BellmanFord 贝尔曼-福特算法
// 返回从起点 start 到每个节点的最短距离，以及图中是否存在负权环
// 若 dist[i] == math.MaxInt64，表示从 start 到 i 不可达
// 时间复杂度: O(V * E)  空间复杂度: O(V)
func (g *Graph) BellmanFord(start int) ([]int, bool) {
	if start < 0 || start >= g.n {
		return nil, false
	}
	dist := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	// 最多进行 n-1 轮松弛
	for i := 0; i < g.n-1; i++ {
		updated := false
		for u := 0; u < g.n; u++ {
			if dist[u] == math.MaxInt64 {
				continue
			}
			for _, e := range g.adj[u] {
				if dist[e.To] > dist[u]+e.Weight {
					dist[e.To] = dist[u] + e.Weight
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	// 第 n 轮还能松弛，说明存在负权环
	for u := 0; u < g.n; u++ {
		if dist[u] == math.MaxInt64 {
			continue
		}
		for _, e := range g.adj[u] {
			if dist[e.To] > dist[u]+e.Weight {
				return dist, true
			}
		}
	}

	return dist, false
}

// item 是 Dijkstra 优先队列中的元素
type item struct {
	node int
	dist int
}

// priorityQueue 是实现 container/heap 接口的最小堆，按 dist 升序排列
type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }

func (pq priorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*item))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	it := old[n-1]
	*pq = old[:n-1]
	return it
}

// Dijkstra 迪杰斯特拉算法
// 返回从起点 start 到每个节点的最短距离
// 要求图中不存在负权边；若 dist[i] == math.MaxInt64，表示从 start 到 i 不可达
// 时间复杂度: O((V + E) * log V)  空间复杂度: O(V)
func (g *Graph) Dijkstra(start int) []int {
	if start < 0 || start >= g.n {
		return nil
	}
	dist := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &item{node: start, dist: 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(*item)
		u := cur.node
		if cur.dist > dist[u] {
			continue
		}
		for _, e := range g.adj[u] {
			if dist[e.To] > dist[u]+e.Weight {
				dist[e.To] = dist[u] + e.Weight
				heap.Push(pq, &item{node: e.To, dist: dist[e.To]})
			}
		}
	}

	return dist
}
