package evaluatereversepolishnotation

import (
	"fmt"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 加减乘混合", []string{"2", "1", "+", "3", "*"}, 9},
		{"示例2: 除法向零截断", []string{"4", "13", "5", "/", "+"}, 6},
		{"示例3: 长表达式与负数", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},

		// 边界：单元素
		{"单个数字", []string{"42"}, 42},

		// 边界：减法与除法的操作数顺序
		{"减法顺序_5减2", []string{"5", "2", "-"}, 3},
		{"减法顺序_2减5为负", []string{"2", "5", "-"}, -3},
		{"除法顺序_15除3", []string{"15", "3", "/"}, 5},
		{"除法结果为负且截断", []string{"7", "-2", "/"}, -3}, // 7 / -2 = -3.5，向零截断为 -3

		// 边界：负数参与运算
		{"负数乘法", []string{"-3", "4", "*"}, -12},
		{"两个负数相乘", []string{"-3", "-4", "*"}, 12},
		{"负数加法", []string{"-5", "-3", "+"}, -8},

		// 边界：嵌套表达式
		{"多层嵌套", []string{"3", "4", "+", "5", "*"}, 35},
		{"除法嵌套", []string{"18", "6", "/", "3", "/", "2", "+"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvalRPN(tt.tokens); got != tt.want {
				t.Errorf("EvalRPN(%v) = %d, want %d", tt.tokens, got, tt.want)
			}
		})
	}
}

func BenchmarkEvalRPN(b *testing.B) {
	b.Run("官方示例3长表达式", func(b *testing.B) {
		tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
		for i := 0; i < b.N; i++ {
			EvalRPN(tokens)
		}
	})
	b.Run("长表达式_长度10001", func(b *testing.B) {
		tokens := generateRPN(5000)
		for i := 0; i < b.N; i++ {
			EvalRPN(tokens)
		}
	})
}

// generateRPN 生成长度为 2n+1 的合法逆波兰表达式
// 结构为 1 2 + 3 + 4 + ...，结果为 1+2+...+n
func generateRPN(n int) []string {
	tokens := make([]string, 0, 2*n+1)
	tokens = append(tokens, "1")
	for i := 2; i <= n; i++ {
		tokens = append(tokens, fmt.Sprintf("%d", i), "+")
	}
	return tokens
}
