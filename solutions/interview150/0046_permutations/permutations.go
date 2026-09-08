package permutations

// Permute 全排列
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列，可以按任意顺序返回答案。
// 时间复杂度: O(n * n!)  空间复杂度: O(n) 递归栈深度
func Permute(nums []int) [][]int {
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
				continue // 同一排列中每个数只能用一次
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
