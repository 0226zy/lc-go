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
// 时间复杂度: O(V+E) 每个节点和每条边只处理一次  空间复杂度: O(V) visited 哈希表 + 递归栈
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	// visited 记录「原节点 -> 克隆节点」的映射，既防止环导致无限递归，也保证每个节点只克隆一次
	visited := make(map[*Node]*Node)
	var dfs func(cur *Node) *Node
	dfs = func(cur *Node) *Node {
		// 已克隆过的节点直接从表里取，环在这里被截断
		if cloned, ok := visited[cur]; ok {
			return cloned
		}
		// 先创建克隆节点并登记，再递归填充邻居（应对环）
		cloned := &Node{Val: cur.Val}
		visited[cur] = cloned
		for _, neighbor := range cur.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor))
		}
		return cloned
	}
	return dfs(node)
}
