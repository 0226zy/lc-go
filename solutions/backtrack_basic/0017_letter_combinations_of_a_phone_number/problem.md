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

这是回溯入门中最经典的 **“多叉树全枚举”** 问题：把每个数字看作一层，该数字映射的每个字母都是这一层的一个分支，穷举所有“每位数字各选一个字母”的方案，就是 DFS 走完整棵多叉树、收集所有根到叶子的路径。

回溯三要素：

- **路径**：`path[i]` 记录第 `i` 位数字当前选中的字母（用固定长度的 `[]byte` 原位写入）。
- **选择列表**：第 `i` 位数字映射的字母串，例如 `digits[i] == '7'` 时选择列表为 `"pqrs"`。
- **结束条件**：`index == len(digits)`，即每一位数字都已经选好字母，把 `path` 转成字符串收集进结果。

由于每一位数字之间没有“是否已用过”的冲突（不同位独立选择），本题 **不需要 visited 标记，也不需要显式撤销选择**：`path[index]` 在下一次循环迭代中会被直接覆盖。

### 算法步骤

1. 特判：`digits` 为空时直接返回空结果。
2. 建立映射表 `phoneMap`：下标 2~9 对应电话按键上的字母串。
3. 从 `index = 0` 开始回溯：
   - 若 `index == len(digits)`，将 `path` 转为字符串加入结果并返回；
   - 否则取出 `digits[index]` 对应的字母串，逐个字母尝试：
     - 把该字母写入 `path[index]`；
     - 递归处理下一位数字 `index + 1`。
4. 回溯结束后返回结果。

### 复杂度分析

- **时间复杂度**: O(4^n × n)，n 为 `digits` 长度。每位数字最多 4 个字母（7、9），叶子路径共 O(4^n) 条，每条路径生成一个长度为 n 的字符串需要 O(n)。
- **空间复杂度**: O(n)，递归栈最大深度为 n（不计输出结果本身占用的空间）。

## 代码实现

```go
// 数字到字母的映射表，索引即数字，0 和 1 不对应任何字母
var phoneMap = [10]string{
	"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
}

func LetterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}
	path := make([]byte, len(digits))
	var backtrack func(index int)
	backtrack = func(index int) {
		// 结束条件：所有数字都选完了
		if index == len(digits) {
			result = append(result, string(path))
			return
		}
		letters := phoneMap[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			path[index] = letters[i] // 做选择
			backtrack(index + 1)
			// path[index] 会被下一轮循环覆盖，无需显式撤销
		}
	}
	backtrack(0)
	return result
}
```

**执行过程示例**（`digits = "23"`）：

```
index=0，数字 '2' 映射 "abc"，逐层展开：
  path[0]='a' → index=1，数字 '3' 映射 "def"：
    path[1]='d' → index=2 == len(2) → 收集 "ad"
    path[1]='e' → 收集 "ae"
    path[1]='f' → 收集 "af"
  path[0]='b' → index=1：
    依次收集 "bd"、"be"、"bf"
  path[0]='c' → index=1：
    依次收集 "cd"、"ce"、"cf"
最终结果: ["ad","ae","af","bd","be","bf","cd","ce","cf"]（共 3 × 3 = 9 条路径）
```
