# 692. 前K个高频单词 (Top K Frequent Words)

## 题目描述
给定一个单词列表 words 和一个整数 k，返回前 k 个出现次数最多的单词。

返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字典顺序排序。

## 示例
```
输入：words = ["i","love","leetcode","i","love","coding"], k = 2
输出：["i","love"]
解释："i" 和 "love" 为出现次数最多的两个单词，均为 2 次。按字典序 "i" 在 "love" 之前。

输入：words = ["the","day","is","sunny","the","the","the","sunny","is","is"], k = 4
输出：["the","is","sunny","day"]
```

## 约束
- 1 <= words.length <= 500
- 1 <= words[i].length <= 10
- words[i] 由小写英文字母组成
- k 的取值范围是 [1, 不同 words[i] 的数量]

## 分类
前缀树 / 哈希表 / 堆 / 排序
