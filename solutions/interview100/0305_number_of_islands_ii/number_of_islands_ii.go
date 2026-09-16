package numberofislandsii

import "github.com/0226zy/lc-go/pkg/datastructures"

// NumIslands2 岛屿数量 II
// m x n 的网格初始全为水，依次执行 positions 中的填海操作，
// 返回每次操作后网格中的岛屿数量（四方向连通）。
// 思路：并查集维护动态连通性，新增陆地先 +1，每与一个已连通分量合并成功就 -1。
// 时间复杂度: O(k * α(mn))，近似 O(k)，k 为操作次数  空间复杂度: O(m*n)
func NumIslands2(m int, n int, positions [][]int) []int {
	uf := datastructures.NewUnionFind(m * n)
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	ans := make([]int, 0, len(positions))
	count := 0 // 当前岛屿数量（初始全为水，不从 uf.Count 读取）
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, pos := range positions {
		r, c := pos[0], pos[1]
		// 重复填海：状态不变，直接记录当前岛屿数
		if grid[r][c] == 1 {
			ans = append(ans, count)
			continue
		}
		grid[r][c] = 1
		count++ // 先视为一个孤立岛屿

		// 与四周已存在的陆地合并，每成功合并一个连通分量，岛屿数减一
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || grid[nr][nc] == 0 {
				continue
			}
			if uf.Union(r*n+c, nr*n+nc) {
				count--
			}
		}
		ans = append(ans, count)
	}
	return ans
}
