# 39. 组合总和 (Combination Sum)

## 题目描述

给你一个 **无重复元素** 的整数数组 `candidates` 和一个目标整数 `target`，找出 `candidates` 中可以使数字和为目标数 `target` 的所有 **不同组合** ，并以列表形式返回。你可以按 **任意顺序** 返回这些组合。

`candidates` 中的 **同一个** 数字可以 **无限制重复被选取** 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 `target` 的不同组合少于 `150` 个。

### 示例 1

```
输入: candidates = [2,3,6,7], target = 7
输出: [[2,2,3],[7]]
解释:
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
只有一种组合。
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

这是 **回溯“可重复选取”的组合模板**，和「组合」一题（77 题）是同一个骨架，区别只有两处：

1. **元素可以无限重复选**：递归时传入的下标是 `i` 而不是 `i + 1`——下一轮仍然可以从当前元素开始选。比如 `[2,2,...]` 就是这样产生的。
2. **依然要去重**：还是靠 `start` 保证不回退。比如 `[2,2,3]` 和 `[3,2,2]` 是同一种组合，靠“下标只增不减”保证每个组合按数组顺序生成，只出现一次。

再叠加一个强力 **剪枝**：先把 `candidates` 排序，当 `sum + candidates[i] > target` 时，后面的元素只会更大，直接 `break` 跳出循环，不用继续尝试。

回溯三要素：

- **路径**：已选中的数（`path`）与其和 `sum`。
- **选择列表**：下标 `>= start` 的候选数（可重复）。
- **结束条件**：`sum == target`，收集答案；`sum > target` 由剪枝提前拦截。

### 算法步骤

1. 将 `candidates` 升序排序。
2. 从 `start = 0, sum = 0` 开始回溯：
   - 若 `sum == target`，把 `path` 的拷贝加入结果，返回。
   - 遍历 `i` 从 `start` 到末尾：
     - 若 `sum + candidates[i] > target`，直接 `break`（排序保证后面更大）；
     - 把 `candidates[i]` 加入 `path`；
     - 递归 `backtrack(i, sum + candidates[i])`（注意传 `i`，允许重复选）；
     - 回溯：弹出 `candidates[i]`。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(2^t) 量级（t 为 target 值），最坏情况下每个元素都有“选/不选”两种决策；剪枝能显著减少实际搜索量。
- **空间复杂度**: O(t)，递归栈深度最坏为 target/min(candidates) 层（不计输出）。

## 代码实现

```go
func CombinationSum(candidates []int, target int) [][]int {
    sort.Ints(candidates) // 排序后可利用单调性提前剪枝
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
            path = append(path, candidates[i])
            backtrack(i, sum+candidates[i]) // 传 i 而非 i+1：允许重复选取当前元素
            path = path[:len(path)-1]
        }
    }
    backtrack(0, 0)
    return result
}
```

**执行过程示例**（`candidates = [2,3,6,7], target = 7`）：

```
path=[] sum=0:
  选2 → path=[2] sum=2:
    选2 → path=[2,2] sum=4:
      选2 → sum=6 → 再选2 超界剪掉，选3 → sum=7 → 收集 [2,2,3]
    选3 → path=[2,3] sum=5 → 选3 超界，选6、7 超界
    选6 → 超界，break
    选7 → 超界，break
  选3 → path=[3] sum=3:
    选3 → sum=6，后续全超界
    选6、7 超界
  选6 → sum=6，后续超界
  选7 → sum=7 → 收集 [7]
结果: [[2,2,3],[7]]
```
