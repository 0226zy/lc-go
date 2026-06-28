# 676. 实现一个魔法字典 (Implement Magic Dictionary)

## 题目描述
设计一个使用单词列表进行初始化的数据结构，单词列表中的单词互不相同。如果能将单词中的一个字母替换成任意字母，使得替换后的单词与列表中的某个单词相同，则称该单词是魔法等价的。

实现 `MagicDictionary` 类：
- `MagicDictionary()` 初始化对象
- `void buildDict(String[] dictionary)` 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
- `bool search(String searchWord)` 给定一个字符串 searchWord，判定能否恰好将 searchWord 中的一个字母替换成另一个字母，使得替换后的字符串是数据结构中的某个字符串

## 示例
```
输入：
["MagicDictionary", "buildDict", "search", "search", "search", "search"]
[[], [["hello","leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
输出：[null,null,false,true,false,false]
```

## 约束
- 1 <= dictionary.length <= 100
- 1 <= dictionary[i].length <= 100
- dictionary[i] 仅由小写英文字母组成
- dictionary 中的所有字符串互不相同
- 1 <= searchWord.length <= 100
- searchWord 仅由小写英文字母组成
- buildDict 仅在 search 之前调用一次
- 最多调用 100 次 search

## 分类
前缀树 / 暴力枚举
