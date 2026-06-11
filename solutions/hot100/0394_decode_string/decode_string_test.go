package decodestring

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "示例1",
			s:    "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			name: "示例2-嵌套",
			s:    "3[a2[c]]",
			want: "accaccacc",
		},
		{
			name: "示例3-混合",
			s:    "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			name: "单层重复",
			s:    "10[a]",
			want: "aaaaaaaaaa",
		},
		{
			name: "多层嵌套",
			s:    "2[3[a]b]",
			want: "aaabaaab",
		},
		{
			name: "无编码字符串",
			s:    "abc",
			want: "abc",
		},
		{
			name: "纯数字括号",
			s:    "3[z]",
			want: "zzz",
		},
		{
			name: "数字大于10",
			s:    "12[a]",
			want: "aaaaaaaaaaaa",
		},
		{
			name: "三重嵌套",
			s:    "2[2[2[b]]]",
			want: "bbbbbbbb",
		},
		{
			name: "前后都有普通字符",
			s:    "abc3[cd]xyz",
			want: "abccdcdcdxyz",
		},
		{
			name: "连续嵌套",
			s:    "2[a]2[b]2[c]",
			want: "aabbcc",
		},
		{
			name: "大数字",
			s:    "100[a]",
			want: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DecodeString(tt.s)
			if got != tt.want {
				t.Errorf("DecodeString() = %q, want %q", got, tt.want)
			}
		})
	}
}

func BenchmarkDecodeString(b *testing.B) {
	s := "3[a2[c]]"
	for i := 0; i < b.N; i++ {
		_ = DecodeString(s)
	}
}
