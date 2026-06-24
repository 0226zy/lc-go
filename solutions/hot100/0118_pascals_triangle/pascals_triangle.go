package pascalstriangle

// Generate 杨辉三角
// 给定非负整数 numRows，生成杨辉三角的前 numRows 行。
// 每个数是它左上方和右上方数的和。
// 时间复杂度: O(n²)  空间复杂度: O(n²)
func Generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1 // 首尾元素恒为 1
		for j := 1; j < i; j++ {
			row[j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		triangle[i] = row
	}
	return triangle
}

// GenerateInPlace 杨辉三角（原地追加版）
// 通过 append 构建每一行，逻辑等价但写法更紧凑。
// 时间复杂度: O(n²)  空间复杂度: O(n²)
func GenerateInPlace(numRows int) [][]int {
	triangle := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				row[j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
		triangle = append(triangle, row)
	}
	return triangle
}
