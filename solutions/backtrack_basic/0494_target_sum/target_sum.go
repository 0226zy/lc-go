package targetsum

// FindTargetSumWays 目标和
// 给整数数组 nums 中的每个数添加 '+' 或 '-' 符号，统计有多少种符号组合
// 能使表达式结果等于 target。
// 时间复杂度: O(n*S)  n 为数组长度，S 为搜索过程中可达和的不同取值个数（最坏约 2*1000*n 级别）
// 空间复杂度: O(n*S)  记忆化表开销 + O(n) 递归栈深度
func FindTargetSumWays(nums []int, target int) int {
	// memo[(i, sum)] 记录“从下标 i 开始处理、当前累计和为 sum”时的方案数
	type state struct{ i, sum int }
	memo := make(map[state]int)

	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {
		// 结束条件：所有数都已确定符号，累计和等于 target 则计 1 种方案
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		key := state{i, sum}
		if v, ok := memo[key]; ok {
			return v
		}
		// 选择列表：对 nums[i] 分别尝试加号与减号，无需显式还原（sum 通过参数传递）
		ways := backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		memo[key] = ways
		return ways
	}
	return backtrack(0, 0)
}
