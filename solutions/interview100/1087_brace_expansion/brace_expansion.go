package braceexpansion

import (
	"sort"
	"strings"
)

// Expand 花括号展开
// 给定字符串 s，其中 {a,b,c} 表示该位置可取括号中任意一个字符，
// 未被花括号包裹的字符为固定字符，返回所有能组成的单词，按字典序升序排列。
// 时间复杂度: O(L + P*M)，L 为 s 长度，P 为结果总数，M 为组数  空间复杂度: O(M)
func Expand(s string) []string {
	// 第一步：把 s 解析成若干候选组
	var groups [][]string
	for i := 0; i < len(s); {
		if s[i] == '{' {
			// 找到配对的花括号，按逗号切分候选字符
			j := i
			for s[j] != '}' {
				j++
			}
			groups = append(groups, strings.Split(s[i+1:j], ","))
			i = j + 1
		} else {
			// 普通字符单独成组
			groups = append(groups, []string{string(s[i])})
			i++
		}
	}

	// 第二步：每组候选排序，保证回溯枚举出的结果天然是字典序
	for _, g := range groups {
		sort.Strings(g)
	}

	// 第三步：回溯枚举所有组合
	var res []string
	path := make([]byte, 0, len(groups))
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(groups) {
			res = append(res, string(path))
			return
		}
		for _, ch := range groups[idx] {
			path = append(path, ch[0])
			dfs(idx + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
