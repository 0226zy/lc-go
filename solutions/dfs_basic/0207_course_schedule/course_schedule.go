package courseschedule

// CanFinish 课程表
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses-1。
// 先修关系 prerequisites[i] = [ai, bi] 表示想学课程 ai 必须先学课程 bi，
// 判断能否完成所有课程（等价于判断先修关系构成的有向图中是否存在环）。
// 时间复杂度: O(V+E) 每个节点和每条边至多访问一次  空间复杂度: O(V+E) 邻接表 + 状态数组 + 递归栈
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// 建图：bi -> ai，即“修完 bi 后可以修 ai”
	graph := make([][]int, numCourses)
	for _, pre := range prerequisites {
		graph[pre[1]] = append(graph[pre[1]], pre[0])
	}

	// 三色标记：0=未访问（白色），1=访问中（灰色，在当前递归栈上），2=已完成（黑色）
	state := make([]int, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if state[course] == 1 {
			// 遇到了当前递归栈上的节点，说明存在环
			return true
		}
		if state[course] == 2 {
			// 该节点的整条搜索子树已确认无环，直接剪枝
			return false
		}
		state[course] = 1 // 进入递归前标记为访问中
		for _, next := range graph[course] {
			if hasCycle(next) {
				return true
			}
		}
		state[course] = 2 // 所有后继检查完毕，标记为已完成（不回退到 0）
		return false
	}

	// 图可能不连通，对每个未访问的节点都要启动一次 DFS
	for i := 0; i < numCourses; i++ {
		if state[i] == 0 && hasCycle(i) {
			return false
		}
	}
	return true
}
