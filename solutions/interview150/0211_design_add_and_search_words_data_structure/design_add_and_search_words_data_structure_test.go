package designaddandsearchwordsdatastructure

import "testing"

func TestWordDictionary(t *testing.T) {
	d := Constructor()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")

	tests := []struct {
		name string
		word string
		want bool
	}{
		{"普通词不存在", "pad", false},
		{"普通词存在", "bad", true},
		{"前缀通配符", ".ad", true},
		{"后缀通配符", "b..", true},
		{"中段通配符", "b.d", true},
		{"长度不匹配", "ba", false},
		{"超长不匹配", "bade", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := d.Search(tt.word); got != tt.want {
				t.Errorf("Search(%q) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestWordDictionaryEdgeCases(t *testing.T) {
	t.Run("空词典", func(t *testing.T) {
		d := Constructor()
		if d.Search("a") {
			t.Errorf("Search(a) = true, want false")
		}
	})

	t.Run("全是通配符", func(t *testing.T) {
		d := Constructor()
		d.AddWord("abc")
		if !d.Search("...") {
			t.Errorf("Search(...) = false, want true")
		}
		if d.Search("....") {
			t.Errorf("Search(....) = true, want false（长度不符）")
		}
	})

	t.Run("重复添加", func(t *testing.T) {
		d := Constructor()
		d.AddWord("a")
		d.AddWord("a")
		if !d.Search("a") {
			t.Errorf("Search(a) = false, want true")
		}
	})

	t.Run("前缀是完整单词", func(t *testing.T) {
		d := Constructor()
		d.AddWord("app")
		d.AddWord("apple")
		if !d.Search("app") || !d.Search("apple") {
			t.Errorf("Search(app/apple) 应均为 true")
		}
		if !d.Search("ap.le") {
			t.Errorf("Search(ap.le) = false, want true")
		}
	})

	t.Run("通配符不匹配单词结尾", func(t *testing.T) {
		d := Constructor()
		d.AddWord("ab")
		if !d.Search("a.") {
			t.Errorf("Search(a.) = false, want true")
		}
		if d.Search("a..") {
			t.Errorf("Search(a..) = true, want false")
		}
	})
}

func BenchmarkWordDictionaryAddWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := Constructor()
		for _, w := range []string{"bad", "dad", "mad", "apple", "application"} {
			d.AddWord(w)
		}
	}
}

func BenchmarkWordDictionarySearch(b *testing.B) {
	d := Constructor()
	for _, w := range []string{"bad", "dad", "mad", "apple", "application"} {
		d.AddWord(w)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.Search(".ad")
		d.Search("a.p.e")
	}
}
