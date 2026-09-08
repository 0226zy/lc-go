package romantointeger

// RomanToInt 罗马数字转整数
// 给定一个罗马数字字符串，将其转换为整数。罗马数字规则：左加右减，
// 小数字在大数字左边表示减法（如 IV=4），其余情况表示加法。
// 时间复杂度: O(n)  空间复杂度: O(1)（字符到值的映射表大小固定）
func RomanToInt(s string) int {
	// 罗马数字字符到整数的映射
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		cur := values[s[i]]
		// 如果当前字符的值小于下一个字符的值，则为减法（如 IV、IX）
		if i+1 < len(s) && cur < values[s[i+1]] {
			result -= cur
		} else {
			result += cur
		}
	}
	return result
}
