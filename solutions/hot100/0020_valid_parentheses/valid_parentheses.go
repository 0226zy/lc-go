package validparentheses

// IsValid 有效的括号（栈）
// 使用栈匹配括号，遇到左括号入栈，遇到右括号检查栈顶是否匹配。
// 时间复杂度: O(n)  空间复杂度: O(n)
func IsValid(s string) bool {
	stack := make([]rune, 0, len(s))
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// IsValidFast 有效的括号（栈 + 提前返回）
// 在字符串长度为奇数时提前返回 false，减少不必要的计算。
// 时间复杂度: O(n)  空间复杂度: O(n)
func IsValidFast(s string) bool {
	// 奇数长度不可能有效
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]rune, 0, len(s))
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// IsValidBytes 有效的括号（使用 byte 切片作为栈，性能优化版）
// 使用 byte 而非 rune，减少内存分配。
// 时间复杂度: O(n)  空间复杂度: O(n)
func IsValidBytes(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	stack := make([]byte, 0, len(s))
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
