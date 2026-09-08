package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: anagram 与 nagaram", "anagram", "nagaram", true},
		{"示例2: rat 与 car", "rat", "car", false},

		// 边界：空串
		{"两个空串", "", "", true},
		{"一空一非空", "", "a", false},

		// 边界：相同字符串
		{"完全相同", "abc", "abc", true},

		// 边界：长度不同
		{"长度不同", "ab", "abc", false},

		// 边界：字符种类不同
		{"字符种类不同", "ab", "cd", false},

		// 边界：字符数量不同
		{"字符数量不同", "aab", "abb", false},

		// 边界：单字符
		{"单字符相同", "a", "a", true},
		{"单字符不同", "a", "b", false},

		// 边界：大量重复字符
		{"大量重复字符是异位词", "aaaaabbbbb", "bbbbbaaaaa", true},
		{"大量重复字符非异位词", "aaaaabbbbb", "aaaaabbbbc", false},

		// 边界：部分字符相同
		{"部分字符相同", "hello", "world", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		t    string
	}{
		{"短串", "anagram", "nagaram"},
		{"长串是异位词", generateStr('a', 500) + generateStr('b', 500), generateStr('b', 500) + generateStr('a', 500)},
		{"长串非异位词", generateStr('a', 1000), generateStr('b', 1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsAnagram(bm.s, bm.t)
			}
		})
	}
}

// generateStr 生成 n 个重复字符 c 组成的字符串
func generateStr(c byte, n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = c
	}
	return string(s)
}
