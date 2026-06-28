package designaddandsearchwordsdatastructure

import "testing"

func TestWordDictionary(t *testing.T) {
	t.Skip("未实现")

	d := Constructor()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")

	tests := []struct {
		word string
		want bool
	}{
		{"pad", false},
		{"bad", true},
		{".ad", true},
		{"b..", true},
	}

	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := d.Search(tt.word); got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}
