package combinationsum

import "sort"

// CombinationSum 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个目标数 target，
// 找出所有数字和为 target 的不同组合，candidates 中的元素可以被无限次重复选取。
// 时间复杂度: O(2^t) t 为 target 量级，最坏为指数级  空间复杂度: O(t) 递归栈深度（不计输出）
func CombinationSum(candidates []int, target int) [][]int {
	// 拷贝后排序：利用单调性剪枝，同时避免修改调用方的原始切片
	sorted := make([]int, len(candidates))
	copy(sorted, candidates)
	sort.Ints(sorted)

	var ans [][]int
	path := make([]int, 0, len(sorted))

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			// 收集答案时必须拷贝，path 后续还会被修改
			combo := make([]int, len(path))
			copy(combo, path)
			ans = append(ans, combo)
			return
		}
		for i := start; i < len(sorted); i++ {
			if sum+sorted[i] > target {
				break // 已排序，后续元素更大，全部剪掉
			}
			path = append(path, sorted[i])
			backtrack(i, sum+sorted[i]) // 传 i：当前元素允许重复选取
			path = path[:len(path)-1]   // 撤销选择
		}
	}
	backtrack(0, 0)
	return ans
}
