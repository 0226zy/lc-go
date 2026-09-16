package substringwithconcatenationofallwords

// FindSubstring 串联所有单词的子串
// 给定一个字符串 s 和一个字符串数组 words，words 中所有单词长度相同。
// s 中的串联子串是指一个包含 words 中所有单词（顺序任意）、以任意顺序拼接起来的子串。
// 返回所有串联子串在 s 中的开始下标。题目数据保证 words 中的单词不互为同义词
// （即 words 中可以有重复单词，这些重复单词都要按出现次数使用）。
// 时间复杂度: O(len(s) * wordLen)  空间复杂度: O(len(words) * wordLen)
func FindSubstring(s string, words []string) []int {
	n := len(s)
	wordLen := len(words[0])
	wordCount := len(words)
	if n < wordLen*wordCount {
		return nil
	}

	// need 记录每个单词需要的次数
	need := make(map[string]int, wordCount)
	for _, w := range words {
		need[w]++
	}

	ans := []int{}
	// 按单词长度 wordLen 分组滑动窗口：起点偏移 i 取 0 ~ wordLen-1，
	// 保证同一组内所有切分点对齐，不会错位切词
	for i := 0; i < wordLen; i++ {
		left := i  // 窗口左边界（下标）
		count := 0 // 窗口内有效单词个数
		have := make(map[string]int, wordCount)
		for j := i; j+wordLen <= n; j += wordLen {
			w := s[j : j+wordLen]
			if need[w] > 0 {
				// 单词有效，加入窗口
				have[w]++
				count++
				// 若该单词出现次数超标，从左边收缩，直到次数达标
				for have[w] > need[w] {
					lw := s[left : left+wordLen]
					have[lw]--
					left += wordLen
					count--
				}
				if count == wordCount {
					ans = append(ans, left)
				}
			} else {
				// 切出的片段不在 words 中，窗口直接清空，从下一片段重新开始
				have = make(map[string]int, wordCount)
				count = 0
				left = j + wordLen
			}
		}
	}
	return ans
}
