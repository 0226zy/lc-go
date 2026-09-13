# 39. 组合总和 (Combination Sum)

## 题目描述

给你一个 **无重复元素** 的整数数组 `candidates` 和一个目标整数 `target`，找出 `candidates` 中可以使数字和为目标数 `target` 的所有 **不同组合**，并以列表形式返回。你可以按 **任意顺序** 返回这些组合。

`candidates` 中的 **同一个** 数字可以 **无限制重复被选取**。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 `target` 的不同组合少于 `150` 个。

### 示例 1

```
输入: candidates = [2,3,6,7], target = 7
输出: [[2,2,3],[7]]
解释:
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
```

### 示例 2

```
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
```

### 示例 3

```
输入: candidates = [2], target = 1
输出: []
```

## 提示

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- `candidates` 的所有元素 **互不相同**
- `1 <= target <= 40`

## 题目解析

### 核心思路

本题是回溯算法中「组合搜索」的经典变形：**候选元素可以重复使用**。整体框架仍是标准的回溯搜索树，但有两个关键点需要处理：

1. **允许重复选取**：在递归下一层时，起始下标传 `i` 而不是 `i + 1`，这样当前元素在后续层仍然可选，从而生成 `[2,2,3]` 这类含重复元素的组合。
2. **避免组合重复**：`[2,2,3]` 和 `[3,2,2]` 本质上是同一种组合。通过引入 `start` 下标，规定每一层只能从 `start` 开始向后选（下标只增不减），保证每个组合只按一种顺序被生成，天然去重。

回溯三要素：

- **路径（path）**：当前已经选入的数以及它们的累计和 `sum`。
- **选择列表**：`candidates[start:]` 中的所有元素（含当前位置，因为可重复选）。
- **结束条件**：当 `sum == target` 时，把当前路径的副本收集进答案；当 `sum > target` 时直接剪枝返回。

**剪枝优化**：先将 `candidates` 升序排序。在枚举选择时，一旦 `sum + candidates[i] > target`，由于后面的元素更大，不可能再凑出 target，直接 `break`，大幅缩小搜索空间。

### 算法步骤

1. 复制一份 `candidates` 并升序排序（不修改调用方的原始切片）。
2. 定义递归函数 `backtrack(start, sum)`：
   - 若 `sum == target`：拷贝 `path` 加入结果集，返回。
   - 从 `start` 开始遍历下标 `i`：
     - 若 `sum + candidates[i] > target`，`break` 提前剪枝；
     - 将 `candidates[i]` 追加到 `path`；
     - 递归调用 `backtrack(i, sum+candidates[i])`（传 `i`，表示当前元素可再选）；
     - 撤销选择：弹出 `path` 末尾元素。
3. 从 `backtrack(0, 0)` 启动搜索，返回结果集。

### 复杂度分析

- **时间复杂度**：最坏为 O(2^t) 量级（t 为 target / min(candidates) 量级的递归深度，每层有选与不选的分支）；排序 + 剪枝能显著减少实际搜索节点数。
- **空间复杂度**：O(t)，主要为递归栈与路径占用的空间（不计输出结果）。

## 代码实现

```go
package combinationsum

import "sort"

// CombinationSum 组合总和
// 给定一个无重复元素的正整数数组 candidates 和一个目标数 target，
// 找出所有数字和为 target 的不同组合，candidates 中的元素可以被无限次重复选取。
// 时间复杂度: O(2^t) t 为 target 量级，最坏为指数级  空间复杂度: O(t) 递归栈深度（不计输出）
func CombinationSum(candidates []int, target int) [][]int {
	// 拷贝后排序：利用单调性剪枝，同时避免修改调用方的原始切片
	sorted := make([]int, len(candidates))
	copy(sorted, candidates)
	sort.Ints(sorted)

	var ans [][]int
	path := make([]int, 0, len(sorted))

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			// 收集答案时必须拷贝，path 后续还会被修改
			combo := make([]int, len(path))
			copy(combo, path)
			ans = append(ans, combo)
			return
		}
		for i := start; i < len(sorted); i++ {
			if sum+sorted[i] > target {
				break // 已排序，后续元素更大，全部剪掉
			}
			path = append(path, sorted[i])
			backtrack(i, sum+sorted[i]) // 传 i：当前元素允许重复选取
			path = path[:len(path)-1]   // 撤销选择
		}
	}
	backtrack(0, 0)
	return ans
}
```

**执行过程示例**（`candidates = [2,3,6,7], target = 7`，排序后不变）：

```
backtrack(0, 0)  path=[]
├─ 选 2 → backtrack(0, 2)  path=[2]
│  ├─ 选 2 → backtrack(0, 4)  path=[2,2]
│  │  ├─ 选 2 → backtrack(0, 6)  path=[2,2,2]
│  │  │  └─ 2、3、6、7 全部使 sum 超界，break
│  │  ├─ 选 3 → backtrack(1, 7)  path=[2,2,3] → sum==target，收集 [2,2,3]
│  │  └─ 选 6 → 4+6>7，break
│  ├─ 选 3 → backtrack(1, 5)  path=[2,3]
│  │  └─ 选 3 → 5+3>7，break
│  └─ 选 6 → 2+6>7，break
├─ 选 3 → backtrack(1, 3)  path=[3]
│  ├─ 选 3 → backtrack(1, 6)  path=[3,3]
│  │  └─ 选 3 → 6+3>7，break
│  └─ 选 6 → 3+6>7，break
├─ 选 6 → backtrack(2, 6)  path=[6]
│  └─ 选 6 → 6+6>7，break
└─ 选 7 → backtrack(3, 7)  path=[7] → sum==target，收集 [7]
结果: [[2,2,3],[7]]
```
