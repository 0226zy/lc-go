package bitwiseandofnumbersrange

// RangeBitwiseAnd 数字范围按位与
// 给定两个整数 left 和 right，表示区间 [left, right]，返回此区间内所有数字按位与的结果
// （包含 left、right 端点）。
// 方法：不断清除 right 的最低有效位（right &= right-1），直到 right <= left，
// 此时 right 即为区间内所有数的公共二进制前缀。
// 时间复杂度: O(1)（最多 32 次循环）  空间复杂度: O(1)
func RangeBitwiseAnd(left, right int) int {
	// 区间内只要存在进位/借位差异，低位就会被清零；
	// 反复去掉 right 的最低 1，相当于把区间内会变化的低位抹掉
	for left < right {
		right &= right - 1
	}
	return right
}

// RangeBitwiseAndShift 数字范围按位与（右移对齐版）
// 同时右移 left、right，直到二者相等，得到公共前缀后再左移回去。
// 时间复杂度: O(1)（最多 32 次循环）  空间复杂度: O(1)
func RangeBitwiseAndShift(left, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	return left << shift
}
