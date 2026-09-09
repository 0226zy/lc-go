package coinchange

// CoinChange 零钱兑换（标准 DP 数组版）
// 每种硬币无限使用，求凑成 amount 所需的最少硬币个数，凑不出返回 -1。
// dp[i] 表示凑成金额 i 的最少硬币数，dp[i] = min(dp[i-c] + 1)，c 为不超过 i 的硬币面额。
// 时间复杂度: O(amount × len(coins))  空间复杂度: O(amount)
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		// 最多 amount 枚面额 1 的硬币即可凑出任何金额，
		// 因此 amount+1 大于所有合法答案，可安全当作无穷大
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i {
				// 最后一枚硬币面额为 c：此前凑 i-c 需要 dp[i-c] 枚
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1 // 仍为初始标记，说明凑不出
	}
	return dp[amount]
}

// CoinChangeAlternative 零钱兑换（BFS 版）
// 把金额看作图的节点，用一枚硬币从 x 走到 x-c；
// 从 amount 到 0 的最短路径长度即最少硬币数，BFS 逐层扩展第一次到达 0 时返回。
// 时间复杂度: O(amount × len(coins))  空间复杂度: O(amount)
func CoinChangeAlternative(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	visited := make([]bool, amount+1)
	queue := []int{amount}
	visited[amount] = true
	steps := 0
	for len(queue) > 0 {
		steps++ // 每扩展一层相当于多用一枚硬币
		next := []int{}
		for _, cur := range queue {
			for _, c := range coins {
				rem := cur - c
				if rem == 0 {
					return steps // 第一次到达 0，层数即最少硬币数
				}
				if rem > 0 && !visited[rem] {
					visited[rem] = true
					next = append(next, rem)
				}
			}
		}
		queue = next
	}
	return -1
}
