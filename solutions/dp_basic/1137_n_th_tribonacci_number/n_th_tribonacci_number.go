package nthtribonumber

// Tribonacci 第 N 个泰波那契数（标准 DP 数组版）
// 泰波那契数列 T0 = 0, T1 = 1, T2 = 1，Tn = Tn-1 + Tn-2 + Tn-3 (n >= 3)。
// dp[i] 表示第 i 个泰波那契数，dp[i] = dp[i-1] + dp[i-2] + dp[i-3]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n < 3 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1 // base case：题目直接给定
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3] // 第 i 项为前三项之和
	}
	return dp[n]
}

// TribonacciOptimized 第 N 个泰波那契数（滚动变量空间优化版）
// dp[i] 只依赖前三个状态，用三个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func TribonacciOptimized(n int) int {
	if n < 2 {
		return n
	}
	if n < 3 {
		return 1
	}
	a, b, c := 0, 1, 1 // 分别代表 dp[i-3]、dp[i-2]、dp[i-1]
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
