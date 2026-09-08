package combinations

// Combine 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 时间复杂度: O(C(n,k) * k)  空间复杂度: O(k) 递归栈深度
func Combine(n int, k int) [][]int {
	var result [][]int
	path := make([]int, 0, k)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			combo := make([]int, k)
			copy(combo, path)
			result = append(result, combo)
			return
		}
		// 剪枝：从 start 到 n 还剩 n-start+1 个数可选，
		// 如果不足以凑齐 k-len(path) 个，直接停止
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i + 1) // 从 i+1 开始，保证组合内升序、不重复
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return result
}
