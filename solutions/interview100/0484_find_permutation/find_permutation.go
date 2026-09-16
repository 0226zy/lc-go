package findpermutation

// FindPermutation 寻找排列
// 给定只含 'I'（递增）和 'D'（递减）的字符串 s，返回满足：
// s[i]=='I' 时 perm[i]<perm[i+1]，s[i]=='D' 时 perm[i]>perm[i+1]
// 的 [1..len(s)+1] 的字典序最小排列。
// 时间复杂度: O(n) 每个元素最多被反转一次  空间复杂度: O(n) 存放结果排列
func FindPermutation(s string) []int {
	n := len(s) + 1
	res := make([]int, n)
	// 恒等排列天然满足所有 'I'，只需修正连续 'D' 段
	for i := range res {
		res[i] = i + 1
	}
	for i := 0; i < len(s); {
		if s[i] == 'I' {
			i++
			continue
		}
		// 定位连续 'D' 段 [i, j-1]，对应的排列区间为 [i, j]（长度多 1）
		j := i
		for j < len(s) && s[j] == 'D' {
			j++
		}
		// 整体反转该区间，使其严格递减且字典序最小
		for l, r := i, j; l < r; l, r = l+1, r-1 {
			res[l], res[r] = res[r], res[l]
		}
		i = j
	}
	return res
}
