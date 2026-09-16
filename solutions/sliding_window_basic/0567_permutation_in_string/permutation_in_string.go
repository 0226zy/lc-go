package permutationinstring

// CheckInclusion 字符串的排列
// 给定两个字符串 s1 和 s2，判断 s2 是否包含 s1 的某个排列（即 s1 的排列之一是 s2 的子串）。
// 思路：固定长度滑动窗口，s1 的任意排列与 s1 的字符计数相同，
// 因此只需比较 s2 中长度为 len(s1) 的窗口内字符计数与 s1 的字符计数。
// 时间复杂度: O(len(s2)) 每次比较长度为 26 的数组是常数开销  空间复杂度: O(1) 仅使用两个长度为 26 的计数数组
func CheckInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	// s1 比 s2 还长，不可能存在排列子串
	if m > n {
		return false
	}

	// need 记录 s1 中每个字符的出现次数
	// have 记录 s2 当前窗口（长度 m）中每个字符的出现次数
	var need, have [26]int
	for i := 0; i < m; i++ {
		need[s1[i]-'a']++
		have[s2[i]-'a']++
	}

	// 滑动窗口：右端进一个字符，左端出一个字符，每次比较计数
	for i := m; ; i++ {
		// 计数完全相等，说明当前窗口是 s1 的一个排列
		if need == have {
			return true
		}
		if i == n {
			break
		}
		// 右端字符 s2[i] 进入窗口，左端字符 s2[i-m] 移出窗口
		have[s2[i]-'a']++
		have[s2[i-m]-'a']--
	}
	return false
}
