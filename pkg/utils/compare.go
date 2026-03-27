package utils

import "sort"

// EqualIntSlice 比较两个 int 切片是否相等（顺序敏感）
func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// EqualIntSliceUnordered 比较两个 int 切片是否包含相同元素（顺序无关）
func EqualIntSliceUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	return EqualIntSlice(aCopy, bCopy)
}

// EqualStringSlice 比较两个 string 切片是否相等
func EqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Equal2DIntSlice 比较两个二维 int 切片是否相等
func Equal2DIntSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !EqualIntSlice(a[i], b[i]) {
			return false
		}
	}
	return true
}
