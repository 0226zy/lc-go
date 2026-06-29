package nthtribonumber

// Tribonacci 第N个泰波那契数
// 时间复杂度: O(?)  空间复杂度: O(?)
func Tribonacci(n int) int {
	if n < 2 {
		return n
	}
	if n < 3 {
		return 1
	}
	f := make([]int, n+1)
	f[0], f[1], f[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2] + f[i-3]
	}
	return f[n]
}
