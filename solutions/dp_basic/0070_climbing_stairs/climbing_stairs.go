package climbingstairs

// ClimbStairs 爬楼梯
// 每次可以爬 1 或 2 个台阶，求爬到第 n 阶的方法数。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
