package numberofdistinctislands

import "strings"

// NumDistinctIslands 不同岛屿的数量
// 给定 m x n 的 0/1 网格，统计形状不同（只允许平移、不允许旋转/翻转）的岛屿数量。
// 注意：函数会原地修改 grid（把访问过的陆地标记为 2）。
// 时间复杂度: O(m*n)  空间复杂度: O(m*n)
func NumDistinctIslands(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	// 四个方向：上、下、左、右，序列化时分别记录为 '1'~'4'
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	shapes := make(map[string]struct{})

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			// 对一个岛屿做 DFS，把搜索路径序列化为形状指纹
			var sb strings.Builder
			var dfs func(x, y int)
			dfs = func(x, y int) {
				grid[x][y] = 2 // 原地标记已访问，避免重复使用 visited 数组
				for d, dir := range dirs {
					nx, ny := x+dir[0], y+dir[1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == 1 {
						sb.WriteByte(byte('1' + d))
						dfs(nx, ny)
						// 回溯标记：缺少它会让不同形状产生相同序列
						sb.WriteByte('0')
					}
				}
			}
			dfs(i, j)
			shapes[sb.String()] = struct{}{}
		}
	}
	return len(shapes)
}
