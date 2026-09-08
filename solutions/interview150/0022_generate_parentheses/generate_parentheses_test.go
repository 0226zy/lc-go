package generateparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: n=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"示例2: n=1", 1, []string{"()"}},

		// 边界：n=2
		{"n=2", 2, []string{"(())", "()()"}},
		// 边界：最大值 8，共 C8=1430 种，验证数量
		{"n=8数量", 8, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
				}
				return
			}
			if len(got) != 1430 {
				t.Errorf("GenerateParenthesis(8) 长度 = %d, want 1430（卡特兰数 C8）", len(got))
			}
		})
	}
}

// TestGenerateParenthesisValid 验证生成的每个组合都是合法的
func TestGenerateParenthesisValid(t *testing.T) {
	for n := 1; n <= 6; n++ {
		got := GenerateParenthesis(n)
		wantCount := catalan(n)
		if len(got) != wantCount {
			t.Fatalf("GenerateParenthesis(%d) 长度 = %d, want %d", n, len(got), wantCount)
		}
		for _, s := range got {
			if !isValid(s, n) {
				t.Errorf("GenerateParenthesis(%d) 生成了非法组合 %q", n, s)
			}
		}
	}
}

// catalan 计算卡特兰数 C_n
func catalan(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

// isValid 检查括号串是否合法：任意前缀中 '(' 不少于 ')'，且总数都为 n
func isValid(s string, n int) bool {
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
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=6", 6},
		{"n=8", 8},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateParenthesis(bm.n)
			}
		})
	}
}
