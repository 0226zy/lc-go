package minimumwindowsubstring

// MinWindow 最小覆盖子串
// 给定两个字符串 s 和 t，返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""。
// need/have 计数 + 滑动窗口模板。
// 时间复杂度: O(len(s) + len(t))  空间复杂度: O(|Σ|) Σ 为字符集大小
func MinWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	// need 记录 t 中每个字符需要的次数
	// have 记录当前窗口中每个字符已有的次数
	// required 记录 need 中不同字符的个数；formed 记录已满足要求的字符个数
	var need, have [128]int
	for i := 0; i < n; i++ {
		need[t[i]]++
	}
	required := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			required++
		}
	}

	left, start, minLen := 0, 0, m+1
	formed := 0
	for right := 0; right < m; right++ {
		c := s[right]
		have[c]++
		// 字符 c 恰好达到所需次数，满足要求的字符种类 +1
		if need[c] > 0 && have[c] == need[c] {
			formed++
		}
		// 窗口已覆盖 t 中所有字符，尝试收缩左边界取更短的窗口
		for formed == required {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			d := s[left]
			have[d]--
			// 移出的字符 d 跌破所需次数，满足要求的字符种类 -1
			if need[d] > 0 && have[d] < need[d] {
				formed--
			}
			left++
		}
	}

	if minLen == m+1 {
		return ""
	}
	return s[start : start+minLen]
}
