package fibonaccinumber

// Fib 斐波那契数
// 时间复杂度: O(?)  空间复杂度: O(?)
func Fib(n int) int {
	if n < 2 {
		return n
	}
	f := make([]int, n+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
