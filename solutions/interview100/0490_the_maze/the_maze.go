package themaze

// HasPath 迷宫
// 球在迷宫中沿四个方向滚动，撞到墙壁（或边界）才停下，停下后才能换方向。
// 判断球能否从 start 恰好停在 destination。maze 中 0 为空地，1 为墙壁。
// 时间复杂度: O(m*n*max(m,n)) 每个停靠点向四个方向滚动到底  空间复杂度: O(m*n) 访问标记与队列
func HasPath(maze [][]int, start []int, destination []int) bool {
	m, n := len(maze), len(maze[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// BFS：队列中只放球能停下的位置
	queue := [][2]int{{start[0], start[1]}}
	visited[start[0]][start[1]] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur[0] == destination[0] && cur[1] == destination[1] {
			return true
		}
		for _, d := range dirs {
			// 沿当前方向一直滚到下一格是墙壁或边界为止
			nx, ny := cur[0], cur[1]
			for {
				tx, ty := nx+d[0], ny+d[1]
				if tx < 0 || tx >= m || ty < 0 || ty >= n || maze[tx][ty] == 1 {
					break
				}
				nx, ny = tx, ty
			}
			if !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}
	return false
}
