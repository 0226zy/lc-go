package generateparentheses

import (
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n=1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n=2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "n=3 官方示例",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n=4",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()",
				"(()(()))", "(()()())", "(()())()", "(())(())",
				"(())()()", "()((()))", "()(()())", "()(())()",
				"()()(())", "()()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesis(tt.n)
			if !equalStringSlices(got, tt.want) {
				t.Errorf("GenerateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesisDP(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n=1",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "n=2",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "n=3 官方示例",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n=4",
			n:    4,
			want: []string{
				"(((())))", "((()()))", "((())())", "((()))()",
				"(()(()))", "(()()())", "(()())()", "(())(())",
				"(())()()", "()((()))", "()(()())", "()(())()",
				"()()(())", "()()()()",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateParenthesisDP(tt.n)
			if !equalStringSlices(got, tt.want) {
				t.Errorf("GenerateParenthesisDP(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// equalStringSlices 比较两个字符串切片是否包含相同元素（顺序无关）
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	sa := make([]string, len(a))
	copy(sa, a)
	sb := make([]string, len(b))
	copy(sb, b)
	sort.Strings(sa)
	sort.Strings(sb)
	for i := range sa {
		if sa[i] != sb[i] {
			return false
		}
	}
	return true
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParenthesis(8)
	}
}

func BenchmarkGenerateParenthesisDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParenthesisDP(8)
	}
}
