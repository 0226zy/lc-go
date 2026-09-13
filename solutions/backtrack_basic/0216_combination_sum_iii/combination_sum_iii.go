package combinationsumiii

// CombinationSum3 组合总和 III
// 从 1 到 9 中选取 k 个互不相同的数字，使其和为 n，返回所有满足条件的组合。
// 时间复杂度: O(C(9,k))，即从 9 个数中选 k 个的组合数  空间复杂度: O(k) 递归栈深度
func CombinationSum3(k int, n int) [][]int {
	var result [][]int
	path := []int{}

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		// 路径长度达到 k 时检查是否满足和为 n
		if len(path) == k {
			if sum == n {
				combo := make([]int, len(path))
				copy(combo, path)
				result = append(result, combo)
			}
			return
		}
		// 剩余位置数
		remain := k - len(path)
		for i := start; i <= 9; i++ {
			// 剪枝一：加入 i 后和已经超过 n，后面更大，直接剪掉
			if sum+i > n {
				break
			}
			// 剪枝二：剩余数字不足以凑满 k 个，后面没必要继续
			if 9-i+1 < remain {
				break
			}
			path = append(path, i)
			backtrack(i+1, sum+i) // 传 i+1：每个数字最多使用一次
			path = path[:len(path)-1]
		}
	}
	backtrack(1, 0)
	return result
}
