package decodestring

// DecodeString 字符串解码
// 给定一个经过编码的字符串，返回它解码后的字符串。
// 编码规则: k[encoded_string]，表示方括号内的 encoded_string 恰好重复 k 次。
// 时间复杂度: O(?)  空间复杂度: O(?)
func DecodeString(s string) string {
	stack := []byte{}
	for _, c := range s {
		b := byte(c)
		stack = append(stack, b)
	}
	return popStack(stack)
}

func popStack(stack []byte) string {

	stack = stack[:len(stack)-1]
	curr := ""
	str := []byte{}
	for len(stack) > 0 {
		n := len(stack)
		top := stack[n-1]
		stack = stack[:n-1]
		if top == '[' || top == ']' {
			continue
		}

		if isDigit(top) {
			curr = repeate(string(str)+curr, int(top-'0'))
		} else {
			str = append([]byte{top}, str...)
		}

	}
	return curr

}

func isDigit(b byte) bool {
	if int(b-'0') >= 0 && int(b-'0') <= 9 {
		return true
	}
	return false

}

func repeate(str string, n int) string {
	ret := ""
	for i := 0; i < n; i++ {
		ret = ret + str
	}
	return ret
}
