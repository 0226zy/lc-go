package implementmagicdictionary

import "testing"

func TestMagicDictionary(t *testing.T) {
	t.Skip("未实现")

	d := Constructor()
	d.BuildDict([]string{"hello", "leetcode"})

	tests := []struct {
		word string
		want bool
	}{
		{"hello", false},
		{"hhllo", true},
		{"hell", false},
		{"leetcoded", false},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := d.Search(tt.word); got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
