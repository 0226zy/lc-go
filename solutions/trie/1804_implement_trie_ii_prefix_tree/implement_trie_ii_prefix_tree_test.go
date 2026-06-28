package implementtrieiiprefixtree

import "testing"

func TestTrie(t *testing.T) {
	t.Skip("未实现")

	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("apple")
	if got := trie.CountWordsEqualTo("apple"); got != 2 {
		t.Errorf("CountWordsEqualTo(apple) = %d, want 2", got)
	}
	if got := trie.CountWordsStartingWith("app"); got != 2 {
		t.Errorf("CountWordsStartingWith(app) = %d, want 2", got)
	}
	trie.Erase("apple")
	if got := trie.CountWordsEqualTo("apple"); got != 1 {
		t.Errorf("CountWordsEqualTo(apple) after erase = %d, want 1", got)
	}
	if got := trie.CountWordsStartingWith("app"); got != 1 {
		t.Errorf("CountWordsStartingWith(app) after erase = %d, want 1", got)
	}
	trie.Erase("apple")
	if got := trie.CountWordsStartingWith("app"); got != 0 {
		t.Errorf("CountWordsStartingWith(app) after second erase = %d, want 0", got)
	}
}
