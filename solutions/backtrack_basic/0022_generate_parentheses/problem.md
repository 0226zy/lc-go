# 22. 括号生成 (Generate Parentheses)

## 题目描述

数字 `n` 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 **有效的** 括号组合。

有效括号组合需满足：任意前缀中左括号 `(` 的数量不少于右括号 `)` 的数量，且左、右括号总数都恰好为 `n`。

### 示例 1

```
输入: n = 3
输出: ["((()))","(()())","(())()","()(())","()()()"]
```

### 示例 2

```
输入: n = 1
输出: ["()"]
```

## 提示

- `1 <= n <= 8`

## 题目解析

### 核心思路

这是一道经典的 **回溯 + 剪枝** 题。如果先暴力枚举 `2n` 个位置的所有括号排列（共 2^(2n) 种），再逐个验证合法性，搜索量太大且绝大多数是非法串。更好的做法是：**在生成过程中就保证每一步的前缀都是合法的**，这样回溯搜索树上根本不会长出非法分支。

维护两个计数器：`open` 表示当前已使用的左括号数，`close` 表示已使用的右括号数。每一步只有两种选择，且各有约束：

- **放左括号 `(`**：只要 `open < n` 就可以放。左括号只会增加“余额”，永远不会让前缀变非法。
- **放右括号 `)`**：只有当 `close < open` 时才能放。右括号一旦追平甚至超过左括号，前缀就不合法了，这条分支直接剪掉。

回溯三要素：

- **路径**：当前正在构造的括号串 `path`（隐含 `open`、`close` 两个计数）。
- **选择列表**：满足上述约束条件的 `(` 或 `)`。
- **结束条件**：`path` 长度达到 `2n`，此时 `open == close == n`，必为合法组合，收集之。

最终答案的个数恰好是第 n 个 **卡特兰数** Cₙ（如 n=3 时为 5，n=8 时为 1430）。

### 算法步骤

1. 初始化空路径 `path`，从 `open = 0, close = 0` 开始回溯。
2. 若 `len(path) == 2*n`，将 `path` 转为字符串加入结果集，返回。
3. 若 `open < n`：把 `(` 追加到 `path`，递归 `backtrack(open+1, close)`，回溯时弹出末尾字符。
4. 若 `close < open`：把 `)` 追加到 `path`，递归 `backtrack(open, close+1)`，回溯时弹出末尾字符。
5. 返回结果集。

### 复杂度分析

- **时间复杂度**：O(4ⁿ/√n)。合法组合数为卡特兰数 Cₙ，其渐近量级为 4ⁿ/(n√(πn))；每个组合需要 O(n) 时间拷贝字符串。
- **空间复杂度**：O(n)。递归栈深度最大为 2n（不计输出占用的空间）。

## 代码实现

```go
package generateparentheses

// GenerateParenthesis 括号生成
// 数字 n 代表生成括号的对数，返回所有可能的并且有效的括号组合。
// 时间复杂度: O(4^n / √n) 即卡特兰数 Cn 量级  空间复杂度: O(n) 递归栈深度
func GenerateParenthesis(n int) []string {
	result := []string{}
	path := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		// 路径长度达到 2n，收集一个合法组合
		if len(path) == 2*n {
			result = append(result, string(path))
			return
		}
		// 左括号还有剩余，添加左括号永远不会破坏前缀合法性
		if open < n {
			path = append(path, '(')
			backtrack(open+1, close)
			path = path[:len(path)-1]
		}
		// 右括号数量不能超过左括号，否则前缀非法，直接剪枝
		if close < open {
			path = append(path, ')')
			backtrack(open, close+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return result
}
```

**执行过程示例**（`n = 2`）：

```
backtrack(0,0) path=""
  只能放 '('（close<open 不成立）
backtrack(1,0) path="("
  放 '(' → backtrack(2,0) path="(("
    open==n 不能再放 '('，只能放 ')'
    backtrack(2,1) path="(()"
      只能放 ')' → backtrack(2,2) path="(())" 长度=4，收集 "(())"
  放 ')' → backtrack(1,1) path="()"
    放 '(' → backtrack(2,1) path="()("
      只能放 ')' → backtrack(2,2) path="()()" 长度=4，收集 "()()"
    close==open 不能再放 ')'
结果: ["(())", "()()"]
```
