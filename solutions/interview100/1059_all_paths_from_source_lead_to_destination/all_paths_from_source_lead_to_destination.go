package allpathsfromsourceleadtodestination

// LeadsToDestination 从起点到终点的所有路径
// 给定 n 个节点的有向图（edges 为有向边列表）、起点 source 和终点 destination，
// 判断从 source 出发的所有路径是否都最终终止于 destination。
// 要求：终点出度为 0、无环可达、不存在终点以外的死胡同。
// 时间复杂度: O(n + e)  空间复杂度: O(n + e)，e 为边数
func LeadsToDestination(n int, edges [][]int, source int, destination int) bool {
	// 建邻接表
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	const (
		white = 0 // 未访问
		gray  = 1 // 在当前递归栈上（访问中）
		black = 2 // 已确认从该节点出发的所有路径都到达终点
	)
	color := make([]int, n)

	var dfs func(u int) bool
	dfs = func(u int) bool {
		// 遇到灰色节点说明存在环，环上的路径永远走不到终点
		if color[u] == gray {
			return false
		}
		// 黑色节点之前已验证安全，直接复用结论
		if color[u] == black {
			return true
		}
		// 出度为 0：路径在此终止，只有终点才合法
		if len(graph[u]) == 0 {
			return u == destination
		}
		color[u] = gray
		for _, v := range graph[u] {
			if !dfs(v) {
				return false
			}
		}
		color[u] = black
		return true
	}

	return dfs(source)
}
