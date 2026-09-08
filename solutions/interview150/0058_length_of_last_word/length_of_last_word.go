package lengthoflastword

import "strings"

// LengthOfLastWord 最后一个单词的长度
// 给定一个字符串 s，由若干单词组成，单词之间用空格分隔，返回最后一个单词的长度。
// 时间复杂度: O(n)  空间复杂度: O(1)（标准库 TrimRight 内部不分配新切片）
func LengthOfLastWord(s string) int {
	// 从后往前跳过所有尾部空格
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	end := i // 最后一个单词的结尾

	// 继续向前数非空格字符的个数
	for i >= 0 && s[i] != ' ' {
		i--
	}
	return end - i
}

// LengthOfLastWordTrim 标准库实现：先 Trim 掉尾部空格，再定位最后一个空格
// 时间复杂度: O(n)  空间复杂度: O(n)（strings.TrimRight 会创建子切片副本）
func LengthOfLastWordTrim(s string) int {
	s = strings.TrimRight(s, " ")
	if idx := strings.LastIndex(s, " "); idx >= 0 {
		return len(s) - idx - 1
	}
	return len(s)
}
