package perfectsquares

import "math"

// NumSquares 完全平方数（标准 DP 数组版）
// 求和为 n 的完全平方数的最少数量（完全背包模型）。
// dp[i] 表示和为 i 所需的最少完全平方数个数，
// dp[i] = min(dp[i-j*j]) + 1，其中 j*j <= i。
// 时间复杂度: O(n*sqrt(n))  空间复杂度: O(n)
func NumSquares(n int) int {
	dp := make([]int, n+1) // dp[0] = 0 为 base case
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32 // 初始化为足够大的值，方便取 min
		for j := 1; j*j <= i; j++ {
			// 枚举最后一个平方数 j*j，剩余 i-j*j 用子问题的最优解
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

// NumSquaresAlternative 完全平方数（数学定理版）
// 拉格朗日四平方和定理：答案至多为 4；
// 勒让德三平方和定理：答案为 4 当且仅当 n = 4^a * (8b+7)。
// 依次判定：完全平方数→1；形如 4^a(8b+7)→4；两平方数之和→2；其余→3。
// 时间复杂度: O(sqrt(n))  空间复杂度: O(1)
func NumSquaresAlternative(n int) int {
	if isSquare(n) {
		return 1
	}
	m := n
	for m%4 == 0 { // 不断除掉因子 4，检查剩余部分是否形如 8b+7
		m /= 4
	}
	if m%8 == 7 {
		return 4
	}
	for i := 1; i*i <= n; i++ { // 检查 n - i*i 是否为完全平方数
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

// isSquare 判断 x 是否为完全平方数
func isSquare(x int) bool {
	r := int(math.Sqrt(float64(x)))
	return r*r == x
}
