package permutations

// Permute 全排列
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列，可以按任意顺序返回答案。
// 时间复杂度: O(n * n!)  空间复杂度: O(n) 递归栈深度（不计输出）
func Permute(nums []int) [][]int {
	n := len(nums)

	// 预分配 n! 的容量，避免 append 反复扩容
	capacity := 1
	for i := 2; i <= n; i++ {
		capacity *= i
	}
	result := make([][]int, 0, capacity)

	var backtrack func(first int)
	backtrack = func(first int) {
		// 结束条件：所有位置都已填完，收集一份当前排列的拷贝
		if first == n {
			perm := make([]int, n)
			copy(perm, nums)
			result = append(result, perm)
			return
		}
		// 选择列表：下标 [first, n) 中尚未固定的元素，依次交换到第 first 位
		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first] // 做选择
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first] // 撤销选择
		}
	}
	backtrack(0)
	return result
}
