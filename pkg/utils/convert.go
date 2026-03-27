package utils

import (
	"strconv"
	"strings"
)

// IntSliceToString 将 int 切片转为字符串，如 [1, 2, 3]
func IntSliceToString(nums []int) string {
	strs := make([]string, len(nums))
	for i, v := range nums {
		strs[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(strs, ", ") + "]"
}

// StringToIntSlice 将 "[1,2,3]" 格式字符串解析为 int 切片
func StringToIntSlice(s string) []int {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "[]")
	if s == "" {
		return []int{}
	}
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if n, err := strconv.Atoi(p); err == nil {
			result = append(result, n)
		}
	}
	return result
}

// IntToString 将 int 转为 string
func IntToString(n int) string {
	return strconv.Itoa(n)
}

// StringToInt 将 string 转为 int，失败返回 0
func StringToInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}
