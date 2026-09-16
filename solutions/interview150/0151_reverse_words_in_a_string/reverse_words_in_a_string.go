package reversewordsinastring

import "strings"

// ReverseWords 反转字符串中的单词
// 给定字符串 s，按顺序反转其中所有单词（单词顺序颠倒，单词内字符顺序不变），
// 单词间只保留一个空格，无首尾空格。
// 时间复杂度: O(n)  空间复杂度: O(n)（切片重建）
func ReverseWords(s string) string {
	return ReverseWords2(s)
}

func reverse(s string) string {

	// 按空白字符分割，自动忽略多余空格
	words := strings.Fields(s)

	// 双指针原地反转单词切片
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func ReverseWords2(s string) string {
	b := []byte(s)
	n := len(b)
	res := make([]byte, 0, n)
	right := n - 1
	for right >= 0 {
		for right >= 0 && b[right] == ' ' {
			right--
		}
		if right < 0 {
			break
		}
		left := right
		for left >= 0 && b[left] != ' ' {
			left--
		}
		word := b[left+1 : right+1]
		if len(res) > 0 {
			res = append(res, ' ')
		}
		res = append(res, word...)
		right = left
	}
	return string(res)
}
