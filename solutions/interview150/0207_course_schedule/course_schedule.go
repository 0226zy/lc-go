package courseschedule

// CanFinish 课程表
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses-1。
// 在选修某些课程之前需要一些先修课程，先修课程以数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] 表示如果要学习课程 ai 则必须先学习课程 bi。
// 请你判断是否可能完成所有课程的学习（即图中是否有环）。
// 时间复杂度: O(V+E) 每个节点和每条边只处理一次  空间复杂度: O(V+E) 邻接表 + 入度数组
func CanFinish(numCourses int, prerequisites [][]int) bool {
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

	// 依次修完入度为 0 的课程，并解除后续课程的入度
	finished := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		finished++
		for _, next := range graph[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	// 能修完所有课程说明无环，否则存在环
	return finished == numCourses
}

// CanFinishDFS 课程表（DFS 三色标记解法）
// 0=未访问（白色），1=访问中（灰色，在递归栈上），2=已完成（黑色）。
// 一旦遇到灰色节点说明碰到了环。
// 时间复杂度: O(V+E)  空间复杂度: O(V+E)
func CanFinishDFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}
	state := make([]int, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if state[course] == 1 {
			// 遇到了访问中的节点，存在环
			return true
		}
		if state[course] == 2 {
			return false
		}
		state[course] = 1 // 标记为访问中
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		state[course] = 2 // 本节点及其后继全部检查完毕
		return false
	}

	for i := 0; i < numCourses; i++ {
		if state[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}
