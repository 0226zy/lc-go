package minimumknightmoves

// MinKnightMoves 进击的骑士
// 在无限大的棋盘上，骑士（马走日）从 (0, 0) 出发，
// 返回到达 (x, y) 所需的最少移动次数。
// 思路：利用对称性把目标映射到第一象限，然后在有界区域内做 BFS。
// 时间复杂度: O(|x|·|y|)  空间复杂度: O(|x|·|y|)
func MinKnightMoves(x, y int) int {
	// 利用对称性，只需在第一象限搜索
	x, y = abs(x), abs(y)

	// 骑士的 8 个移动方向
	dirs := [][2]int{
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	}

	queue := [][2]int{{0, 0}}
	visited := map[[2]int]bool{{0, 0}: true}
	steps := 0
	for len(queue) > 0 {
		// 按层扩展，当前层所有位置都对应相同的步数
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur[0] == x && cur[1] == y {
				return steps
			}
			for _, d := range dirs {
				nx, ny := cur[0]+d[0], cur[1]+d[1]
				next := [2]int{nx, ny}
				// 最短路径不会偏离目标太远，限制在 [-2, x+2] × [-2, y+2] 内
				if nx < -2 || nx > x+2 || ny < -2 || ny > y+2 || visited[next] {
					continue
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}
		steps++
	}
	return -1 // 题目保证可达，不会走到这里
}

// abs 返回非负绝对值
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
