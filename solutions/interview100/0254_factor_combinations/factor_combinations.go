package factorcombinations

// GetFactors 因子的组合
// 给定整数 n，返回其所有因子组合（每个因子严格大于 1 且小于 n，组合内非递减、不重复）。
// 时间复杂度: 约 O(n log n) 上界，实际远小于此  空间复杂度: O(log n) 递归栈（不计结果集）
func GetFactors(n int) [][]int {
	result := [][]int{}
	backtrack(n, 2, []int{}, &result)
	return result
}

// backtrack 枚举 n 的因子 f（f >= start），收集组合并对商继续递归分解
func backtrack(n, start int, path []int, result *[][]int) {
	for f := start; f*f <= n; f++ {
		if n%f != 0 {
			continue
		}
		q := n / f
		// 收集「path + f + q」这种以 f 开头的分解
		combo := make([]int, len(path)+2)
		copy(combo, path)
		combo[len(path)] = f
		combo[len(path)+1] = q
		*result = append(*result, combo)

		// 以 f 为下界继续拆分 q，保证组合非递减、不重复
		backtrack(q, f, append(path, f), result)
	}
}
