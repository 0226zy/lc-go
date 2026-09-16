package wallsandgates

// inf 表示空房间（到门的距离尚未计算）
const inf = 1<<31 - 1

// WallsAndGates 墙与门
// rooms 中 -1 表示墙，0 表示门，INF 表示空房间。
// 多源 BFS：把所有门同时入队逐层扩散，为每个空房间填入到最近门的距离；不可达的房间保持 INF。
// 结果就地修改 rooms。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func WallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])

	// 所有门作为 BFS 源点入队
	queue := make([][2]int, 0)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if rooms[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			// 越界、是墙、或已被更近的门访问过（值不再是 INF）时跳过
			if nr < 0 || nr >= m || nc < 0 || nc >= n || rooms[nr][nc] != inf {
				continue
			}
			rooms[nr][nc] = rooms[r][c] + 1
			queue = append(queue, [2]int{nr, nc})
		}
	}
}
