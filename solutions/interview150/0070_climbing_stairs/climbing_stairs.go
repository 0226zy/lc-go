package climbingstairs

// ClimbStairs 爬楼梯
// 爬 n 阶楼梯，每次可以爬 1 或 2 个台阶，返回爬到楼顶的不同方法总数。
// 本质是斐波那契数列：f(n) = f(n-1) + f(n-2)，f(1) = 1，f(2) = 2。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ClimbStairs(n int) int {
	// n <= 2 时方法数就是 n 本身：f(1) = 1，f(2) = 2
	if n <= 2 {
		return n
	}

	// 滚动变量：prev2 表示 f(i-2)，prev1 表示 f(i-1)
	prev2, prev1 := 1, 2
	for i := 3; i <= n; i++ {
		// 爬到第 i 阶 = 从第 i-1 阶迈 1 步 + 从第 i-2 阶迈 2 步
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
