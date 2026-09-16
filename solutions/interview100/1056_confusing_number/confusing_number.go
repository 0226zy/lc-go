package confusingnumber

// ConfusingNumber 易混淆数
// 给定一个整数 n，判断它是否为易混淆数：将 n 整体旋转 180° 后仍是一个有效数字，
// 且旋转后的数字与 n 不同。数字 0/1/8 旋转后不变，6/9 互相转换，其余数字旋转后无效。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func ConfusingNumber(n int) bool {
	// rotate[d] 表示数字 d 旋转 180° 后的数字，-1 表示旋转后无效
	rotate := [10]int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}
	origin := n
	rotated := 0
	// 从低位到高位取原数的每一位，累加后天然得到旋转后的数字
	for n > 0 {
		d := n % 10
		if rotate[d] == -1 {
			return false
		}
		rotated = rotated*10 + rotate[d]
		n /= 10
	}
	return rotated != origin
}
