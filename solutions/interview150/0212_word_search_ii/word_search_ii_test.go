package wordsearchii

import (
	"sort"
	"testing"
)

// equalStringSet 比较两个字符串集合是否相等（结果顺序无关）
func equalStringSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sa := append([]string{}, a...)
	sb := append([]string{}, b...)
	sort.Strings(sa)
	sort.Strings(sb)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		// LeetCode 官方示例
		{"示例1", [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{"示例2", [][]byte{{'a', 'b'}, {'c', 'd'}}, []string{"abcb"}, nil},

		// 边界：单词列表为空
		{"空单词列表", [][]byte{{'a'}}, nil, nil},

		// 边界：单格棋盘命中与未命中
		{"单格命中", [][]byte{{'a'}}, []string{"a"}, []string{"a"}},
		{"单格未命中", [][]byte{{'a'}}, []string{"ab"}, nil},

		// 边界：同一格不能重复使用（aba 在 [a,b] 上不可构造）
		{"禁止重复用格", [][]byte{{'a', 'b'}}, []string{"aba"}, nil},

		// 边界：同一单词被找到一次（棋盘含重复出现也只收一次）
		{"重复单词只收一次", [][]byte{
			{'a', 'a'},
			{'a', 'a'},
		}, []string{"aaa"}, []string{"aaa"}},

		// 边界：单词需转弯才能找到
		{"转弯路径", [][]byte{
			{'a', 'b', 'c'},
			{'f', 'e', 'd'},
		}, []string{"abcdef"}, []string{"abcdef"}},

		// 边界：部分命中（b 与 c 对角不相邻，abc/abcd 无法构成）
		{"部分命中", [][]byte{
			{'a', 'b'},
			{'c', 'd'},
		}, []string{"abc", "abcd", "abe", "cdba"}, []string{"cdba"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindWords(tt.board, tt.words)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !equalStringSet(got, tt.want) {
				t.Errorf("FindWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindWords(b *testing.B) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindWords(board, words)
	}
}
