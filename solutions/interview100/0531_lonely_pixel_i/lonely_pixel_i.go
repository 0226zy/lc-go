package lonelypixeli

// FindLonelyPixel 孤独像素 I
// 给定 m×n 的图片 picture（'W' 为白色像素、'B' 为黑色像素），统计黑色孤独像素的数量：
// 黑色孤独像素指其所在行与所在列都没有其他黑色像素的 'B'。
// 时间复杂度: O(m*n) 两遍扫描  空间复杂度: O(m+n) 行列计数数组
func FindLonelyPixel(picture [][]byte) int {
	if len(picture) == 0 {
		return 0
	}
	m, n := len(picture), len(picture[0])
	// rowCnt[i] 表示第 i 行黑色像素个数，colCnt[j] 表示第 j 列黑色像素个数
	rowCnt := make([]int, m)
	colCnt := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rowCnt[i]++
				colCnt[j]++
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 行、列都恰好只有这一个黑色像素时才是孤独像素
			if picture[i][j] == 'B' && rowCnt[i] == 1 && colCnt[j] == 1 {
				ans++
			}
		}
	}
	return ans
}
