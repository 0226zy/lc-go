package wordpattern

import "strings"

// WordPattern 单词规律
// 判断字符串 pattern 中的每个字符是否按固定的一对一映射与 str 中的单词一一对应。
// 时间复杂度: O(n)  空间复杂度: O(n)（需要存储切分出的单词）
func WordPattern(pattern string, s string) bool {
	// 注意：strings.Split("", " ") 会返回 [""]（长度 1），
	// 而空串按题意应视为 0 个单词，需要特殊处理
	var words []string
	if s != "" {
		words = strings.Split(s, " ")
	}
	if len(pattern) != len(words) {
		return false
	}
	// char2Word: pattern 中字符到单词的映射；word2Char: 单词到字符的反向映射，保证双射
	char2Word := make(map[byte]string, len(pattern))
	word2Char := make(map[string]byte, len(pattern))
	for i := 0; i < len(pattern); i++ {
		c, w := pattern[i], words[i]
		if mappedWord, ok := char2Word[c]; ok {
			if mappedWord != w {
				return false
			}
		} else if _, ok := word2Char[w]; ok {
			// 该单词已被其他字符占用，出现“多对一”冲突
			return false
		} else {
			char2Word[c] = w
			word2Char[w] = c
		}
	}
	return true
}
