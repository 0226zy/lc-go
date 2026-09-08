# 17. 电话号码的字母组合 (Letter Combinations of a Phone Number)

## 题目描述

给定一个仅包含数字 `2-9` 的字符串，返回所有它能表示的字母组合。答案可以按 **任意顺序** 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

```
2: abc   3: def
4: ghi   5: jkl   6: mno
7: pqrs  8: tuv   9: wxyz
```

### 示例 1

```
输入: digits = "23"
输出: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

### 示例 2

```
输入: digits = ""
输出: []
```

### 示例 3

```
输入: digits = "2"
输出: ["a","b","c"]
```

## 提示

- `0 <= digits.length <= 4`
- `digits[i]` 是数字范围 `['2', '9']` 的一个数字

## 题目解析

### 核心思路

这道题是典型的 **回溯“多叉树枚举”模板**：每个数字对应一组可选字母，我们要枚举所有“为每个数字选一个字母”的组合方案。

想象一棵多叉树：第 0 层为第 1 个数字选字母（3 或 4 个分支），第 1 层为第 2 个数字选字母……当走到最后一层时，就得到一个完整的组合。深度优先遍历（DFS）这棵树，把每一条从根到叶子的路径收集起来，就是全部答案。

回溯三要素在这道题里体现得很干净：

- **路径**：`path[i]` 表示第 `i` 位数字已经选好的字母。
- **选择列表**：当前数字对应的字母串（如 `'2'` 对应 `"abc"`）。
- **结束条件**：所有数字都选完了（`index == len(digits)`），把路径加入结果。

### 算法步骤

1. 如果 `digits` 为空，直接返回空结果。
2. 定义映射表 `phoneMap`：数字 → 对应字母串。
3. 从 `index = 0` 开始回溯：
   - 若 `index == len(digits)`，说明每一位都选好了，将 `path` 转为字符串加入结果，返回。
   - 否则取出 `digits[index]` 对应的字母串，逐个尝试：
     - 把字母写入 `path[index]`；
     - 递归处理下一位（`index + 1`）。
4. 所有递归结束后返回结果。

### 复杂度分析

- **时间复杂度**: O(4^n × n)，其中 n 是 digits 长度。每个数字最多 4 个字母（7、9），共 O(4^n) 条路径，每条路径复制一次字符串需要 O(n)。
- **空间复杂度**: O(n)，递归栈最大深度为 n（不计输出结果占用的空间）。

## 代码实现

```go
var phoneMap = [10]string{
    "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

func LetterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    var result []string
    path := make([]byte, len(digits)) // path[i] 表示第 i 位数字选中的字母

    var backtrack func(index int)
    backtrack = func(index int) {
        if index == len(digits) {
            result = append(result, string(path))
            return
        }
        letters := phoneMap[digits[index]-'0']
        for i := 0; i < len(letters); i++ {
            path[index] = letters[i]
            backtrack(index + 1) // 进入下一位数字
        }
    }
    backtrack(0)
    return result
}
```

**执行过程示例**（`digits = "23"`）：

```
index=0: 数字2对应"abc"，依次尝试：
  path[0]='a' → index=1: 数字3对应"def"
    path[1]='d' → index=2 == len → 收集 "ad"
    path[1]='e' → 收集 "ae"
    path[1]='f' → 收集 "af"
  path[0]='b' → 同理收集 "bd", "be", "bf"
  path[0]='c' → 同理收集 "cd", "ce", "cf"
```
