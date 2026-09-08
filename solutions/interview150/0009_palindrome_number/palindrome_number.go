package palindromenumber

// IsPalindrome 回文数
// 判断一个整数 x 是否是回文数：正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 时间复杂度: O(log n) 数字位数的一半  空间复杂度: O(1) 只使用常数变量
func IsPalindrome(x int) bool {
	// 负数一定不是回文数；个位是 0 且不是 0 本身的数倒过来会丢前导零，也不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	// 只反转数字的后一半，避免整个反转可能发生的 int 溢出
	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 偶数位：前后两半相等；奇数位：中间位单独剩下，去掉它再比较
	return x == reversed || x == reversed/10
}
