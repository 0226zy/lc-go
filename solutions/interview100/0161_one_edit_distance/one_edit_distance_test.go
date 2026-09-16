package oneeditdistance

import "testing"

func TestIsOneEditDistance(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: ab->acb 插入", "ab", "acb", true},
		{"示例2: cab->ad 需两次编辑", "cab", "ad", false},
		{"示例3: 1203->1213 替换", "1203", "1213", true},

		// 边界：空串
		{"空串与单字符", "", "a", true},
		{"双空串距离为0", "", "", false},

		// 边界：完全相同
		{"完全相等", "abc", "abc", false},
		{"单字符相等", "a", "a", false},

		// 替换场景
		{"首字符替换", "abc", "zbc", true},
		{"尾字符替换", "abc", "abz", true},
		{"多处不同", "abc", "axy", false},

		// 插入/删除场景
		{"末尾插入", "abc", "abcd", true},
		{"末尾删除", "abcd", "abc", true},
		{"开头插入", "bc", "abc", true},
		{"中间插入", "acd", "abcd", true},
		{"前缀后多处不同", "abcd", "abxyd", false},
		{"长度差1但内容差两处", "ab", "ba", false},

		// 边界：长度差大于 1
		{"长度差为2", "abc", "abcde", false},
		{"长度差悬殊", "", "abc", false},

		// 长短串颠倒（验证交换逻辑）
		{"长串在前仍正确", "acb", "ab", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneEditDistance(tt.s, tt.t); got != tt.want {
				t.Errorf("IsOneEditDistance(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkIsOneEditDistance(b *testing.B) {
	// 构造两个仅末尾差一个字符的长字符串
	long := make([]byte, 10000)
	for i := range long {
		long[i] = 'a'
	}
	s := string(long)
	t1 := s + "b"

	b.Run("前缀全同_末尾插入", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsOneEditDistance(s, t1)
		}
	})
	b.Run("短串立即命中", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsOneEditDistance("abc", "abcd")
		}
	})
}
