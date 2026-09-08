package validpalindrome

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: A man, a plan, a canal: Panama 是回文", "A man, a plan, a canal: Panama", true},
		{"示例2: race a car 不是回文", "race a car", false},
		{"示例3: 单个空格视为回文", " ", true},

		// 边界：单字符与纯无效字符
		{"单个字母是回文", "a", true},
		{"单个数字是回文", "0", true},
		{"单个标点是回文", ".", true},
		{"纯标点串是回文", ",.;:!", true},

		// 边界：大小写不敏感
		{"大小写混合回文 abBa", "abBa", true},
		{"大小写混合回文 Noon", "Noon", true},
		{"大小写不同但非回文 aBc", "aBc", false},

		// 边界：数字也是有效字符
		{"含数字回文 1abba1", "1abba1", true},
		{"数字与字母不匹配 0P 不是回文", "0P", false},
		{"纯数字回文 12321", "12321", true},
		{"纯数字非回文 12345", "12345", false},

		// 边界：无效字符分布
		{"首尾都是无效字符", "!!aba!!", true},
		{"无效字符夹杂回文 aba", "a,b-a", true},
		{"无效字符夹杂非回文 abc", "a,b-c", false},
		{"偶数个有效字符回文", "ab,,ba", true},
		{"奇数个有效字符回文", "a b a", true},

		// 常规：常见回文短语
		{"Madam 是回文", "Madam", true},
		{"hello 不是回文", "hello", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// TestIsPalindromeStress 压力场景：构造超长字符串验证正确性与稳定性
func TestIsPalindromeStress(t *testing.T) {
	// 20 万个字符的回文：大量无效字符 + 镜像字母数字
	half := strings.Repeat("a1, ", 25000) // 100000 个字符
	longPalindrome := half + strings.Repeat(" ,1a", 25000)
	t.Run("20万字符回文", func(t *testing.T) {
		if !IsPalindrome(longPalindrome) {
			t.Errorf("IsPalindrome(20万字符回文) = false, want true")
		}
	})

	// 20 万个字符的非回文：首尾有效字符不同，应尽早返回
	longNotPalindrome := "a" + strings.Repeat("b", 199998) + "c"
	t.Run("20万字符非回文", func(t *testing.T) {
		if IsPalindrome(longNotPalindrome) {
			t.Errorf("IsPalindrome(20万字符非回文) = true, want false")
		}
	})

	// 20 万个纯无效字符：视为空串，是回文
	t.Run("20万个纯空格是回文", func(t *testing.T) {
		if !IsPalindrome(strings.Repeat(" ", 200000)) {
			t.Errorf("IsPalindrome(20万个空格) = false, want true")
		}
	})
}

func BenchmarkIsPalindrome(b *testing.B) {
	half := strings.Repeat("a1, ", 25000)
	longPalindrome := half + strings.Repeat(" ,1a", 25000)

	benchmarks := []struct {
		name string
		s    string
	}{
		{"短回文 A man, a plan, a canal: Panama", "A man, a plan, a canal: Panama"},
		{"短非回文 race a car", "race a car"},
		{"20万字符回文", longPalindrome},
		{"20万字符非回文首尾即失败", "a" + strings.Repeat("b", 199998) + "c"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPalindrome(bm.s)
			}
		})
	}
}
