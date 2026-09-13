package subsetsii

import "sort"

// SubsetsWithDup 子集 II
// 给定一个可能包含重复元素的整数数组 nums，返回所有可能的子集（幂集），
// 解集不能包含重复的子集，可以按任意顺序返回。
// 时间复杂度: O(n × 2^n)  空间复杂度: O(n) 递归栈深度（不计输出）
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 排序使相同元素相邻，是同层去重的前提
	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// 每个节点都是一个合法子集，进入递归先收集当前路径的拷贝
		combo := make([]int, len(path))
		copy(combo, path)
		result = append(result, combo)

		for i := start; i < len(nums); i++ {
			// 同层去重：同一递归层中，与前一个元素相同的分支会产生完全一致的子树，跳过
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			backtrack(i + 1) // 传 i+1：每个元素至多使用一次
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
