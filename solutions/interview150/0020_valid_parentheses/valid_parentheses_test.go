package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 单对圆括号", "()", true},
		{"示例2: 三种括号各自成对", "()[]{}", true},
		{"示例3: 类型不匹配", "(]", false},

		// 边界：长度限制内的典型情况
		{"嵌套匹配", "([{}])", true},
		{"多组顺序匹配", "(){}[]", true},
		{"交叉嵌套非法", "([)]", false},
		{"左括号未闭合", "(", false},
		{"右括号无匹配", "]", false},
		{"多个左括号未闭合", "(((", false},
		{"多个右括号无匹配", ")))", false},
		{"混合部分未闭合", "(((", false},
		{"开头即右括号", "}{", false},
		{"同类型未闭合", "((", false},
		{"长字符串全部匹配", "(((((((((())))))))))", true},
		{"长字符串末尾不匹配", "((((((((((())))))))", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.s); got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"短串_嵌套匹配", "([{}])"},
		{"短串_交叉非法", "([)]"},
		{"长串_全部匹配", generateBrackets(10000)},
		{"长串_左括号未闭合", generateUnclosed(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsValid(bm.s)
			}
		})
	}
}

// generateBrackets 生成 n 对完全嵌套匹配的括号字符串
func generateBrackets(n int) string {
	s := make([]byte, 0, 2*n)
	for i := 0; i < n; i++ {
		s = append(s, '(')
	}
	for i := 0; i < n; i++ {
		s = append(s, ')')
	}
	return string(s)
}

// generateUnclosed 生成长度为 n 的全左括号字符串（未闭合）
func generateUnclosed(n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = '('
	}
	return string(s)
}
