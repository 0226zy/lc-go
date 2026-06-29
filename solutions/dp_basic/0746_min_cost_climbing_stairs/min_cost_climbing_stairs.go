package mincostclimbingstairs

// MinCostClimbingStairs 使用最小花费爬楼梯
// 每次可以爬 1 或 2 个台阶，cost[i] 是从第 i 级出发的花费，求到达顶部的最小花费。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	prev, curr := 0, 0 // f[0]=0, f[1]=0
	for i := 2; i <= n; i++ {
		prev, curr = curr, min(curr+cost[i-1], prev+cost[i-2])
	}
	return curr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
