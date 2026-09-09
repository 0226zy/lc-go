package coinchangeii

// Change 零钱兑换II（标准二维 DP 数组版）
// 每种硬币无限供应，求凑出总金额 amount 的组合数。
// dp[i][j] 表示只使用前 i 种硬币凑出金额 j 的组合数：
// dp[i][j] = dp[i-1][j]（不用第 i 种）+ dp[i][j-coins[i-1]]（至少用一枚第 i 种）。
// 时间复杂度: O(m*amount)  空间复杂度: O(m*amount)
func Change(amount int, coins []int) int {
	m := len(coins)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 // base case：凑金额 0 只有「一枚都不选」一种组合
	}
	for i := 1; i <= m; i++ {
		c := coins[i-1]
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j] // 不用第 i 种硬币
			if j >= c {
				dp[i][j] += dp[i][j-c] // 至少用一枚第 i 种硬币（完全背包，同层左侧）
			}
		}
	}
	return dp[m][amount]
}

// ChangeOptimized 零钱兑换II（一维滚动数组空间优化版）
// dp[j] 滚动表示使用到当前硬币种类时凑金额 j 的组合数。
// 硬币种类在外层保证按组合计数而非排列；金额正序允许重复使用当前硬币。
// 每个金额的计数依赖同层所有更小金额的历史值，空间最多压到 O(amount)。
// 时间复杂度: O(m*amount)  空间复杂度: O(amount)
func ChangeOptimized(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1 // base case
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			dp[j] += dp[j-c]
		}
	}
	return dp[amount]
}
