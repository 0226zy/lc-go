package decodestring

import (
	"strings"
)

// DecodeString 字符串解码
// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则: k[encoded_string]，表示方括号内的 encoded_string 恰好重复 k 次。
// 时间复杂度: O(?)  空间复杂度: O(?)
func DecodeString(s string) string {
	if len(s) < 1 {
		return s
	}
	stack := []string{}
	ptr := 0
	for ptr < len(s) {
		if isDigit(s[ptr]) {
			tmp := ""
			tmp, ptr = getNum(s, ptr)
			stack = append(stack, tmp)
			continue
		}
		if s[ptr] == '[' || isAplha(s[ptr]) {
			stack = append(stack, string(s[ptr]))
			ptr++
			continue
		}

		// s[ptr]==']'
		ptr++
		str := ""
		for stack[len(stack)-1] != "[" {
			str = stack[len(stack)-1] + str
			stack = stack[:len(stack)-1]
		}

		// pop '['
		stack = stack[:len(stack)-1]
		// pop 'num'
		strNum := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		num := toNum(strNum)
		// repeated string
		str = strings.Repeat(str, num)
		stack = append(stack, str)

	}
	ret := ""
	for _, str := range stack {
		ret = ret + str
	}
	return ret
}

func toNum(str string) int {
	ret := 0
	for _, ch := range str {
		ret = ret*10 + int(ch-'0')
	}
	return ret
}

func isAplha(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}

func getNum(s string, ptr int) (string, int) {
	ret := ""
	for ; isDigit(s[ptr]); ptr++ {
		ret = ret + string(s[ptr])
	}
	return ret, ptr
}

func isDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}
