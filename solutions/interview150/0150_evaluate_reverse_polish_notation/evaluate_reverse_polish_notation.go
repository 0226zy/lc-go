package evaluatereversepolishnotation

import "strconv"

// EvalRPN 逆波兰表达式求值
// 根据逆波兰表示法（后缀表达式）计算算术表达式的值，除法向零截断。
// 时间复杂度: O(n) 每个 token 只处理一次  空间复杂度: O(n) 栈最多存放所有操作数
func EvalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, tok := range tokens {
		if len(tok) == 1 && (tok[0] == '+' || tok[0] == '-' ||
			tok[0] == '*' || tok[0] == '/') {
			// 运算符：先弹出的是右操作数，后弹出的是左操作数
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var res int
			switch tok[0] {
			case '+':
				res = a + b
			case '-':
				res = a - b
			case '*':
				res = a * b
			case '/':
				res = a / b // Go 整数除法天然向零截断
			}
			stack = append(stack, res)
		} else {
			num, _ := strconv.Atoi(tok)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}
