package integerbreak

// IntegerBreak 整数拆分（标准 DP 数组版）
// 将 n 拆成至少两个正整数之和，求这些整数的最大乘积。
// dp[i] 表示拆分 i 所得的最大乘积；
// dp[i] = max(j*(i-j), j*dp[i-j])，j 为拆出的第一份，剩下的部分直接用或继续拆。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1 // base case：2 只能拆成 1+1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// 拆出第一份 j，剩下的 i-j 要么整体使用，要么继续拆
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// IntegerBreakAlternative 整数拆分（数学贪心版）
// 数学结论：每份不应大于 4，最优拆法是尽可能多地拆出 3；
// 余 1 时把一个 3 与 1 合并成 4（3×1 不如 2×2）。
// 时间复杂度: O(n)  空间复杂度: O(1)
func IntegerBreakAlternative(n int) int {
	if n <= 3 {
		return n - 1 // 2 → 1，3 → 2（题目要求至少拆成两份）
	}
	product := 1
	for n > 4 {
		product *= 3 // 每次拆出一个 3
		n -= 3
	}
	return product * n // 剩下的 n 必为 2、3 或 4，直接乘上即可
}
