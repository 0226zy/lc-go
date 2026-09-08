package basiccalculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 简单加法", "1 + 1", 2},
		{"示例2: 含首尾空格与减法", " 2-1 + 2 ", 3},
		{"示例3: 多层嵌套括号", "(1+(4+5+2)-3)+(6+8)", 23},

		// 边界：单数字与纯空格
		{"单个数字", "42", 42},
		{"单个0", "0", 0},
		{"数字间空格被忽略", "1 2", 12}, // 空格直接跳过，数字连续累加
		{"前后大量空格", "   7   ", 7},

		// 边界：连续加减与符号位
		{"连续减法", "3-2-1", 0},
		{"减加混合", "5-3+2", 4},
		{"结果为负", "2-5", -3},
		{"末尾是负号后的数字", "1-9", -8},

		// 边界：括号
		{"整段括号", "(1)", 1},
		{"括号内负数", "1-(5)", -4},
		{"空括号效果_括号内为0", "1+((2))", 3},
		{"括号前是减号", "2-(5-3)", 0},
		{"括号前是减号且嵌套", "10-((2+3)-1)", 6},
		{"多层嵌套", "((1+2)+3)+4", 10},
		{"括号内单独数字", "1-(2+3-(4+5))", 5},

		// 边界：多位数
		{"大数计算", "2147483647", 2147483647},
		{"大数加减", "100000-99999", 1},

		// 边界：空格穿插
		{"运算符两侧空格", " 3 + 5 - 2 ", 6},
		{"括号两侧空格", "( 1 + 2 ) - 3", 0},

		// 综合
		{"综合长表达式", "  (  12 + 8 )-( 7 - ( 3 + 2 ) )  ", 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.s); got != tt.want {
				t.Errorf("Calculate(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	b.Run("官方示例3_嵌套括号", func(b *testing.B) {
		s := "(1+(4+5+2)-3)+(6+8)"
		for i := 0; i < b.N; i++ {
			Calculate(s)
		}
	})
	b.Run("长表达式_长度300000", func(b *testing.B) {
		s := generateExpression(100000)
		for i := 0; i < b.N; i++ {
			Calculate(s)
		}
	})
}

// generateExpression 生成长度约为 3n 的表达式 "1+2+3+..."
func generateExpression(n int) string {
	buf := make([]byte, 0, 3*n)
	for i := 1; i <= n; i++ {
		buf = append(buf, byte('0'+i%10))
		buf = append(buf, '+')
		if i%100 == 0 {
			buf = append(buf, ' ')
		}
	}
	return string(buf)
}
