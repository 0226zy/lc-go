# 734. 句子相似性

> 难度：简单 ｜ 分类：哈希 ｜ 尊享面试 100 题 · 第 14 题
> 链接：https://leetcode.cn/problems/sentence-similarity/

## 题目描述

一个句子用一个字符串数组表示，其中每个元素是一个单词，句子中的单词之间以空格分隔。

给定两个句子 `sentence1` 和 `sentence2`（均为字符串数组），以及一个字符串对列表 `similarPairs`，其中每个元素 `[xi, yi]` 表示单词 `xi` 和 `yi` 是**相似**的。

如果两个句子满足以下条件，则称它们是**相似**的：

- 两个句子包含的单词数量相同（长度相等）；
- 对于每个下标 `i`，`sentence1[i]` 与 `sentence2[i]` **相同**，或者二者是相似的单词。

注意：相似关系是**双向**的（`[x, y]` 意味着 `x` 与 `y` 相似、`y` 也与 `x` 相似），但**不具有传递性**（`x` 与 `y` 相似、`y` 与 `z` 相似，并不能推出 `x` 与 `z` 相似）。每个单词总是与它自身相似。

请判断两个句子是否相似，相似则返回 `true`，否则返回 `false`。

### 示例 1

```
输入: sentence1 = ["great","acting","skills"],
     sentence2 = ["fine","drama","talent"],
     similarPairs = [["great","fine"],["drama","acting"],["skills","talent"]]
输出: true
解释: 两个句子长度相同，且每个位置上的单词都是相似的。
```

### 示例 2

```
输入: sentence1 = ["great"],
     sentence2 = ["great"],
     similarPairs = []
输出: true
解释: 单词总是与它自身相似。
```

### 示例 3

```
输入: sentence1 = ["great"],
     sentence2 = ["doubleplus","good"],
     similarPairs = [["great","doubleplus"]]
输出: false
解释: 两个句子长度不同，直接判定不相似。
```

### 提示

- `1 <= sentence1.length, sentence2.length <= 1000`
- `1 <= sentence1[i].length, sentence2[i].length <= 20`
- `sentence1[i]` 和 `sentence2[i]` 仅由大小写英文字母组成
- `0 <= similarPairs.length <= 2000`
- `similarPairs[i].length == 2`
- `1 <= xi.length, yi.length <= 20`
- `xi` 和 `yi` 仅由英文字母组成
- 所有单词对 `(xi, yi)` 互不相同

## 思路解析

### 核心思路

这是一道典型的**哈希表**应用题。判断条件是逐位置独立的，所以：

1. 先比较两个句子的长度，不同则直接返回 `false`；
2. 把 `similarPairs` 建成「单词 → 相似单词集合」的哈希表（由于相似是双向的，两个方向都要插入）；
3. 逐位置比较：对应单词相同，或者在彼此的相似集合中，才算匹配；只要有一个位置不匹配就返回 `false`。

需要注意两个坑：

- **方向性**：词对 `[x, y]` 表示双向相似，建表时 `x -> y` 和 `y -> x` 都要加入，否则查 `y` 的相似词时会漏掉 `x`；
- **不可传递**：不要试图用并查集把相似词合并成连通块，题目明确说明相似关系不能传递，只能查「直接相似」。

### 算法步骤

1. 若 `len(sentence1) != len(sentence2)`，返回 `false`。
2. 建立哈希表 `sim`，类型为 `map[string]map[string]bool`；遍历 `similarPairs`，把 `pair[1]` 加入 `sim[pair[0]]` 的集合，同时把 `pair[0]` 加入 `sim[pair[1]]` 的集合。
3. 对每个下标 `i`：
   - 若 `sentence1[i] == sentence2[i]`，继续；
   - 否则检查 `sentence2[i]` 是否在 `sim[sentence1[i]]` 中，不在则返回 `false`。
4. 所有位置都匹配，返回 `true`。

## 复杂度分析

- **时间复杂度**: O(n + p · L)，其中 n 为句子长度，p 为相似词对数量，L 为单词平均长度（哈希表插入/查询的代价与字符串长度成正比）。
- **空间复杂度**: O(p · L)，哈希表最多存储 2p 个单词到集合的映射项。
