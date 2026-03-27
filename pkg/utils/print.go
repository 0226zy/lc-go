package utils

import "fmt"

// PrintIntSlice 打印 int 切片
func PrintIntSlice(nums []int) {
	fmt.Println(IntSliceToString(nums))
}

// PrintMatrix 打印二维矩阵
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(IntSliceToString(row))
	}
}

// PrintStringSlice 打印 string 切片
func PrintStringSlice(strs []string) {
	fmt.Printf("[%s]\n", joinStrings(strs))
}

func joinStrings(strs []string) string {
	result := ""
	for i, s := range strs {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%q", s)
	}
	return result
}
