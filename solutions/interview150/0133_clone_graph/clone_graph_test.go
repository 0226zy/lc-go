package clonegraph

import "testing"

// buildGraph 按 LeetCode 的邻接表表示构建图，返回节点数组
// adjList[i] 表示节点 i+1 的邻居值列表
func buildGraph(adjList [][]int) []*Node {
	if len(adjList) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbors := range adjList {
		for _, v := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[v-1])
		}
	}
	return nodes
}

// graphToAdjList 把图转回邻接表（按邻居值排序），用于比较
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]bool)
	var dfs func(n *Node)
	order := make(map[int][]int)
	dfs = func(n *Node) {
		if visited[n] {
			return
		}
		visited[n] = true
		for _, nb := range n.Neighbors {
			order[n.Val] = append(order[n.Val], nb.Val)
			dfs(nb)
		}
	}
	dfs(node)
	maxVal := 0
	for v := range order {
		if v > maxVal {
			maxVal = v
		}
	}
	result := make([][]int, maxVal)
	for v, list := range order {
		result[v-1] = list
	}
	return result
}

// assertCloned 校验克隆结果：结构一致、节点是全新副本
func assertCloned(t *testing.T, adjList [][]int, cloneFn func(*Node) *Node) {
	t.Helper()
	nodes := buildGraph(adjList)
	var head *Node
	if nodes != nil {
		head = nodes[0]
	}
	cloned := cloneFn(head)

	if adjList == nil || len(adjList) == 0 {
		if cloned != nil {
			t.Errorf("期望 nil，得到 %v", cloned.Val)
		}
		return
	}

	// 校验结构一致
	got := graphToAdjList(cloned)
	if len(got) != len(adjList) {
		t.Fatalf("节点数不一致: got %d, want %d", len(got), len(adjList))
	}
	for i := range adjList {
		if len(got[i]) != len(adjList[i]) {
			t.Fatalf("节点 %d 邻居数不一致: got %v, want %v", i+1, got[i], adjList[i])
		}
	}

	// 校验是深拷贝：克隆图中所有节点都是新节点
	origVisited := make(map[*Node]bool)
	var markOrig func(n *Node)
	markOrig = func(n *Node) {
		if origVisited[n] {
			return
		}
		origVisited[n] = true
		for _, nb := range n.Neighbors {
			markOrig(nb)
		}
	}
	markOrig(head)

	cloneVisited := make(map[*Node]bool)
	var checkNew func(n *Node)
	checkNew = func(n *Node) {
		if cloneVisited[n] {
			return
		}
		if origVisited[n] {
			t.Errorf("节点 %d 没有深拷贝，与原图共享内存", n.Val)
		}
		cloneVisited[n] = true
		for _, nb := range n.Neighbors {
			checkNew(nb)
		}
	}
	checkNew(cloned)
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 四个节点的环", [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		{"示例2: 单节点自环", [][]int{{1}}},
		{"示例3: 空图", nil},
		{"两个节点互相连接", [][]int{{2}, {1}}},
		{"三角形", [][]int{{2, 3}, {1, 3}, {1, 2}}},
		{"链式图", [][]int{{2}, {1, 3}, {2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCloned(t, tt.adjList, CloneGraph)
		})
	}
}

func TestCloneGraphBFS(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{"示例1: 四个节点的环", [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
		{"示例2: 单节点自环", [][]int{{1}}},
		{"示例3: 空图", nil},
		{"三角形", [][]int{{2, 3}, {1, 3}, {1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertCloned(t, tt.adjList, CloneGraphBFS)
		})
	}
}

// buildLargeGraph 构建 n 个节点组成的环（节点 i 连接 i-1 和 i+1），用于基准测试
func buildLargeGraph(n int) *Node {
	if n == 0 {
		return nil
	}
	nodes := make([]*Node, n)
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}
	for i := range nodes {
		nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[(i+1)%n])
		nodes[(i+1)%n].Neighbors = append(nodes[(i+1)%n].Neighbors, nodes[i])
	}
	return nodes[0]
}

func BenchmarkCloneGraph(b *testing.B) {
	b.Run("100节点", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CloneGraph(buildLargeGraph(100))
		}
	})
	b.Run("1000节点", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CloneGraph(buildLargeGraph(1000))
		}
	})
}
