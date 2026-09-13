package coinchange

// CoinChange 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount，计算凑成总金额所需的最少硬币个数。
// 如果无法凑成，返回 -1。每种硬币数量无限（完全背包）。
// dp[x] 表示凑成金额 x 所需的最少硬币数：dp[x] = min(dp[x-coin]+1)。
// 时间复杂度: O(amount * len(coins))  空间复杂度: O(amount)
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	inf := amount + 1 // 任何合法方案的硬币数都不会超过 amount
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = inf
	}
	dp[0] = 0

	for x := 1; x <= amount; x++ {
		for _, coin := range coins {
			if coin <= x && dp[x-coin]+1 < dp[x] {
				dp[x] = dp[x-coin] + 1
			}
		}
	}

	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}
