# 1150. 检查一个数是否在数组中占绝大多数

> 难度：简单 ｜ 分类：二分查找 ｜ 尊享面试 100 题 · 第 77 题
> 链接：https://leetcode.cn/problems/check-if-a-number-is-majority-element-in-a-sorted-array/

## 题目描述

给你一个按**非递减顺序**排序的整数数组 `nums` ，以及一个整数 `target` 。如果 `target` 是 `nums` 中的**多数元素**（即 `target` 在 `nums` 中出现的次数**严格大于** `nums.length / 2`），返回 `true` ，否则返回 `false` 。

### 示例 1

```
输入: nums = [2,4,5,5,5,5,5,6,6], target = 5
输出: true
解释: 数字 5 出现了 5 次，数组长度为 9，5 > 9 / 2 = 4，所以 5 是多数元素。
```

### 示例 2

```
输入: nums = [10,100,101,101], target = 101
输出: false
解释: 数字 101 出现了 2 次，数组长度为 4，2 并不大于 4 / 2 = 2，所以 101 不是多数元素。
```

### 提示

- `1 <= nums.length <= 1000`
- `1 <= nums[i], target <= 10^9`
- `nums` 按非递减顺序排序

## 思路解析

### 核心思路

数组已排序，相同的元素必然连续排列成一段。因此只需用**二分查找**找到 `target` 第一次出现的下标 `first`：

- 若 `nums[first] != target`（`target` 根本不存在），直接返回 `false`；
- 否则，`target` 是多数元素当且仅当 `first + n/2` 位置上的元素仍然是 `target`（出现次数 > n/2 等价于从 `first` 开始数第 `n/2 + 1` 个位置仍是 `target`）。

也可以用两次二分分别定位 `target` 区间的左右端点再比较长度，但利用「排序 + 多数元素必占中点」的性质，一次二分即可。

### 算法步骤

1. 用二分查找（`sort.SearchInts`）找到 `target` 在 `nums` 中第一次出现的位置 `first`。
2. 若 `first == n` 或 `nums[first] != target`，返回 `false`。
3. 计算 `next = first + n/2`，若 `next < n` 且 `nums[next] == target`，返回 `true`，否则返回 `false`。

## 复杂度分析

- **时间复杂度**: O(log n)，只进行了一次二分查找。
- **空间复杂度**: O(1)，只使用常数个变量。
