package validparentheses

// IsValid 有效的括号
// 给定一个只包含 '(', ')', '[', ']', '{', '}' 的字符串 s，判断括号是否有效匹配。
// 时间复杂度: O(n) 每个字符最多入栈、出栈一次  空间复杂度: O(n) 最坏全部为左括号
func IsValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			// 左括号入栈，等待匹配的右括号
			stack = append(stack, c)
		case ')', ']', '}':
			// 右括号：必须与最近的未闭合左括号（栈顶）匹配
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (c == ')' && top != '(') ||
				(c == ']' && top != '[') ||
				(c == '}' && top != '{') {
				return false
			}
		}
	}
	// 栈为空说明所有左括号都被闭合
	return len(stack) == 0
}
