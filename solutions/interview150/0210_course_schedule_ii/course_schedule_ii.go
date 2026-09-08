package coursescheduleii

// FindOrder 课程表 II
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses-1。
// 在选修某些课程之前需要一些先修课程，先修课程以数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] 表示如果要学习课程 ai 则必须先学习课程 bi。
// 返回一个合理的修课顺序，如果不可能完成所有课程，返回空数组。
// 时间复杂度: O(V+E) 每个节点和每条边只处理一次  空间复杂度: O(V+E) 邻接表 + 入度数组
func FindOrder(numCourses int, prerequisites [][]int) []int {
	// 建图：边 bi -> ai，并统计每门课的入度
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		ai, bi := pre[0], pre[1]
		graph[bi] = append(graph[bi], ai)
		indegree[ai]++
	}

	// Kahn 算法：把所有入度为 0 的课程（没有先修要求）入队
	queue := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 依次修完入度为 0 的课程，顺序即为拓扑序
	order := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)
		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 存在环时无法修完所有课程，返回空数组
	if len(order) != numCourses {
		return []int{}
	}
	return order
}
