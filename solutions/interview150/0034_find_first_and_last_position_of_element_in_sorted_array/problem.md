# 34. 在排序数组中查找元素的第一个和最后一个位置 (Find First and Last Position of Element in Sorted Array)

## 题目描述

给你一个按照非递减顺序排列的整数数组 `nums`，和一个目标值 `target`。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 `target`，返回 `[-1, -1]`。

你必须设计并实现一个时间复杂度为 `O(log n)` 的算法解决此问题。

### 示例 1

```
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
```

### 示例 2

```
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
```

### 示例 3

```
输入: nums = [], target = 0
输出: [-1,-1]
```

### 提示

- `0 <= nums.length <= 10^5`
- `-10^9 <= nums[i] <= 10^9`
- `nums` 是一个非递减数组
- `-10^9 <= target <= 10^9`

## 题目解析

### 核心思路

数组有序，且要求 `O(log n)`，显然要用 **二分查找**。但普通的二分找到 `target` 就停，无法保证是“第一个”或“最后一个”——因为数组里有重复元素。

这类问题有一个经典模板：**lower_bound**（下界），即在有序数组中找到“第一个 `>= target` 的元素下标”。有了它，答案就出来了：

1. **第一个位置** = `lower_bound(nums, target)`。如果该位置的值不是 `target`，说明 `target` 不存在，返回 `[-1, -1]`。
2. **最后一个位置** = `lower_bound(nums, target + 1) - 1`。即找“第一个 `> target` 的元素”，它的前一个就是最后一个 `target`。

lower_bound 的写法是二分模板里最容易记错的一种，要点是：**循环条件是 `left < right`（不是 `<=`），收缩时保留 `mid` 这一侧**。这样循环结束时 `left == right`，正好落在第一个满足条件的下标上。

本题属于 **“有序数组 + 二分边界模板”** 模型，同模板还可以解决“插入位置（35 题）”“第一个错误版本”等题。

### 算法步骤

1. 实现 `lowerBound(nums, target)`：返回第一个 `>= target` 的下标（可能等于 `len(nums)`）。
   - `left = 0, right = len(nums)`
   - 当 `left < right`：`mid = (left+right)/2`
     - 若 `nums[mid] >= target`，说明答案在 `[left, mid]`，`right = mid`；
     - 否则答案在 `(mid, right)`，`left = mid + 1`。
2. `first = lowerBound(nums, target)`；若 `first == len(nums)` 或 `nums[first] != target`，返回 `[-1, -1]`。
3. `last = lowerBound(nums, target+1) - 1`，返回 `[first, last]`。

### 复杂度分析

- **时间复杂度**: O(log n)，两次二分查找
- **空间复杂度**: O(1)，只使用常数额外空间

## 代码实现

```go
func lowerBound(nums []int, target int) int {
    left, right := 0, len(nums) // 注意：right 开区间，循环条件 left < right
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid // mid 可能是答案，保留
        } else {
            left = mid + 1
        }
    }
    return left
}

func SearchRange(nums []int, target int) []int {
    first := lowerBound(nums, target)
    if first == len(nums) || nums[first] != target {
        return []int{-1, -1}
    }
    last := lowerBound(nums, target+1) - 1
    return []int{first, last}
}
```

**执行过程示例**（`nums = [5,7,7,8,8,10]`, `target = 8`）：

```
lowerBound(8): 最终 first = 3（第一个 >= 8 的元素是 nums[3]=8）
nums[3] == 8，目标存在
lowerBound(9): 最终 left = 5（第一个 >= 9 的元素是 nums[5]=10）
last = 5 - 1 = 4
返回 [3, 4]
```
