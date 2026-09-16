package fourkeyskeyboard

// MaxA 四个键的键盘
// 键盘只有 A / Ctrl-A / Ctrl-C / Ctrl-V 四个键，初始屏幕为空，
// 求最多按 n 次键后屏幕上最多能出现多少个 A。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func MaxA(n int) int {
	// f[i] 表示按 i 次键最多能得到的 A 的个数
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 情况一：最后一次按 A
		f[i] = f[i-1] + 1
		// 情况二：第 j 次按键后进行 全选+复制+(i-j-2) 次粘贴，
		// 最终 A 的个数为 f[j] * (i-j-1)
		for j := 1; j <= i-3; j++ {
			if v := f[j] * (i - j - 1); v > f[i] {
				f[i] = v
			}
		}
	}
	return f[n]
}
