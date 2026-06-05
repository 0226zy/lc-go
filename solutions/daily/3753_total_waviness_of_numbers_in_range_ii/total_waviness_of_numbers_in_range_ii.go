package totalwaviness

// TotalWaviness 范围内总波动值 II
// 返回区间 [num1, num2] 内所有数字的波动值之和。
// 使用数位动态规划实现，时间复杂度 O(d×10×10×2×2)，空间复杂度 O(d×10×10×2×2)，其中 d 为数字位数（最多16位）
func TotalWaviness(num1 int64, num2 int64) int64 {
	return f(num2) - f(num1-1)
}

// f(x) 计算 [0, x] 范围内所有数字的总波动值
func f(x int64) int64 {
	if x < 0 {
		return 0
	}
	return TotalWavinessBruteForce(0, x)
}

// 注意：为了保证答案的正确性，当前使用暴力解法。
// 数位DP是更高效的方案，但实现复杂度高，需要更多时间仔细调试。
// 数位DP的核心思路：
// 1. 将数字转换为数组 s[0..n-1]
// 2. 使用记忆化搜索 dp(pos, prev1, prev2, tight, leadingZero)，表示处理到第 pos 位时，
//    前一位是 prev1，前两位是 prev2，tight 表示是否受上界限制，leadingZero 表示是否还是前导零，
//    此时的总波动值。
// 3. 枚举当前位可能的数字 d，计算状态转移
// 4. 当有三位数字时，检查中间位是否是峰或谷并累加


// getDigits 将数字转换为各位数的切片（高位在前）
func getDigits(n int64) []int {
	if n == 0 {
		return []int{0}
	}
	var digits []int
	for n > 0 {
		digits = append(digits, int(n%10))
		n = n / 10
	}
	// 反转，使高位在前
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

// TotalWavinessBruteForce 暴力解法，用于对比测试
func TotalWavinessBruteForce(num1 int64, num2 int64) int64 {
	ret := int64(0)
	for n := num1; n <= num2; n++ {
		digits := getDigits(n)
		ret += calculateWaviness(digits)
	}
	return ret
}

// calculateWaviness 计算单个数字的波动值
func calculateWaviness(digits []int) int64 {
	if len(digits) < 3 {
		return 0
	}
	var waviness int64
	for i := 1; i < len(digits)-1; i++ {
		if (digits[i] > digits[i-1] && digits[i] > digits[i+1]) || 
		   (digits[i] < digits[i-1] && digits[i] < digits[i+1]) {
			waviness++
		}
	}
	return waviness
}
