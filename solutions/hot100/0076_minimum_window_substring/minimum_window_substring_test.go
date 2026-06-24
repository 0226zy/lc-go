package minimumwindowsubstring

import "testing"

func TestMinWindow(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "示例1：基本用例",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "示例2：单字符匹配",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "示例3：无法覆盖",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "s和t相同",
			s:    "abc",
			t:    "abc",
			want: "abc",
		},
		{
			name: "子串在中间",
			s:    "xxabcxx",
			t:    "abc",
			want: "abc",
		},
		{
			name: "重复字符",
			s:    "aaabbbccc",
			t:    "abc",
			want: "abbbc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinWindow(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("MinWindow(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkMinWindow(b *testing.B) {
	s := "ADOBECODEBANC"
	t := "ABC"
	for i := 0; i < b.N; i++ {
		MinWindow(s, t)
	}
}
