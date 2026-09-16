package findklengthsubstringswithnorepeatedcharacters

// NumKLenSubstrNoRepeats 长度为 K 的无重复字符子串
// 统计字符串 s 中长度为 k 且不含重复字符的子串个数。
// 时间复杂度: O(n)  空间复杂度: O(1)（字符集固定为 26 个小写字母）
func NumKLenSubstrNoRepeats(s string, k int) int {
	if k > len(s) {
		return 0
	}

	var freq [26]int
	dup := 0 // 窗口内出现次数 >= 2 的字符种数
	ans := 0
	for i := 0; i < len(s); i++ {
		// 右端字符进入窗口
		c := s[i] - 'a'
		freq[c]++
		if freq[c] == 2 {
			dup++
		}
		// 窗口已满后，左端字符滑出
		if i >= k {
			out := s[i-k] - 'a'
			if freq[out] == 2 {
				dup--
			}
			freq[out]--
		}
		// 窗口长度为 k 且无重复字符时计数
		if i >= k-1 && dup == 0 {
			ans++
		}
	}
	return ans
}
