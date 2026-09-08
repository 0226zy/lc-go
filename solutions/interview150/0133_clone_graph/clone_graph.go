package clonegraph

// Node 无向连通图中的节点
// pkg/datastructures 中没有图节点类型，因此在本包内定义。
type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph 克隆图
// 给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）。
// 图中每个节点都包含它的值 val 和其邻居的列表 neighbors。
// 时间复杂度: O(V+E) 每个节点和每条边只处理一次  空间复杂度: O(V) visited 哈希表
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// visited 记录「原节点 -> 克隆节点」，防止环导致无限递归
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if cloned, ok := visited[node]; ok {
		return cloned
	}
	// 先创建克隆节点并登记，再递归填充邻居（应对环）
	cloned := &Node{Val: node.Val}
	visited[node] = cloned
	for _, neighbor := range node.Neighbors {
		cloned.Neighbors = append(cloned.Neighbors, clone(neighbor, visited))
	}
	return cloned
}

// CloneGraphBFS 克隆图（BFS 解法）
// 用队列逐层遍历，visited 同样保存「原节点 -> 克隆节点」的映射。
// 时间复杂度: O(V+E)  空间复杂度: O(V)
func CloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := map[*Node]*Node{node: {Val: node.Val}}
	queue := []*Node{node}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbor := range curr.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			// 把邻居的克隆节点挂到当前节点的克隆节点上
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}
