package perfectsquares

import "math"

// NumSquares 完全平方数（动态规划·完全背包）
// 给定正整数 n，返回和为 n 的完全平方数的最少数量。
// 时间复杂度: O(n√n)  空间复杂度: O(n)
func NumSquares(n int) int {
	// dp[i] 表示和为 i 的最少完全平方数个数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i // 最坏情况：全用 1，需要 i 个
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j]+1 < dp[i] {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

// NumSquaresMath 完全平方数（数学法·拉格朗日四平方和定理）
// 利用四平方和定理与三平方定理，答案只可能是 1/2/3/4。
// 时间复杂度: O(√n)  空间复杂度: O(1)
func NumSquaresMath(n int) int {
	// 1. 判断是否为完全平方数 → 答案 1
	if isPerfectSquare(n) {
		return 1
	}

	// 2. 判断是否能写成两个平方数之和 → 答案 2
	for i := 1; i*i <= n; i++ {
		if isPerfectSquare(n - i*i) {
			return 2
		}
	}

	// 3. 勒让日三平方定理：n = 4^a * (8b+7) 时答案为 4，否则为 3
	//    先不断除 4 去掉 4^a 因子
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	return 3
}

// NumSquaresBFS 完全平方数（BFS 最短路径）
// 将问题建模为从 0 到 n 的最短路径，每条边代表加一个完全平方数。
// 时间复杂度: O(n√n)  空间复杂度: O(n)
func NumSquaresBFS(n int) int {
	// 预计算所有不超过 n 的完全平方数
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	visited := make([]bool, n+1)
	queue := []int{0}
	visited[0] = true
	level := 0

	for len(queue) > 0 {
		level++
		size := len(queue)
		for k := 0; k < size; k++ {
			cur := queue[k]
			for _, s := range squares {
				next := cur + s
				if next == n {
					return level
				}
				if next > n || visited[next] {
					continue
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}
		queue = queue[size:]
	}
	return level
}

// isPerfectSquare 判断 n 是否为完全平方数
func isPerfectSquare(n int) bool {
	s := int(math.Sqrt(float64(n)))
	return s*s == n
}
