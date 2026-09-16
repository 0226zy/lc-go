package shortestdistancefromallbuildings

import "math"

// ShortestDistance 离建筑物最近的距离
// 在网格的空地（0）上建一栋房子，使得到所有建筑物（1）的旅行距离总和最小；
// 障碍物（2）不可通过。返回最小距离总和，无法建造时返回 -1。
// 思路：从每栋建筑物出发做多源 BFS，累加每块空地到各建筑物的距离并统计可达建筑物数。
// 时间复杂度: O(b·m·n) b 为建筑物数量  空间复杂度: O(m·n)
func ShortestDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// distSum[i][j] 记录空地 (i,j) 到所有可达建筑物的距离之和
	// reach[i][j] 记录能到达空地 (i,j) 的建筑物数量
	distSum := make([][]int, m)
	reach := make([][]int, m)
	for i := range distSum {
		distSum[i] = make([]int, n)
		reach[i] = make([]int, n)
	}

	// 统计建筑物总数
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				total++
			}
		}
	}

	// bfs 从建筑物 (si, sj) 出发，把距离累加到所有可达空地
	bfs := func(si, sj int) {
		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		visited[si][sj] = true
		type point struct{ x, y int }
		queue := []point{{si, sj}}
		steps := 0
		for len(queue) > 0 {
			steps++
			size := len(queue)
			// 逐层扩展，同一层的节点到源点距离相同
			for s := 0; s < size; s++ {
				cur := queue[0]
				queue = queue[1:]
				for _, d := range dirs {
					x, y := cur.x+d[0], cur.y+d[1]
					if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
						continue
					}
					visited[x][y] = true
					// 只在空地上累加距离并继续扩展；遇到其他建筑物或障碍物则停止该方向
					if grid[x][y] == 0 {
						distSum[x][y] += steps
						reach[x][y]++
						queue = append(queue, point{x, y})
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	// 只有能到达全部建筑物的空地才有资格成为答案
	ans := math.MaxInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 && reach[i][j] == total && distSum[i][j] < ans {
				ans = distSum[i][j]
			}
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
