# 46. 全排列 (Permutations)

## 题目描述

给定一个不含重复数字的数组 `nums`，返回其所有可能的全排列。你可以按 **任何顺序** 返回答案。

### 示例 1

```
输入: nums = [1,2,3]
输出: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

### 示例 2

```
输入: nums = [0,1]
输出: [[0,1],[1,0]]
```

### 示例 3

```
输入: nums = [1]
输出: [[1]]
```

## 提示

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- `nums` 中的所有整数 **互不相同**

## 题目解析

### 核心思路

这是 **回溯“排列模板”** 的标准应用。与组合不同，排列**关心顺序**，`[1,2]` 和 `[2,1]` 是两个不同的答案，所以不能用 `start` 下标限制选择范围。

那如何防止同一个数字在一个排列里被重复使用？答案是 **`used` 数组**：记录每个位置上的数字是否已经进入当前路径。

回溯三要素：

- **路径**：已选中的数字（`path`）。
- **选择列表**：所有**还没被用过**的数字（`used[i] == false`）。
- **结束条件**：`path` 长度等于 `n`，收集答案。

### 算法步骤

1. 初始化 `used` 数组（全 false）、空路径 `path`。
2. 开始回溯：
   - 若 `len(path) == n`，把 `path` 的拷贝加入结果，返回。
   - 否则遍历数组下标 `i`：
     - 若 `used[i]` 为 true，跳过（已在路径中）；
     - 标记 `used[i] = true`，把 `nums[i]` 加入 `path`；
     - 递归进入下一层；
     - 回溯：`used[i] = false`，从 `path` 弹出 `nums[i]`。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(n × n!)，共 n! 个排列，生成每个排列需要 O(n) 的拷贝。
- **空间复杂度**: O(n)，递归栈深度与 `used` 数组均为 O(n)（不计输出）。

## 代码实现

```go
func Permute(nums []int) [][]int {
    n := len(nums)
    var result [][]int
    used := make([]bool, n)       // used[i] 表示 nums[i] 是否已在当前路径中
    path := make([]int, 0, n)

    var backtrack func()
    backtrack = func() {
        if len(path) == n {
            p := make([]int, n)
            copy(p, path)
            result = append(result, p)
            return
        }
        for i := 0; i < n; i++ {
            if used[i] {
                continue // 同一排列中每个数只能用一次
            }
            used[i] = true
            path = append(path, nums[i])
            backtrack()
            path = path[:len(path)-1]
            used[i] = false
        }
    }
    backtrack()
    return result
}
```

**执行过程示例**（`nums = [1,2,3]`）：

```
第1层: 三个都没用过，依次尝试 1, 2, 3
  path=[1]: 第2层可用 {2,3}
    path=[1,2]: 第3层只剩 {3} → 收集 [1,2,3]
    path=[1,3]: 收集 [1,3,2]
  path=[2]: 可用 {1,3} → 收集 [2,1,3]、[2,3,1]
  path=[3]: 可用 {1,2} → 收集 [3,1,2]、[3,2,1]
```
