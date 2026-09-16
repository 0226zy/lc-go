package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: ADOBECODEBANC / ABC", "ADOBECODEBANC", "ABC", "BANC"},
		{"示例2: a / aa", "a", "aa", ""},

		// 边界：s 比 t 短
		{"s比t短", "ab", "abc", ""},
		// 边界：s 与 t 完全相等
		{"s与t相等", "abc", "abc", "abc"},
		// 边界：单个字符命中
		{"单字符命中", "a", "a", "a"},
		// 边界：t 中字符重复，需要按次数覆盖
		{"t含重复字符", "aa", "aa", "aa"},
		{"t含重复字符2", "ab", "aab", ""},
		{"重复字符只需部分覆盖", "ADOBECODEBANC", "AABC", "ADOBECODEBA"},
		// 边界：t 的字符全在 s 末尾
		{"字符集中在末尾", "xyzabc", "abc", "abc"},
		// 边界：最小窗口在开头
		{"窗口在开头", "abcdef", "abc", "abc"},
		// 边界：大小写敏感
		{"大小写敏感", "aA", "aA", "aA"},
		{"大小写不同不匹配", "ab", "A", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinWindow(tt.s, tt.t); got != tt.want {
				t.Errorf("MinWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkMinWindow(b *testing.B) {
	// 构造长串：随机大小写字母，并在中间插入一次 t
	long := make([]byte, 100000)
	for i := range long {
		long[i] = byte('a' + i%26)
	}
	copy(long[50000:], "ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	benchmarks := []struct {
		name string
		s    string
		t    string
	}{
		{"len=13", "ADOBECODEBANC", "ABC"},
		{"len=1000", string(long[:1000]), "ABCNOP"},
		{"len=10000", string(long[:10000]), "ABCDEFGHIJ"},
		{"len=100000", string(long), "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinWindow(bm.s, bm.t)
			}
		})
	}
}
