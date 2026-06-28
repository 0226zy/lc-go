package minimumwindowsubstring

// MinWindow 最小覆盖子串
// 返回 s 中涵盖 t 所有字符的最小子串。
// 思路：滑动窗口，right 右扩到覆盖 t，left 左缩到最短，记录最小窗口。
// 时间复杂度: O(m+n)  空间复杂度: O(1)（字符集固定）
func MinWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0 // 已满足需求的字符种类数
	start, minLen := 0, len(s)+1

	for right < len(s) {
		c := s[right] // 即将加入窗口的字符
		right++

		// 右扩：c 加入窗口
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++ // 该字符刚好满足需求
			}
		}

		// 左缩：当窗口覆盖 t 时，尝试收缩
		for valid == len(need) {
			// 更新最短
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			d := s[left] // 即将移出窗口的字符
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid-- // 移除后该字符不再满足
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
