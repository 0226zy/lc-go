package fibonaccinumber

// Fib 斐波那契数（标准 DP 数组版）
// F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2)，求第 n 个斐波那契数。
// dp[i] 表示 F(i)，dp[i] = dp[i-1] + dp[i-2]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1 // base case
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 前两项之和
	}
	return dp[n]
}

// FibOptimized 斐波那契数（滚动变量空间优化版）
// dp[i] 只依赖前两个状态，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func FibOptimized(n int) int {
	if n <= 1 {
		return n
	}
	prev2, prev1 := 0, 1 // 分别代表 dp[i-2]、dp[i-1]
	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
