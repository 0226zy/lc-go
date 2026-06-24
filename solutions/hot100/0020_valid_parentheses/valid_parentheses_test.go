package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "示例1：单对括号",
			s:    "()",
			want: true,
		},
		{
			name: "示例2：多种括号",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "示例3：不匹配",
			s:    "(]",
			want: false,
		},
		{
			name: "嵌套括号",
			s:    "{[()]}",
			want: true,
		},
		{
			name: "错误顺序",
			s:    "([)]",
			want: false,
		},
		{
			name: "只有左括号",
			s:    "(((",
			want: false,
		},
		{
			name: "只有右括号",
			s:    ")))",
			want: false,
		},
		{
			name: "空字符串",
			s:    "",
			want: true,
		},
		{
			name: "单个字符",
			s:    "(",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.s)
			if got != tt.want {
				t.Errorf("IsValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	s := "{[()]}"
	for i := 0; i < b.N; i++ {
		IsValid(s)
	}
}
