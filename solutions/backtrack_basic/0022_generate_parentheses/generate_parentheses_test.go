package generateparentheses

import (
	"sort"
	"testing"
)

// normalize 对字符串切片排序，用于无序比较两组括号组合
func normalize(s []string) []string {
	cp := make([]string, len(s))
	copy(cp, s)
	sort.Strings(cp)
	return cp
}

func equalStringSlice(a, b []string) bool {
	a, b = normalize(a), normalize(b)
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

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: n=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"示例2: n=1", 1, []string{"()"}},

		// 边界情况
		{"边界: n=2 两种组合", 2, []string{"(())", "()()"}},
		{"边界: n=4 共14种", 4, []string{
			"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
			"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
			"()(()())", "()(())()", "()()(())", "()()()()",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if !equalStringSlice(got, tt.want) {
				t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// TestGenerateParenthesisCountAndValid 验证数量等于卡特兰数且每个组合都合法
func TestGenerateParenthesisCountAndValid(t *testing.T) {
	for n := 1; n <= 8; n++ {
		got := GenerateParenthesis(n)
		if len(got) != catalan(n) {
			t.Errorf("GenerateParenthesis(%d) 数量 = %d, 期望卡特兰数 %d", n, len(got), catalan(n))
		}
		seen := make(map[string]bool, len(got))
		for _, s := range got {
			if !isValidParentheses(s, n) {
				t.Errorf("GenerateParenthesis(%d) 生成了非法组合 %q", n, s)
			}
			if seen[s] {
				t.Errorf("GenerateParenthesis(%d) 生成了重复组合 %q", n, s)
			}
			seen[s] = true
		}
	}
}

// catalan 计算第 n 个卡特兰数
func catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

// isValidParentheses 检查括号串是否合法：任意前缀中 '(' 不少于 ')'，且长度恰为 2n
func isValidParentheses(s string, n int) bool {
	if len(s) != 2*n {
		return false
	}
	balance := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
		} else {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParenthesis(8)
	}
}
