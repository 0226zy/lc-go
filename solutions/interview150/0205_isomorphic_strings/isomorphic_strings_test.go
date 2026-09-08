package isomorphicstrings

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: egg 与 add 同构", "egg", "add", true},
		{"示例2: foo 与 bar 不同构", "foo", "bar", false},
		{"示例3: paper 与 title 同构", "paper", "title", true},

		// 边界：空串
		{"两个空串", "", "", true},

		// 边界：单字符
		{"单字符相同", "a", "a", true},
		{"单字符不同", "a", "b", true},

		// 边界：长度不同
		{"长度不同", "ab", "a", false},

		// 边界：一对多冲突（同一个 s 字符映射到不同 t 字符）
		{"s 中字符映射不一致", "ab", "aa", false},

		// 边界：多对一冲突（不同 s 字符映射到同一个 t 字符）
		{"t 中字符被重复映射", "aa", "ab", false},

		// 边界：相同字符串
		{"完全相同的字符串", "abc", "abc", true},

		// 边界：重复模式
		{"重复模式", "abab", "cdcd", true},
		{"重复模式冲突", "abab", "cdce", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("IsIsomorphic(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkIsIsomorphic(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		t    string
	}{
		{"短串", "egg", "add"},
		{"长串同构", generateRepeatStr("ab", 500), generateRepeatStr("xy", 500)},
		{"长串非同构", generateRepeatStr("aab", 400), generateRepeatStr("xyz", 400)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsIsomorphic(bm.s, bm.t)
			}
		})
	}
}

// generateRepeatStr 将 pattern 重复 n 次拼接成字符串
func generateRepeatStr(pattern string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += pattern
	}
	return result
}
