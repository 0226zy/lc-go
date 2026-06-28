package longestwordindictionary

import "testing"

func TestLongestWord(t *testing.T) {
	t.Skip("未实现")

	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "示例1",
			words: []string{"w","wo","wor","worl","world"},
			want:  "world",
		},
		{
			name:  "示例2",
			words: []string{"a","banana","app","appl","ap","apply","apple"},
			want:  "apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestWord(tt.words); got != tt.want {
				t.Errorf("LongestWord() = %q, want %q", got, tt.want)
			}
		})
	}
}
