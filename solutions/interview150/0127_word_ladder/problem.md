# 127. 单词接龙 (Word Ladder)

## 题目描述

字典 `wordList` 中从单词 `beginWord` 到 `endWord` 的**转换序列**是一个按下列规格形成的序列 `beginWord -> s1 -> s2 -> ... -> sk`：

- 序列中第一个单词是 `beginWord`，最后一个单词是 `endWord`；
- 每一对相邻的单词**只差一个字母**；
- 对于 `1 <= i <= k`，每个 `si` 都在 `wordList` 中。注意 `beginWord` 不需要在 `wordList` 中；
- `sk == endWord`。

给你 `beginWord`、`endWord` 和 `wordList`，返回从 `beginWord` 到 `endWord` 的**最短转换序列中的单词数目**。如果不存在这样的转换序列，返回 `0`。

### 示例 1

```
输入: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
输出: 5
解释: 一个最短转换序列为 "hit" -> "hot" -> "dot" -> "dog" -> "cog"，包含 5 个单词。
```

### 示例 2

```
输入: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
输出: 0
解释: endWord "cog" 不在字典中，无法转换。
```

## 提示

- `1 <= beginWord.length <= 10`
- `endWord.length == beginWord.length`
- `1 <= wordList.length <= 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`、`endWord`、`wordList[i]` 由小写英文字母组成
- `beginWord != endWord`
- `wordList` 中的所有字符串**互不相同**

## 题目解析

### 核心思路

最短转换序列 = **无权图最短路**，主体是 BFS。难点在于：如何高效地找出一个单词"只差一个字母"的所有邻居？

暴力做法是每对单词两两比较（O(n² × L)），n 最大 5000，会超时。经典的技巧是**通配符中间态建图**：

> 对单词 `hot`，它所有的"只差一个字母"的邻居，一定出现在 `*ot`、`h*t`、`ho*` 这三个模式组里。

于是预处理阶段把每个单词按"挖去第 i 个字母"的模式分组（如 `hot`、`dot`、`lot` 都进 `*ot` 组）；BFS 时对一个单词枚举它的 L 个模式，取每组里的所有单词，就是全部合法邻居。这样建图 + BFS 的总复杂度是 O(n × L²)，轻松通过。

答案要的是**单词数目**（含 `beginWord` 和 `endWord`），所以深度从 1 开始计数，碰到终点直接返回当前深度。

两个剪枝/特判：

- `endWord` 不在 `wordList` 中直接返回 0（规则要求中间词必须在字典中，`endWord` 作为 `sk` 也必须在）；
- 已访问集合必须加，否则字典里成环会死循环。

### 算法步骤

1. 把 `wordList` 放入哈希集合，特判 `endWord` 不存在则返回 0。
2. 预处理：对每个单词枚举每个位置，生成 `w[:i] + "*" + w[i+1:]` 模式，建立 `模式 -> 单词列表` 的映射。
3. BFS：队列放入 `beginWord`，深度 `depth = 1`。
4. 逐层处理：取出单词，若等于 `endWord` 返回 `depth`；否则枚举它的 L 个模式，把每组中未访问的单词标记并入队。
5. 一层结束 `depth++`；队列耗尽返回 0。

### 复杂度分析

- **时间复杂度**: O(n × L²)，n 为单词数，L 为单词长度。预处理 O(n × L) 个模式（每个模式复制字符串 O(L)）；BFS 中每个单词被访问一次，每次扩展 O(L) 个模式
- **空间复杂度**: O(n × L)，模式映射表与 visited 集合

## 代码实现

```go
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool, len(wordList))
	for _, w := range wordList {
		wordSet[w] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	// 预处理：把所有单词按"去掉第 i 个字母"的通配符模式分组
	patternMap := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			pattern := w[:i] + "*" + w[i+1:]
			patternMap[pattern] = append(patternMap[pattern], w)
		}
	}

	visited := map[string]bool{beginWord: true}
	queue := []string{beginWord}
	depth := 1 // beginWord 本身算第 1 个单词

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return depth
			}
			for j := 0; j < len(word); j++ {
				pattern := word[:j] + "*" + word[j+1:]
				for _, next := range patternMap[pattern] {
					if !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		depth++
	}
	return 0
}
```

**执行过程示例**（示例 1，模式组部分节选）：

```
模式组: *ot -> [hot, dot, lot]   *og -> [dog, log, cog]   h*t -> [hot] ...

初始: queue=[hit], depth=1
第 1 层: hit 的模式 *it, h*t, hi* 在映射中命中的邻居: hot，入队
第 2 层: hot 命中 *ot 组 -> dot, lot 入队
第 3 层: dot -> *og 组命中 dog；lot -> 命中 log
第 4 层: dog / log -> *og? cog 经 d*g/l*g? cog 在 *og 组，由 dog、log 扩展得到，入队
第 5 层: 取出 cog == endWord，返回 5
```
