# 211. 添加与搜索单词 - 数据结构设计 (Design Add and Search Words Data Structure)

## 题目描述
请你设计一个数据结构，支持添加新单词和查找字符串是否与任何先前添加的字符串匹配。

实现词典类 `WordDictionary`：
- `WordDictionary()` 初始化词典对象
- `void addWord(word)` 将 word 添加到数据结构中，之后可以对它进行匹配
- `bool search(word)` 如果数据结构中存在字符串与 word 匹配，返回 true；否则返回 false。word 中可能包含一些 '.'，每个 '.' 可以表示任何字母。

## 示例
```
输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]
```

## 约束
- 1 <= word.length <= 25
- word 由小写英文字母或 '.' 组成
- 最多调用 10^4 次 addWord 和 search

## 分类
数据结构 / 前缀树 / DFS
