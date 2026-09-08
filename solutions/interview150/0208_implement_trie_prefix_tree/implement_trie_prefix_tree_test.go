package implementtrieprefixtree

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Errorf("Search(apple) = false, want true")
	}
	if trie.Search("app") {
		t.Errorf("Search(app) = true, want false（app 不是完整单词）")
	}
	if !trie.StartsWith("app") {
		t.Errorf("StartsWith(app) = false, want true")
	}
	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Search(app) = false, want true（插入 app 后应为 true）")
	}
}

func TestTrieEdgeCases(t *testing.T) {
	t.Run("空字典", func(t *testing.T) {
		trie := Constructor()
		if trie.Search("a") {
			t.Errorf("Search(a) = true, want false")
		}
		if trie.StartsWith("a") {
			t.Errorf("StartsWith(a) = true, want false")
		}
	})

	t.Run("单字符单词", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("a")
		if !trie.Search("a") {
			t.Errorf("Search(a) = false, want true")
		}
		if trie.StartsWith("ab") {
			t.Errorf("StartsWith(ab) = true, want false")
		}
	})

	t.Run("前缀是完整单词", func(t *testing.T) {
		trie := Constructor()
		trie.Insert("app")
		trie.Insert("apple")
		if !trie.Search("app") || !trie.Search("apple") {
			t.Errorf("Search(app/apple) 应均为 true")
		}
		if !trie.StartsWith("app") || !trie.StartsWith("appl") {
			t.Errorf("StartsWith(app/appl) 应均为 true")
		}
	})

	t.Run("同前缀多分支", func(t *testing.T) {
		trie := Constructor()
		for _, w := range []string{"apple", "apply", "apex", "banana"} {
			trie.Insert(w)
		}
		if !trie.Search("apply") || !trie.Search("apex") {
			t.Errorf("Search(apply/apex) 应均为 true")
		}
		if trie.Search("ap") || trie.Search("ban") {
			t.Errorf("Search(ap/ban) 应为 false（非完整单词）")
		}
		if !trie.StartsWith("ban") {
			t.Errorf("StartsWith(ban) 应为 true")
		}
	})
}

func BenchmarkTrieInsert(b *testing.B) {
	words := []string{"apple", "application", "apply", "apex", "banana", "band", "cat", "category"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie := Constructor()
		for _, w := range words {
			trie.Insert(w)
		}
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	words := []string{"apple", "application", "apply", "apex", "banana", "band", "cat", "category"}
	trie := Constructor()
	for _, w := range words {
		trie.Insert(w)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Search("application")
		trie.Search("apple")
	}
}
