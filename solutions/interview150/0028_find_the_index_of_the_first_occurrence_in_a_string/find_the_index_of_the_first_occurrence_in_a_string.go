package findfirstoccurrence

import "strings"

// StrStr 找出字符串中第一个匹配项的下标
// 在 haystack 中查找 needle 第一次出现的下标（从 0 开始），不存在返回 -1。
// 采用 KMP 算法，失配时模式串指针按 next 数组回退，主串指针不回退。
// 时间复杂度: O(m + n)  m、haystack 长度, n、needle 长度  空间复杂度: O(n) next 数组
func StrStr(haystack, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}

	// 预处理 next 数组：next[i] 表示 needle[0..i] 的最长相等真前后缀长度
	next := make([]int, n)
	for i, j := 1, 0; i < n; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1] // 失配则回退到更短的前缀
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}

	// 主匹配：i 遍历 haystack，j 遍历 needle
	for i, j := 0, 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// StrStrBuiltin 基于 strings.Index 的实现
// Go 标准库 strings.Index 内部使用 Rabin-Karp 算法，作为 KMP 的对照实现。
// 时间复杂度: 平均 O(m + n)  空间复杂度: O(1)
func StrStrBuiltin(haystack, needle string) int {
	return strings.Index(haystack, needle)
}

// StrStrBrute 暴力匹配实现（对照解法）
// 枚举每个起始位置，逐字符比较，失配整体后移一位。
// 时间复杂度: O(m * n)  空间复杂度: O(1)
func StrStrBrute(haystack, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	for i := 0; i+n <= len(haystack); i++ {
		j := 0
		for j < n && haystack[i+j] == needle[j] {
			j++
		}
		if j == n {
			return i
		}
	}
	return -1
}
