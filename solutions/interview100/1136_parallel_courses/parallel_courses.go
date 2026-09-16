package parallelcourses

// MinimumSemesters 并行课程
// 有 n 门课程（编号 1..n），relations[j] = [prev, next] 表示 prev 是 next 的先修课。
// 每学期可以修读任意多门先修课均已完成的课程，返回学完全部课程的最少学期数；
// 若存在循环依赖导致无法学完，返回 -1。
// 时间复杂度: O(n + e)  空间复杂度: O(n + e)，e 为先修关系数
func MinimumSemesters(n int, relations [][]int) int {
	// 建图：邻接表 + 入度统计（节点编号 1..n）
	graph := make([][]int, n+1)
	indegree := make([]int, n+1)
	for _, r := range relations {
		prev, next := r[0], r[1]
		graph[prev] = append(graph[prev], next)
		indegree[next]++
	}

	// 入度为 0 的课程构成当前学期可修读的层
	queue := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	semesters, count := 0, 0
	// 分层 BFS（Kahn 拓扑排序）：每弹空一层即过一个学期
	for len(queue) > 0 {
		size := len(queue)
		semesters++
		count += size
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, next := range graph[cur] {
				indegree[next]--
				if indegree[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
	}

	// 未能修完全部课程说明有环
	if count != n {
		return -1
	}
	return semesters
}
