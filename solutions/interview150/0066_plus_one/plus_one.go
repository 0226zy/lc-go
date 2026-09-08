// Package plusone 提供 LeetCode 66. 加一 的题解实现。
package plusone

// PlusOne 加一
// 给定一个由非负整数组成的数组 digits 表示一个大整数（最高位在数组首位，无前导零），
// 将该整数加一，并以同样的数组形式返回结果。
// 时间复杂度: O(n) 最坏情况（全部为 9）需要从末位进位到首位  空间复杂度: O(1)（不考虑进位扩容时为 O(n)）
func PlusOne(digits []int) []int {
	// 从末位开始处理进位
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			// 当前位加一后不需要继续进位，直接返回
			digits[i]++
			return digits
		}
		// 当前位是 9，加一后变为 0，向高位进位
		digits[i] = 0
	}
	// 所有位都是 9，结果需要扩容一位（如 999 + 1 = 1000）
	return append([]int{1}, digits...)
}
