package subsets

// Subsets 子集
// 给定一个元素互不相同的整数数组 nums，返回其所有可能的子集（幂集），解集中不能包含重复的子集。
// 时间复杂度: O(n × 2^n)  空间复杂度: O(n) 递归栈深度（不计输出）
func Subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// 每个递归节点的路径都是一个合法子集，进入即收集
		subset := make([]int, len(path))
		copy(subset, path)
		result = append(result, subset)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1) // 传 i+1：每个元素只选一次，避免重复组合
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
