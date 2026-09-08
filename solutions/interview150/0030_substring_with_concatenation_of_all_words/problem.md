# 30. 串联所有单词的子串 (Substring with Concatenation of All Words)

## 题目描述

给定一个字符串 `s` 和一个字符串数组 `words`，`words` 中**所有单词长度相同**。

`s` 中的**串联子串**是指：一个恰好包含 `words` 中**所有单词**（每个单词按 `words` 中出现的次数使用，顺序任意）依次拼接而成的子串。

返回所有串联子串在 `s` 中的**起始下标**。答案顺序不限。

### 示例 1

```
输入: s = "barfoothefoobarman", words = ["foo","bar"]
输出: [0,9]
解释: 从索引 0 开始的子串是 "barfoo"，包含 ["bar","foo"]；从索引 9 开始的子串是 "foobar"，包含 ["foo","bar"]。
```

### 示例 2

```
输入: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出: []
解释: words 中有两个 "word"，但 s 中没有子串能按次数覆盖它们（中间 "goodgoodgood" 无法匹配）。
```

### 示例 3

```
输入: s = "barfoofoobarthefoobarman", words = ["the","bar","foo"]
输出: [6,9,12]
解释: 从 6 开始 "foobarthe"，从 9 开始 "barthefoo"，从 12 开始 "thefoobar"。
```

### 提示

- `1 <= s.length <= 10^4`
- `1 <= words.length <= 5000`
- `1 <= words[i].length <= 30`
- `s` 和 `words[i]` 由小写英文字母组成

## 题目解析

### 核心思路

这道题可以看成 [3. 无重复字符的最长子串](../0003_longest_substring_without_repeating_characters/problem.md) 和 [76. 最小覆盖子串](../0076_minimum_window_substring/problem.md) 的「复合版」：覆盖目标不再是字符，而是**一组定长单词**。

直接套用滑窗会遇到一个陷阱：窗口按什么步长滑动？如果把 `s` 按**每个单词的长度 `wordLen`** 切成一段一段的，那么任何合法答案的起点必然和其中一种切分方式对齐（起点模 `wordLen` 的余数只有 `0 ~ wordLen-1` 这 `wordLen` 种）。所以外层枚举偏移 `i`（`0` 到 `wordLen-1`），**内层固定从 `i` 开始按 `wordLen` 步长滑窗**，就保证窗口边界永远切在单词缝隙上，不会把单词切成两半。

固定偏移后，问题就退化成经典的「**定长单元滑窗 + 哈希计数**」：

- 用一个 `need` 表记录每个单词需要出现几次；
- 窗口内用 `have` 表统计已收集的单词次数，同时维护窗口内有效单词个数 `count`；
- 新切出的片段不在 `need` 里 → 前面整段都作废，窗口清零，从下一片段重新开始；
- 新片段导致某单词**超标** → 从左边按单词粒度收缩，直到该单词次数回到合法范围；
- `count` 恰好等于 `words` 长度 → 当前窗口就是一个串联子串，记录左边界。

因为单词顺序任意、只看计数，所以 `foobar` 和 `barfoo` 都算合法。

### 算法步骤

1. 统计 `need`：遍历 `words`，`need[w]++`；记 `wordLen = len(words[0])`，`wordCount = len(words)`。
2. 枚举起点偏移 `i`（`0 <= i < wordLen`），对每个 `i` 执行一次滑窗：
   - `left = i`，`count = 0`，清空 `have`；
   - `j` 从 `i` 开始、每次跳 `wordLen`，取片段 `w = s[j:j+wordLen]`：
     - 若 `need[w] == 0`：`have`、`count` 清零，`left = j + wordLen`（窗口作废）；
     - 否则 `have[w]++`、`count++`；若 `have[w] > need[w]`，循环从左边移出单词（`left += wordLen`）直到不超标；
     - 若 `count == wordCount`，把 `left` 加入答案。
3. 返回答案列表。

### 复杂度分析

- **时间复杂度**: O(len(s) × wordLen)。对每个偏移 `i`，内层扫描 s 的每个字符一次，每个片段的哈希比较代价为 O(wordLen)，共 wordLen 个偏移。
  由于每段只被切一次、比较一次，总代价也可记为 O(len(s) × wordLen)；当 wordLen 很小时接近 O(len(s))。
- **空间复杂度**: O(wordCount × wordLen)，哈希表最多存 words 中所有单词。

## 代码实现

```go
func FindSubstring(s string, words []string) []int {
	n := len(s)
	wordLen := len(words[0])
	wordCount := len(words)
	if n < wordLen*wordCount {
		return nil
	}

	need := make(map[string]int, wordCount)
	for _, w := range words {
		need[w]++
	}

	ans := []int{}
	// 按单词长度 wordLen 分组滑动窗口：起点偏移 i 取 0 ~ wordLen-1，
	// 保证同一组内所有切分点对齐，不会错位切词
	for i := 0; i < wordLen; i++ {
		left := i                 // 窗口左边界（下标）
		count := 0                // 窗口内有效单词个数
		have := make(map[string]int, wordCount)
		for j := i; j+wordLen <= n; j += wordLen {
			w := s[j : j+wordLen]
			if need[w] > 0 {
				have[w]++
				count++
				// 若该单词出现次数超标，从左边收缩，直到次数达标
				for have[w] > need[w] {
					lw := s[left : left+wordLen]
					have[lw]--
					left += wordLen
					count--
				}
				if count == wordCount {
					ans = append(ans, left)
				}
			} else {
				// 切出的片段不在 words 中，窗口直接清零，从下一片段重新开始
				have = make(map[string]int, wordCount)
				count = 0
				left = j + wordLen
			}
		}
	}
	return ans
}
```

**执行过程示例**（`s = "barfoofoobarthefoobarman"`, `words = ["the","bar","foo"]`，`wordLen = 3`）：

以偏移 `i = 0` 为例，切分：bar|foo|foo|bar|the|foo|bar|man：

```
j=0  "bar": have={bar:1}, count=1
j=3  "foo": have={bar:1,foo:1}, count=2
j=6  "foo": have[foo]=2 > need 1 → 左移 "bar"(left:0→3), have[bar]=0,
      仍超标 → 左移 "foo"(left:3→6), have[foo]=1; count=1
j=9  "bar": have={bar:1,foo:1}, count=2
j=12 "the": have={bar:1,foo:1,the:1}, count=3 == wordCount → 记录 left=6
j=15 "foo": 加入后 count=4, 无超标（各次数均 ≤ 1）→ count != 3, 继续
j=18 "bar": 加入后 have={bar:2,...} 超标 → 左移 "the"(left:9), count=3
      → have={bar:2,foo:2} 仍超标 → 左移 "foo"(left:12), count=2
      → have[bar]=2 > 1 → 左移 "bar"(left:15), count=1, have[bar]=1
j=21 "man": 不在 need 中 → 窗口清零, left=24
```
最终偏移 0 这一组贡献答案 6；偏移 1、2 两组分别贡献 9、12，合计 `[6,9,12]`。
