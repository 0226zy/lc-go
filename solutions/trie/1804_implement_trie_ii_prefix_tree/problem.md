# 1804. 实现 Trie II (前缀树 II) (Implement Trie II (Prefix Tree II))

## 题目描述
Trie 是一棵前缀树，包含 insert、countWordsEqualTo、countWordsStartingWith 和 erase 四种操作。

实现 Trie 类：
- `Trie()` 初始化前缀树对象。
- `void insert(String word)` 将 word 插入前缀树。
- `int countWordsEqualTo(String word)` 返回前缀树中等于 word 的字符串实例个数。
- `int countWordsStartingWith(String prefix)` 返回前缀树中以 prefix 为前缀的字符串实例个数。
- `void erase(String word)` 从前缀树中移除 word 的一个实例。

## 示例
```
输入：
["Trie", "insert", "insert", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsEqualTo", "countWordsStartingWith", "erase", "countWordsStartingWith"]
[[], ["apple"], ["apple"], ["apple"], ["app"], ["apple"], ["apple"], ["app"], ["apple"], ["app"]]
输出：[null, null, null, 2, 2, null, 1, 1, null, 0]
```

## 约束
- 1 <= word.length, prefix.length <= 2000
- word 和 prefix 仅由小写英文字母组成
- 最多调用 10^4 次 insert、countWordsEqualTo、countWordsStartingWith 和 erase
- 保证 erase 调用时 word 在前缀树中至少存在一个实例

## 分类
数据结构 / 前缀树
