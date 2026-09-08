# 208. 实现 Trie (前缀树) (Implement Trie (Prefix Tree))

## 题目描述

**Trie**（发音类似 "try"）或者说**前缀树**是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

- `Trie()` 初始化前缀树对象。
- `void insert(String word)` 向前缀树中插入字符串 `word`。
- `boolean search(String word)` 如果字符串 `word` 在前缀树中，返回 `true`（即，在检索之前已经插入）；否则，返回 `false`。
- `boolean startsWith(String prefix)` 如果之前已经插入的字符串 `word` 的前缀之一为 `prefix`，返回 `true`；否则，返回 `false`。

### 示例 1

```
输入:
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出:
[null, null, true, false, true, null, true]

解释:
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
```

## 提示

- `1 <= word.length, prefix.length <= 2000`
- `word` 和 `prefix` 仅由小写英文字母组成
- `insert`、`search` 和 `startsWith` 调用次数总计不超过 `3 * 10^4` 次

## 题目解析

### 核心思路

前缀树的核心思想是**用公共前缀压缩存储**：从根到某个节点的路径就代表一个前缀，树上一条从根出发的路径唯一对应一个字符串。

- 每个节点有 26 个槽位（本题只有小写字母），`children[i]` 指向以 `'a'+i` 为下一个字符的子树；
- 路径上的字符隐含在"走哪个子节点"里，节点本身不必存字符；
- `isEnd` 标记当前节点是否是一个**完整单词的结尾**——这是 `search("app")` 与 `startsWith("app")` 的区别关键：`app` 可能是 `apple` 的中途节点，只有 `isEnd = true` 才算插入过这个单词。

这是"设计类"题目的模板题：Trie 会被后续很多题（单词搜索 II、添加与搜索单词等）复用，值得背熟。

### 算法步骤

**insert(word)**：
1. 从根节点出发，逐字符读取 `word[i]`，算出槽位 `idx = word[i] - 'a'`；
2. 若 `children[idx]` 为空则新建节点；
3. 下移，直到处理完所有字符，把最后节点的 `isEnd` 置 `true`。

**search(word)**：与 insert 走相同路径，任一步槽位为空则返回 `false`；走完看落点 `isEnd` 是否为 `true`。

**startsWith(prefix)**：与 search 相同，但**不要求** `isEnd`，路径能走完就返回 `true`。

三个操作都抽出一个公共的 `searchPrefix`：沿前缀逐字符下钻，返回最后一个字符对应的节点（或 nil）。

### 复杂度分析

设 L 为字符串长度：

- **insert**: 时间 O(L)，空间 O(L)（最坏情况全部是新前缀，新建 L 个节点）
- **search / startsWith**: 时间 O(L)，空间 O(1)

## 代码实现

```go
// Trie 前缀树（字典树）节点
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd // 必须是完整单词
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil // 只要求前缀存在
}

// searchPrefix 沿前缀逐字符下钻，返回最后一个字符对应的节点；前缀不存在返回 nil
func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}
```

**执行过程示例**（依次插入 `"app"`、`"apple"` 后树的形态）：

```
root
└─ a ── p ── p (isEnd)          <- 插入 "app" 到此为止，isEnd=true
             └─ l ── e (isEnd)  <- 继续插入 "apple"

Search("app"):  沿 a→p→p 走完，落点 isEnd=true  -> true
Search("appl"): 落点 isEnd=false               -> false
StartsWith("ap"): 路径能走完                   -> true
```
