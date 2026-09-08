package combinationsum

import "sort"

// CombinationSum 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个目标数 target，
// 找出 candidates 中所有可以使数字和为 target 的组合，同一元素可以重复选取。
// 时间复杂度: O(2^t)  t 为 target 值，最坏为指数级  空间复杂度: O(t) 递归栈深度
func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序后可利用单调性提前剪枝
	var result [][]int
	path := []int{}

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			combo := make([]int, len(path))
			copy(combo, path)
			result = append(result, combo)
			return
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break // candidates 已排序，后面的更大，直接剪掉
			}
			path = append(path, candidates[i])
			backtrack(i, sum+candidates[i]) // 传 i 而非 i+1：允许重复选取当前元素
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
