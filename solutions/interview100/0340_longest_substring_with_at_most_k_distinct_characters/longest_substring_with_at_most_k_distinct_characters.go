package longestsubstringwithatmostkdistinctcharacters

// LengthOfLongestSubstringKDistinct 至多包含 K 个不同字符的最长子串
// 给定一个字符串 s 和一个整数 k，返回 s 中至多包含 k 个不同字符的最长子串的长度。
// 时间复杂度: O(n) 每个字符最多进出窗口一次  空间复杂度: O(k) 哈希表最多存 k+1 个字符
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	// cnt 统计当前窗口 [left, right] 内每个字符的出现次数
	cnt := make(map[byte]int)
	left, ans := 0, 0
	for right := 0; right < len(s); right++ {
		cnt[s[right]]++
		// 窗口内不同字符超过 k 个时，从左端收缩直到恢复合法
		for len(cnt) > k {
			c := s[left]
			cnt[c]--
			if cnt[c] == 0 {
				delete(cnt, c)
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
