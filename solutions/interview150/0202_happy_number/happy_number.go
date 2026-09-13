package happynumber

// IsHappy 快乐数
// 判断正整数 n 是否为快乐数：反复将各位数字平方和替换自身，最终能得到 1 则为快乐数。
// 使用快慢指针检测环，避免哈希表额外空间。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func IsHappy(n int) bool {
	slow, fast := n, getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

// getNext 计算 n 各位数字的平方和
func getNext(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}
