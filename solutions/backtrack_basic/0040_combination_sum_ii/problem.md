# 40. 组合总和 II (Combination Sum II)

## 题目描述

给定一个候选人编号的集合 `candidates` 和一个目标数 `target`，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的每个数字在每个组合中只能使用 **一次** 。

**注意：** 解集不能包含重复的组合。

### 示例 1

```
输入: candidates = [10,1,2,7,6,1,5], target = 8
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

### 示例 2

```
输入: candidates = [2,5,2,1,2], target = 5
输出:
[
[1,2,2],
[5]
]
```

## 提示

- `1 <= candidates.length <= 100`
- `1 <= candidates[i] <= 50`
- `1 <= target <= 30`

## 题目解析

### 核心思路

本题是「39. 组合总和」的变体，但有两个关键区别，都直接影响回溯的写法：

1. **数组中有重复元素**：比如 `[10,1,2,7,6,1,5]` 里有两个 1。如果不做处理，`[1,7]` 会因为选了不同下标的 1 而生成两次，导致结果重复。
2. **每个元素只能用一次**：递归时传入的下标必须是 `i + 1` 而不是 `i`——下一轮不能再从当前元素开始选。

解决重复组合的标准套路是 **排序 + 同层去重剪枝**：

- 先排序，让值相同的元素在数组中相邻。
- 在回溯的**同一层循环**里，若 `candidates[i] == candidates[i-1]` 且 `i > start`，说明这个值在本层已经被“前一个相同元素”代表的子树完整探索过了（那个子树里允许选它自己以及之后所有元素），当前再选它只会生成其子集，必然重复，直接 `continue` 跳过。
- 注意判断条件必须是 `i > start` 而不是 `i > 0`：跨层时 `path` 里已经选了一个相同值（比如 `[1,1,6]` 里两个 1 分别来自两层），这种情况是合法的，不能跳过。

回溯三要素：

- **路径**：已选中的数（`path`）与其和 `sum`。
- **选择列表**：下标 `>= start` 的候选数（每个只用一次），同层重复值只取第一次。
- **结束条件**：`sum == target`，收集答案；`sum > target` 由排序剪枝提前拦截。

### 算法步骤

1. 将 `candidates` 升序排序。
2. 从 `start = 0, sum = 0` 开始回溯：
   - 若 `sum == target`，把 `path` 的拷贝加入结果，返回。
   - 遍历 `i` 从 `start` 到末尾：
     - 若 `sum + candidates[i] > target`，直接 `break`（排序保证后面更大）；
     - 若 `i > start` 且 `candidates[i] == candidates[i-1]`，`continue`（同层去重）；
     - 把 `candidates[i]` 加入 `path`；
     - 递归 `backtrack(i+1, sum + candidates[i])`（传 `i+1`，每个元素只用一次）；
     - 回溯：弹出 `candidates[i]`。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(2^n) 量级（n 为数组长度），每个元素对应“选/不选”的决策树；排序与同层去重剪枝能显著减少实际搜索量。
- **空间复杂度**: O(n)，递归栈深度最大为 n 层（不计输出）。

## 代码实现

```go
func CombinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates) // 排序：既能让重复元素相邻便于去重，也能利用单调性剪枝
    var result [][]int
    path := []int{}

    var backtrack func(start, sum int)
    backtrack = func(start, sum int) {
        if sum == target {
            combo := make([]int, len(path))
            copy(combo, path)
            result = append(result, combo)
            return
        }
        for i := start; i < len(candidates); i++ {
            if sum+candidates[i] > target {
                break // candidates 已排序，后面的更大，直接剪掉
            }
            if i > start && candidates[i] == candidates[i-1] {
                continue // 同层去重：同一层的相同值只尝试第一次，避免产生重复组合
            }
            path = append(path, candidates[i])
            backtrack(i+1, sum+candidates[i]) // 传 i+1：每个元素在一个组合中只能用一次
            path = path[:len(path)-1]
        }
    }
    backtrack(0, 0)
    return result
}
```

**执行过程示例**（`candidates = [2,5,2,1,2], target = 5`，排序后为 `[1,2,2,2,5]`）：

```
path=[] sum=0:
  选1 → path=[1] sum=1:
    选2(下标1) → path=[1,2] sum=3:
      选2(下标2) → path=[1,2,2] sum=5 → 收集 [1,2,2]
      选2(下标3) → 同层重复(i>start)，跳过
      选5 → 超界，break
    选2(下标2) → 同层重复，跳过
    选2(下标3) → 同层重复，跳过
    选5 → path=[1,5] sum=6 → 超界，break
  选2(下标1) → path=[2] sum=2:
    选2(下标2) → path=[2,2] sum=4 → 后续选2超界、选5超界
    选2(下标3) → 同层重复，跳过
    选5 → 超界，break
  选2(下标2)、选2(下标3) → 同层重复，均跳过
  选5 → path=[5] sum=5 → 收集 [5]
结果: [[1,2,2],[5]]
```

可以看到：如果没有同层去重，`[1,2,2]` 会因子树中不同下标的 2 被多次生成；加上 `i > start && candidates[i] == candidates[i-1]` 的跳过后，每个组合恰好出现一次。
