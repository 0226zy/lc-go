package implementtrieprefixtree

import "testing"

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		ops  []string
		args []string
		want []interface{}
	}{
		{
			name: "示例1：基本操作",
			ops:  []string{"insert", "search", "search", "startsWith", "insert", "search"},
			args: []string{"apple", "apple", "app", "app", "app", "app"},
			want: []interface{}{nil, true, false, true, nil, true},
		},
		{
			name: "空字符串插入",
			ops:  []string{"insert", "search"},
			args: []string{"", ""},
			want: []interface{}{nil, true},
		},
		{
			name: "单个字符",
			ops:  []string{"insert", "search", "startsWith"},
			args: []string{"a", "a", "a"},
			want: []interface{}{nil, true, true},
		},
		{
			name: "前缀查找不存在",
			ops:  []string{"insert", "startsWith"},
			args: []string{"hello", "hella"},
			want: []interface{}{nil, false},
		},
		{
			name: "重复插入",
			ops:  []string{"insert", "insert", "search"},
			args: []string{"test", "test", "test"},
			want: []interface{}{nil, nil, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := Constructor()
			for i, op := range tt.ops {
				switch op {
				case "insert":
					trie.Insert(tt.args[i])
					if tt.want[i] != nil {
						t.Errorf("Insert(%q) 期望没有返回值", tt.args[i])
					}
				case "search":
					got := trie.Search(tt.args[i])
					want := tt.want[i].(bool)
					if got != want {
						t.Errorf("Search(%q) = %v, want %v", tt.args[i], got, want)
					}
				case "startsWith":
					got := trie.StartsWith(tt.args[i])
					want := tt.want[i].(bool)
					if got != want {
						t.Errorf("StartsWith(%q) = %v, want %v", tt.args[i], got, want)
					}
				}
			}
		})
	}
}

func TestTrieInsertSearch(t *testing.T) {
	t.Skip("未实现")
	// TODO: 实现后取消跳过
	if false {
		trie := Constructor()
		trie.Insert("hello")
		if !trie.Search("hello") {
			t.Error("Search('hello') 应为 true")
		}
		if trie.Search("hell") {
			t.Error("Search('hell') 应为 false（只是前缀不是单词）")
		}
		if !trie.StartsWith("hell") {
			t.Error("StartsWith('hell') 应为 true")
		}
		if trie.StartsWith("world") {
			t.Error("StartsWith('world') 应为 false")
		}
	}
}

func BenchmarkTrie(b *testing.B) {
	trie := Constructor()
	words := []string{"apple", "application", "apply", "app", "appstore"}
	for _, w := range words {
		trie.Insert(w)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Search("apple")
		trie.StartsWith("app")
	}
}
