package countsubstringswithonlyonedistinctletter

import "testing"

func TestCountLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: aaaba", "aaaba", 8},
		{"示例2: aaaaaaaaaa", "aaaaaaaaaa", 55},
		{"示例3: abc", "abc", 3},

		// 边界：空串与单字符
		{"空字符串", "", 0},
		{"单字符", "z", 1},

		// 全部字符相同
		{"两个相同字符", "aa", 3},
		{"全部相同字符", "zzzzz", 15},

		// 多段混合
		{"交替字符", "abab", 4},
		// 段长 3、3、2 分别贡献 6、6、3
		{"长段加短段", "aaabbbaa", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLetters(tt.s); got != tt.want {
				t.Errorf("CountLetters(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkCountLetters(b *testing.B) {
	// 构造长度 1000 的混合字符串：a*10 b*10 ... 循环
	buf := make([]byte, 0, 1000)
	for len(buf) < 1000 {
		for c := byte('a'); c <= byte('j'); c++ {
			for k := 0; k < 10; k++ {
				buf = append(buf, c)
			}
		}
	}
	s := string(buf[:1000])

	b.Run("长度1000混合段", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			CountLetters(s)
		}
	})
}
