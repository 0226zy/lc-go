package integertoroman

// IntToRoman 整数转罗马数字
// 给定一个整数 num，将其转换为罗马数字字符串。num 取值范围为 [1, 3999]。
// 时间复杂度: O(1)  空间复杂度: O(1)（符号表大小固定）
func IntToRoman(num int) string {
	// 预定义所有可能用到的符号（含减法表示的 6 个组合），从大到小排列
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(vals) && num > 0; i++ {
		// 用当前最大的符号尽可能多地表示剩余数值
		for num >= vals[i] {
			result += symbols[i]
			num -= vals[i]
		}
	}
	return result
}
