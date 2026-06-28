package longestvalidparentheses

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "示例1：基本用例",
			s:    "(()",
			want: 2,
		},
		{
			name: "示例2：多个有效段",
			s:    ")()())",
			want: 4,
		},
		{
			name: "示例3：空字符串",
			s:    "",
			want: 0,
		},
		{
			name: "全部有效",
			s:    "()()()",
			want: 6,
		},
		{
			name: "嵌套有效",
			s:    "((()))",
			want: 6,
		},
		{
			name: "全部无效",
			s:    ")))(((",
			want: 0,
		},
		{
			name: "单对括号",
			s:    "()",
			want: 2,
		},
		{
			name: "交错括号",
			s:    "()(()",
			want: 2,
		},
		{
			name: "长有效串",
			s:    "()()((()))()",
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LongestValidParentheses(tt.s)
			if got != tt.want {
				t.Errorf("LongestValidParentheses(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestValidParentheses(b *testing.B) {
	s := ")()())"
	for i := 0; i < b.N; i++ {
		LongestValidParentheses(s)
	}
}
