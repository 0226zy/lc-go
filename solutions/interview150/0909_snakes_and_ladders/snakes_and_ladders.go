package snakesandladders

// SnakesAndLadders 蛇梯棋
// 在 n x n 棋盘上，方格按“从底行开始、奇数行从右往左”的蛇形编号 1 ~ n*n。
// 每步可以掷出 1~6，若落脚格有蛇或梯子则立即传送到目标格，求到达 n*n 的最少移动次数，无法到达返回 -1。
// 时间复杂度: O(n^2) 每个格子最多入队一次  空间复杂度: O(n^2) visited 数组与队列
func SnakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n

	visited := make([]bool, target+1)
	visited[1] = true
	queue := []int{1}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == target {
				return steps
			}
			// 尝试掷出 1~6
			for next := cur + 1; next <= cur+6 && next <= target; next++ {
				r, c := squareToCoord(next, n)
				dest := next
				if board[r][c] != -1 {
					dest = board[r][c]
				}
				if !visited[dest] {
					visited[dest] = true
					queue = append(queue, dest)
				}
			}
		}
		steps++
	}
	return -1
}

// squareToCoord 将方格编号转为棋盘坐标
// 编号从棋盘最后一行（底行）开始，底行为 1~n，倒数第二行为 n+1~2n（从右往左），依此蛇形交替
func squareToCoord(pos, n int) (int, int) {
	pos-- // 转成 0 基编号
	r := n - 1 - pos/n
	c := pos % n
	if (n-1-r)%2 == 1 {
		// 从底数奇数行：编号从右往左
		c = n - 1 - c
	}
	return r, c
}
