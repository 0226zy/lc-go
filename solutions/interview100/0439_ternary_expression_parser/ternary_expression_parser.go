package ternaryexpressionparser

// ParseTernary 三元表达式解析器
// 给定一个表示任意嵌套三元表达式的字符串 expression，返回其求值结果的字符串表示。
// 表达式规则：'T'/'F' 为基本值，expr1?expr2:expr3 表示条件选择，结果值均为单位字符。
// 时间复杂度: O(n) 每个字符最多进出栈一次  空间复杂度: O(n) 栈最多容纳 n 个字符
func ParseTernary(expression string) string {
	// 从右向左扫描，把已归约的字符压栈；遇到 '?' 时完成一次局部求值
	stack := make([]byte, 0, len(expression))
	i := len(expression) - 1
	for i >= 0 {
		if expression[i] != '?' {
			stack = append(stack, expression[i])
			i--
			continue
		}
		// 栈顶依次为：真分支值、':'、假分支值
		first := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 弹出真分支值
		stack = stack[:len(stack)-1] // 弹出 ':'
		second := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 弹出假分支值
		// '?' 左侧紧邻的字符即为条件
		if expression[i-1] == 'T' {
			stack = append(stack, first)
		} else {
			stack = append(stack, second)
		}
		i -= 2 // 跳过 '?' 和条件字符
	}
	return string(stack[0])
}
