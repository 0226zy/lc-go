# 0167. 两数之和 II - 输入有序数组 (Two Sum II - Input Array Is Sorted)

## 题目描述

给你一个下标从 **1** 开始的整数数组 `numbers`，该数组已按 **非递减顺序** 排列。请你从数组中找出满足相加之和等于目标数 `target` 的两个数。如果设这两个数分别是 `numbers[index1]` 和 `numbers[index2]`，则 `1 <= index1 < index2 <= numbers.length`。

以长度为 2 的整数数组 `[index1, index2]` 的形式返回这两个整数的下标 `index1` 和 `index2`。

你可以假设每个输入 **只对应唯一的答案**，而且你 **不可以** 重复使用相同的元素。

你所设计的解决方案必须只使用常量级的额外空间。

### 示例

```
输入: numbers = [2,7,11,15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9。因此 index1 = 1, index2 = 2。返回 [1, 2]。
```

```
输入: numbers = [2,3,4], target = 6
输出: [1,3]
解释: 2 与 4 之和等于目标数 6。因此 index1 = 1, index2 = 3。返回 [1, 3]。
```

```
输入: numbers = [-1,0], target = -1
输出: [1,2]
解释: -1 与 0 之和等于目标数 -1。因此 index1 = 1, index2 = 2。返回 [1, 2]。
```

### 提示

- `2 <= numbers.length <= 3 * 10⁴`
- `-1000 <= numbers[i] <= 1000`
- `numbers` 按 **非递减顺序** 排列
- `-1000 <= target <= 1000`
- 仅存在一个有效答案

## 核心思路

数组已有序，适合用 **对撞双指针**：

- 左指针 `left` 指向最小端，右指针 `right` 指向最大端；
- 若 `numbers[left] + numbers[right] == target`，找到答案；
- 若和偏小，说明需要更大的加数，`left++`；
- 若和偏大，说明需要更小的加数，`right--`。

因为有序性，每次移动都能严格缩小搜索区间，且不会漏掉唯一解。

## 算法步骤

1. 初始化 `left = 0`，`right = n - 1`。
2. 当 `left < right` 时循环：
   - 计算 `sum = numbers[left] + numbers[right]`；
   - `sum == target`：返回 `[left+1, right+1]`（题目下标从 1 开始）；
   - `sum < target`：`left++`；
   - `sum > target`：`right--`。
3. 题目保证有解，正常不会走到循环结束。

## 复杂度分析

- **时间复杂度**: O(n)，两指针最多各走一遍数组。
- **空间复杂度**: O(1)，仅使用常量额外空间。
