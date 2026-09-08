package basiccalculator

// Calculate 基本计算器
// 计算只包含 '+'、'-'、括号与非负整数的有效表达式的值，允许空格。
// 核心技巧：减法看作加负数，用符号位处理；遇到括号把 (结果, 符号) 压栈保存现场。
// 时间复杂度: O(n) 每个字符只处理一次  空间复杂度: O(n) 栈深度为括号嵌套层数
func Calculate(s string) int {
	res := 0  // 当前层已累计的结果
	sign := 1 // 当前数字的符号：+1 或 -1
	num := 0  // 正在读取的多位数字
	var stack []int

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ' ':
			// 空格无意义
		case c >= '0' && c <= '9':
			num = num*10 + int(c) - int('0')
		case c == '+':
			res += sign * num // 结算前一个数字
			num = 0
			sign = 1
		case c == '-':
			res += sign * num
			num = 0
			sign = -1
		case c == '(':
			// 保存现场，进入括号内的新一层
			stack = append(stack, res, sign)
			res, sign = 0, 1
		case c == ')':
			res += sign * num // 括号内最后一个数字
			num = 0
			prevSign := stack[len(stack)-1]
			prevRes := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// 括号整体替换为一个数：上层结果 + 上层符号 * 括号内结果
			res = prevRes + prevSign*res
		}
	}
	res += sign * num // 结算最后一个数字
	return res
}
