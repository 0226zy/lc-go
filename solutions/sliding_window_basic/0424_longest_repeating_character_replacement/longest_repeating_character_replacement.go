package longestrepeatingcharacterreplacement

// LongestRepeatingCharacterReplacement 替换后的最长重复字符
// 给定字符串 s 和整数 k，最多可将 k 个字符替换为任意大写英文字母，
// 返回替换后包含相同字母的最长子串的长度。
// 思路：滑动窗口。窗口内出现次数最多的字符数为 maxCount，
// 当 窗口长度 - maxCount > k 时收缩左边界。
// 时间复杂度: O(n)  空间复杂度: O(1)（计数数组大小固定为 26）
func LongestRepeatingCharacterReplacement(s string, k int) int {
	var count [26]int
	left, maxCount, ans := 0, 0, 0
	for right := 0; right < len(s); right++ {
		// 右端字符加入窗口，更新窗口内最大众数频次
		count[s[right]-'A']++
		if count[s[right]-'A'] > maxCount {
			maxCount = count[s[right]-'A']
		}
		// 需要替换的次数（窗口长度 - maxCount）超过 k，收缩左边界
		if right-left+1-maxCount > k {
			count[s[left]-'A']--
			left++
		}
		// 此时窗口一定合法，用窗口长度更新答案
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
