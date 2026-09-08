# 211. 添加与搜索单词 - 数据结构设计 (Design Add and Search Words Data Structure)

## 题目描述

请你设计一个数据结构，支持「添加新单词」和「查找是否有任意一个单词与指定模式匹配」。

实现 `WordDictionary` 类：

- `WordDictionary()` 初始化数据结构。
- `void addWord(word)` 将 `word` 添加到数据结构中。
- `bool search(word)` 如果数据结构中存在字符串与 `word` 匹配，则返回 `true`；否则返回 `false`。

`word` 中可能包含点号 `'.'`，每个 `.` 可以匹配**任意一个**字母。

### 示例 1

```
输入:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出:
[null,null,null,null,false,true,true,true]

解释:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True（可匹配 bad、dad、mad）
wordDictionary.search("b.."); // 返回 True（可匹配 bad）
```

## 提示

- `1 <= word.length <= 25`
- `addWord` 中的 `word` 由小写英文字母组成
- `search` 中的 `word` 由 `.` 或小写英文字母组成
- 最多调用 `10^4` 次 `addWord` 和 `search`

## 题目解析

### 核心思路

`addWord` 和普通的 `search` 都是前缀树（Trie）的看家本领，和 208 题一样。本题唯一的增量是 `.` **通配符**：到某一位不知道具体字符，需要**尝试所有可能的分支**——这是典型的回溯场景。

- 普通字符：和 208 一样，唯一对应一个子节点，槽位为空直接失败；
- `.`：当前节点下所有非空子节点都可能是答案，逐个递归尝试，**任意一个分支成功就算成功**。

因为不确定走哪个分支，不能像普通 search 那样迭代着一路下钻，需要把"剩余字符串 + 当前节点"作为递归状态写回溯函数。

### 算法步骤

**addWord(word)**：同 208 题，逐字符下钻建节点，结尾标 `isEnd`。

**search(word)**：
1. 从根节点、第 0 个字符开始调用回溯函数 `searchFrom`；
2. `searchFrom(node, i)`：
   - 若 `i == len(word)`，已走完，返回 `node.isEnd`（必须落在完整单词上）；
   - 若 `word[i] == '.'`，遍历 `node.children` 中所有非空子节点，递归 `searchFrom(child, i+1)`，任一返回 `true` 即成功；
   - 否则取对应槽位子节点，为空返回 `false`，否则递归。

### 复杂度分析

设 L 为单词长度，调用次数为 q：

- **addWord**: 时间 O(L)，空间 O(L)
- **search**: 没有 `.` 时时间 O(L)；最坏情况（全是 `.`）需要遍历整棵树，时间 O(26^L)（实际受字典树规模限制，为 O(树中节点数)），递归栈空间 O(L)

## 代码实现

```go
// WordDictionary 支持添加单词与按模式搜索的词典
type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (d *WordDictionary) AddWord(word string) {
	node := d
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionary{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (d *WordDictionary) Search(word string) bool {
	return d.searchFrom(word, 0)
}

// searchFrom 从第 index 个字符开始在当前子树下回溯搜索
func (d *WordDictionary) searchFrom(word string, index int) bool {
	if index == len(word) {
		return d.isEnd
	}
	ch := word[index]
	if ch == '.' {
		// 通配符：尝试所有存在的子分支
		for _, child := range d.children {
			if child != nil && child.searchFrom(word, index+1) {
				return true
			}
		}
		return false
	}
	child := d.children[ch-'a']
	if child == nil {
		return false
	}
	return child.searchFrom(word, index+1)
}
```

**执行过程示例**（词典中已有 `bad`、`dad`、`mad`，搜索 `"b.."`）：

```
Search("b.."):
  第 0 位 'b': 沿 b 分支下钻
  第 1 位 '.': 尝试 b 下的所有子分支: a 分支存在
       -> 递归剩余 "."
  第 2 位 '.': 在 ba 下尝试所有子分支: d 分支存在
       -> 递归剩余 ""
  走完: 落点 isEnd=true（bad 在词典中），返回 true
```
