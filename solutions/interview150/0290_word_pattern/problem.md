# 290. 单词规律 (Word Pattern)

## 题目描述

给定一种规律 `pattern` 和一个字符串 `s`，判断 `s` 是否遵循相同的规律。

这里的"遵循"指完全匹配，例如 `pattern` 里的每个字母和字符串 `s` 中的每个**单词**之间存在着双向一一对应的映射规律。

### 示例 1

```
输入: pattern = "abba", s = "dog cat cat dog"
输出: true
```

### 示例 2

```
输入: pattern = "abba", s = "dog cat cat fish"
输出: false
```

### 示例 3

```
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false
```

### 提示

- `1 <= pattern.length <= 300`
- `pattern` 只包含小写英文字母
- `1 <= s.length <= 3000`
- `s` 只包含小写英文字母和空格 `' '`
- `s` **不包含** 任何前导或尾随的空格
- `s` 中每个单词都被**单个空格**分隔

## 题目解析

### 核心思路

这是 **双映射哈希** 模型的标准题，和 205「同构字符串」同宗同源：那里是两个字符序列互相同构，这里是“字符序列”和“单词序列”互相同构。

把 `pattern` 和 `s` 按空格切出来的单词列表对齐：`pattern` 的第 i 个字符对应单词列表的第 i 个单词。要求：

1. **字符 → 单词** 唯一：`pattern` 中同一个字符，对应位置上的单词必须总是一样的（`a` 出现两次，对应单词就得相同）。
2. **单词 → 字符** 唯一：同一个单词不能同时被两个不同的字符映射（`dog` 不能既是 `a` 的翻译又是 `b` 的翻译）。

两个方向都要检查，缺一不可。

### 算法步骤

1. 按空格把 `s` 切分成单词列表 `words`。
2. 若 `len(pattern) != len(words)`，长度对不上，返回 `false`。
3. 建立两张哈希表：`char2Word`（字符到单词）和 `word2Char`（单词到字符）。
4. 逐个位置对齐遍历：
   - 若字符 `c` 已有映射但对应的单词不是当前单词，返回 `false`（一对多冲突）。
   - 否则若当前单词已被其他字符占用（`word2Char` 中已有记录），返回 `false`（多对一冲突）。
   - 否则建立双向映射。
5. 全部对齐成功，返回 `true`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为 s 的长度（切分 + 一次对齐遍历）
- **空间复杂度**: O(n)，存储切分出的单词与两张映射表

## 代码实现

```go
func WordPattern(pattern string, s string) bool {
    words := strings.Split(s, " ")
    if len(pattern) != len(words) {
        return false
    }
    char2Word := make(map[byte]string, len(pattern))
    word2Char := make(map[string]byte, len(pattern))
    for i := 0; i < len(pattern); i++ {
        c, w := pattern[i], words[i]
        if mappedWord, ok := char2Word[c]; ok {
            if mappedWord != w {
                return false // 同一字符映射到了不同单词
            }
        } else if _, ok := word2Char[w]; ok {
            return false // 同一单词被不同字符映射
        } else {
            char2Word[c] = w
            word2Char[w] = c
        }
    }
    return true
}
```

**执行过程示例**（`pattern = "abba"`, `s = "dog cat cat dog"`）：

```
对齐: a↔dog, b↔cat, b↔cat, a↔dog
i=0: 建立 a→dog, dog→a
i=1: 建立 b→cat, cat→b
i=2: b 已映射 cat，与当前单词一致 ✓
i=3: a 已映射 dog，与当前单词一致 ✓
返回 true
```

反例（`pattern = "abba"`, `s = "dog cat cat fish"`）：i=3 时 `a` 已映射 `dog`，当前单词是 `fish`，一对多冲突，返回 `false`。
