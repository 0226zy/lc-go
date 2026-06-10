package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "空字符串",
			s:    "",
			want: true,
		},
		{
			name: "官方示例1：()",
			s:    "()",
			want: true,
		},
		{
			name: "官方示例2：()[]{}",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "官方示例3：(]",
			s:    "(]",
			want: false,
		},
		{
			name: "多层嵌套：([{}])",
			s:    "([{}])",
			want: true,
		},
		{
			name: "左括号多余：(()",
			s:    "(()",
			want: false,
		},
		{
			name: "右括号多余：())",
			s:    "()",
			want: true,
		},
		{
			name: "错误顺序：([)]",
			s:    "([)]",
			want: false,
		},
		{
			name: "单对括号：()",
			s:    "()",
			want: true,
		},
		{
			name: "只有左括号：(((",
			s:    "(((",
			want: false,
		},
		{
			name: "只有右括号：)))",
			s:    ")))",
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

func TestIsValidFast(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "空字符串",
			s:    "",
			want: true,
		},
		{
			name: "官方示例1：()",
			s:    "()",
			want: true,
		},
		{
			name: "官方示例2：()[]{}",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "官方示例3：(]",
			s:    "(]",
			want: false,
		},
		{
			name: "多层嵌套：([{}])",
			s:    "([{}])",
			want: true,
		},
		{
			name: "左括号多余：(()",
			s:    "(()",
			want: false,
		},
		{
			name: "错误顺序：([)]",
			s:    "([)]",
			want: false,
		},
		{
			name: "奇数长度：({[",
			s:    "({[",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidFast(tt.s)
			if got != tt.want {
				t.Errorf("IsValidFast(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestIsValidBytes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "空字符串",
			s:    "",
			want: true,
		},
		{
			name: "官方示例1：()",
			s:    "()",
			want: true,
		},
		{
			name: "官方示例2：()[]{}",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "官方示例3：(]",
			s:    "(]",
			want: false,
		},
		{
			name: "多层嵌套：([{}])",
			s:    "([{}])",
			want: true,
		},
		{
			name: "错误顺序：([)]",
			s:    "([)]",
			want: false,
		},
		{
			name: "奇数长度：({[",
			s:    "({[",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidBytes(tt.s)
			if got != tt.want {
				t.Errorf("IsValidBytes(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	s := "([{}])"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValid(s)
	}
}

func BenchmarkIsValidFast(b *testing.B) {
	s := "([{}])"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidFast(s)
	}
}

func BenchmarkIsValidBytes(b *testing.B) {
	s := "([{}])"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidBytes(s)
	}
}
