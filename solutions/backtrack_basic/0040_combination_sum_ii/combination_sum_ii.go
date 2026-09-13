package combinationsumii

import "sort"

// CombinationSum2 组合总和 II
// 给定一个可包含重复元素的整数数组 candidates 和一个目标数 target，
// 找出所有和为 target 的不同组合，每个元素在一个组合中只能使用一次。
// 时间复杂度: O(2^n)  n 为数组长度，最坏为指数级  空间复杂度: O(n) 递归栈深度
func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序：既能让重复元素相邻便于去重，也能利用单调性剪枝
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
			if i > start && candidates[i] == candidates[i-1] {
				continue // 同层去重：同一层的相同值只尝试第一次，避免产生重复组合
			}
			path = append(path, candidates[i])
			backtrack(i+1, sum+candidates[i]) // 传 i+1：每个元素在一个组合中只能用一次
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
