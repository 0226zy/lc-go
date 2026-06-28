package courseschedule

// CanFinish 课程表
// 判断是否可以完成所有课程的学习（检测有向图是否存在环）。
// 时间复杂度: O(V+E)  空间复杂度: O(V+E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make(map[int]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	count := 0
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		count++
		for _, next := range graph[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	return count == numCourses
}
