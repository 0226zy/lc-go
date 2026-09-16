package minimumwindowsubstring

// MinWindow 最小覆盖子串
// 给定两个字符串 s 和 t，返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""。
// need/have 计数 + 滑动窗口模板。
// 时间复杂度: O(len(s) + len(t))  空间复杂度: O(|Σ|) Σ 为字符集大小
func MinWindow(s string, t string) string {
	m, n := len(s), len(t)
	if n > m {
		return ""
	}
	left, right := 0, 0
	need := [128]int{}
	for i := 0; i < n; i++ {
		need[t[i]]++
	}
	needNum := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			needNum++
		}
	}
	findNum := 0
	find := [128]int{}
	start, minLen := 0, len(s)+1

	for right < m {
		c := s[right]
		find[c]++
		if need[c] > 0 && find[c] == need[c] {
			findNum++
		}

		for needNum == findNum {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			c = s[left]
			find[c]--
			if need[c] > 0 && find[c] != need[c] {
				findNum--
			}
			left++
		}
		right++
	}
	if minLen == m+1 {
		return ""
	}
	return s[start : start+minLen]
}
