package combinations

// Combine 组合
// 给定两个整数 n 和 k，返回 [1, n] 范围内所有由 k 个数组成的组合，组合可按任意顺序返回。
// 时间复杂度: O(C(n,k) * k)，共 C(n,k) 个组合，每个组合需要 O(k) 拷贝  空间复杂度: O(k) 递归栈深度
func Combine(n int, k int) [][]int {
	result := [][]int{}
	path := make([]int, 0, k)

	var dfs func(start int)
	dfs = func(start int) {
		// 结束条件：路径长度达到 k，收集一份拷贝
		if len(path) == k {
			ans := make([]int, k)
			copy(ans, path)
			result = append(result, ans)
			return
		}
		// 剪枝：还需要 k-len(path) 个数，可选区间为 [start, n-(k-len(path))+1]，
		// 再往后剩余的数不够凑齐 k 个，直接跳过
		upper := n - (k - len(path)) + 1
		for i := start; i <= upper; i++ {
			path = append(path, i)    // 做选择
			dfs(i + 1)                // 下一层从 i+1 起选，保证组合内递增、不重复
			path = path[:len(path)-1] // 撤销选择
		}
	}
	dfs(1)
	return result
}
