package allpaths

// AllPathsSourceTarget 所有可能的路径
// 给定一个有向无环图（DAG），节点编号为 0 到 n-1，graph[i] 存储节点 i 指向的所有节点。
// 找出所有从节点 0 出发到达节点 n-1 的路径，可以按任意顺序返回。
// 时间复杂度: O(2^V * V) 最坏情况下路径总数为指数级，每条路径长度为 O(V)
// 空间复杂度: O(V) 递归栈深度与路径长度（不计输出）
func AllPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var result [][]int
	path := []int{0} // 路径从节点 0 开始

	var dfs func(node int)
	dfs = func(node int) {
		// 到达终点 n-1，收集当前路径的拷贝
		if node == n-1 {
			p := make([]int, len(path))
			copy(p, path)
			result = append(result, p)
			return
		}
		// 依次尝试当前节点的所有后继
		for _, next := range graph[node] {
			path = append(path, next)
			dfs(next)
			path = path[:len(path)-1] // 回溯：撤销选择
		}
	}
	dfs(0)
	return result
}
