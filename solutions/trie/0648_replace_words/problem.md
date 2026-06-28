# 648. 单词替换 (Replace Words)

## 题目描述
在英语中，我们有一个叫做词根的概念，可以词根后面添加其他一些词组成另一个较长的单词。我们称这个词为继承词。例如，词根 help，随着其后加上 ful，就形成了 helpful。

给定一个由许多词根组成的词典 dictionary 和一个用空格分隔句子形成的句子 sentence。你需要将句子中的所有继承词替换成它的词根。如果继承词有许多可以形成它的词根，则用最短的词根替换它。

## 示例
```
输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"
```

## 约束
- 1 <= dictionary.length <= 1000
- 1 <= dictionary[i].length <= 100
- dictionary[i] 仅由小写字母组成
- 1 <= sentence.length <= 10^6
- sentence 仅由小写字母和空格组成
- sentence 中单词的个数为 [1, 1000]

## 分类
前缀树 / 哈希表
