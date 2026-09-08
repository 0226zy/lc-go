# 35. 搜索插入位置 (Search Insert Position)

## 题目描述

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 `O(log n)` 的算法。

### 示例 1

```
输入: nums = [1,3,5,6], target = 5
输出: 2
```

### 示例 2

```
输入: nums = [1,3,5,6], target = 2
输出: 1
```

### 示例 3

```
输入: nums = [1,3,5,6], target = 7
输出: 4
```

## 提示

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 为**无重复元素**的**升序**排列数组
- `-10^4 <= target <= 10^4`

## 题目解析

### 核心思路

这是 **lower_bound 二分模板**的裸题：在升序数组中找到第一个 `>= target` 的位置。

- 如果数组里存在 `target`，第一个 `>= target` 的位置恰好就是 `target` 的下标。
- 如果不存在，第一个 `>= target` 的位置恰好就是它应该插入的位置（前面的元素都比它小，后面的都比它大）。

关键是不停地收缩答案区间 `[left, right]`：每次取中点 `mid`，若 `nums[mid] >= target` 说明答案在 `[left, mid]`（含 mid），否则答案在 `[mid+1, right]`。循环结束时 `left == right`，即为答案。

> 记忆口诀：求「第一个满足某条件的位置」用 `right = mid`，区间端点始终保持「答案可能在其中」，退出时 `left` 就是答案。

### 算法步骤

1. 初始化 `left, right = 0, len(nums)`（注意 right 取 len，答案可能等于 n，即插到末尾）
2. 当 `left < right`：
   - `mid = (left + right) / 2`
   - 若 `nums[mid] >= target`，`right = mid`（mid 可能是答案）
   - 否则 `left = mid + 1`（mid 及左侧都不可能是答案）
3. 返回 `left`

### 复杂度分析

- **时间复杂度**: O(log n)，二分查找
- **空间复杂度**: O(1)，常数变量

## 代码实现

```go
func SearchInsert(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] >= target {
            right = mid // mid 可能是答案，保留
        } else {
            left = mid + 1 // mid 太小，答案在右侧
        }
    }
    return left
}
```

**执行过程示例**（`nums = [1,3,5,6], target = 2`）：

```
left=0, right=4
mid=2: nums[2]=5 >= 2, right=2
mid=1: nums[1]=3 >= 2, right=1
mid=0: nums[0]=1 <  2, left=1
left == right == 1, 返回 1
```
