package permutationsii

import "sort"

// PermuteUnique 全排列 II
// 给定一个可包含重复数字的序列 nums，返回其所有不重复的全排列，可以按任意顺序返回答案。
// 时间复杂度: O(n * n!)  空间复杂度: O(n) 递归栈深度（不计输出）
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 排序：让重复元素相邻，便于同层去重
	n := len(nums)
	var result [][]int
	used := make([]bool, n) // used[i] 表示 nums[i] 是否已在当前路径中
	path := make([]int, 0, n)

	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			p := make([]int, n)
			copy(p, path)
			result = append(result, p)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue // 同一排列中每个元素只能用一次
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue // 同层去重：相同的值只在前一个相同值已使用时才选，保证同一层的相同值只展开一次
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}
