# 49. 字母异位词分组 (Group Anagrams)

## 题目描述

给你一个字符串数组，请你将 **字母异位词** 组合在一起。可以按任意顺序返回结果列表。

**字母异位词** 是由重新排列源单词的所有字母得到的一个新单词。

### 示例 1

```
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

### 示例 2

```
输入: strs = [""]
输出: [[""]]
```

### 示例 3

```
输入: strs = ["a"]
输出: [["a"]]
```

### 提示

- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 仅包含小写字母

## 题目解析

### 核心思路

这是 **哈希分组** 模型：核心问题是——**怎么判断两个字符串互为异位词？** 上两题（242、383）已经回答了：字符多重集相同。进一步想，互为异位词的字符串有一个共同特征：**把它们内部字符排序后，结果完全相同**（`eat`、`tea`、`ate` 排序后都是 `aet`）。

于是问题变成：给每个字符串算出一个“排序后的 key”，用哈希表把 key 相同的字符串放进同一组。一次遍历就完成全部分组。

还有另一种等价思路（不用排序）：用长度为 26 的计数数组把每个字符串编码成一个 key（例如 `"1#0#2#..."` 这样把计数拼起来）。排序法代码更短、更直观，故本题采用排序法。

### 算法步骤

1. 创建哈希表 `groups`，key 为排序后的字符串，value 为原字符串列表。
2. 遍历 `strs` 中每个字符串 `s`：
   - 把 `s` 的字符排序得到 `key`；
   - 将 `s` 追加到 `groups[key]` 中。
3. 把哈希表中的每组收集到结果列表，返回。

### 复杂度分析

- **时间复杂度**: O(n · k log k)，n 为字符串个数，k 为字符串平均长度；每个字符串排序一次
- **空间复杂度**: O(n · k)，哈希表存储所有字符串及其 key

## 代码实现

```go
func GroupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string, len(strs))
    for _, s := range strs {
        key := sortedString(s)            // 排序后的字符串作为分组 key
        groups[key] = append(groups[key], s)
    }
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    return result
}

// sortedString 返回字符串排序后的副本，作为异位词的统一 key
func sortedString(s string) string {
    bytes := []byte(s)
    sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
    return string(bytes)
}
```

**执行过程示例**（`strs = ["eat","tea","tan","ate","nat","bat"]`）：

```
"eat" → key "aet" → groups: {aet:[eat]}
"tea" → key "aet" → groups: {aet:[eat, tea]}
"tan" → key "ant" → groups: {aet:[eat, tea], ant:[tan]}
"ate" → key "aet" → groups: {aet:[eat, tea, ate], ant:[tan]}
"nat" → key "ant" → groups: {aet:[eat, tea, ate], ant:[tan, nat]}
"bat" → key "abt" → groups: {..., abt:[bat]}
收集结果: [["eat","tea","ate"], ["tan","nat"], ["bat"]]
```
