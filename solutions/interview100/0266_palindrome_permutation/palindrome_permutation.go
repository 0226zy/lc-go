package palindromepermutation

// CanPermutePalindrome 回文排列
// 给定一个字符串 s，判断能否通过重新排列其中的字符得到一个回文串。
// 充要条件：出现次数为奇数的字符至多有一个（它可以放在回文正中间）。
// 时间复杂度: O(n)  空间复杂度: O(|Σ|) Σ 为字符集大小，小写字母场景下为 O(1)
func CanPermutePalindrome(s string) bool {
	// 统计每个字符出现的次数
	cnt := make(map[rune]int)
	for _, c := range s {
		cnt[c]++
	}
	// 出现奇数次的字符至多允许一个
	odd := 0
	for _, v := range cnt {
		if v%2 == 1 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}
