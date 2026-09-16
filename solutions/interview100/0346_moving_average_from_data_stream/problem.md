# 346. 数据流中的移动平均值

> 难度：简单 ｜ 分类：队列 ｜ 尊享面试 100 题 · 第 32 题
> 链接：https://leetcode.cn/problems/moving-average-from-data-stream/

## 题目描述

给定一个整数数据流和一个滑动窗口的大小 `size`，请设计一个 `MovingAverage` 类，计算数据流中滑动窗口内所有整数的移动平均值。

实现 `MovingAverage` 类：

- `MovingAverage(int size)` 用窗口大小 `size` 初始化对象。
- `double next(int val)` 向数据流中加入一个整数 `val`，并返回当前滑动窗口（最多包含最近 `size` 个数）内所有整数的平均值。

### 示例 1

```
输入:
["MovingAverage", "next", "next", "next", "next"]
[[3], [1], [10], [3], [5]]
输出:
[null, 1.0, 5.5, 4.66667, 6.0]

解释:
MovingAverage movingAverage = new MovingAverage(3);
movingAverage.next(1);  // 返回 1.0 = 1 / 1
movingAverage.next(10); // 返回 5.5 = (1 + 10) / 2
movingAverage.next(3);  // 返回 4.66667 = (1 + 10 + 3) / 3
movingAverage.next(5);  // 返回 6.0 = (10 + 3 + 5) / 3，此时 1 滑出窗口
```

### 示例 2

```
输入:
["MovingAverage", "next", "next"]
[[1], [5], [7]]
输出:
[null, 5.0, 6.0]

解释:
窗口大小为 1，每次平均值就是当前加入的那个数。
```

### 提示

- `1 <= size <= 1000`
- `-10^5 <= val <= 10^5`
- 最多调用 `10^4` 次 `next`

## 思路解析

### 核心思路

「最近 `size` 个数」是一个先进先出的滑动窗口，用**队列**维护最自然：新值从队尾进，窗口满了就把队首（最老的值）弹出。

如果每次 `next` 都遍历队列求和，复杂度是 O(size)。注意到窗口满员时，新和 = 旧和 − 弹出值 + 新值，因此额外维护一个**滚动和** `sum`，每次操作只需 O(1)。

### 算法步骤

1. 结构体中保存：窗口容量 `size`、存放窗口元素的切片队列 `queue`、当前窗口元素之和 `sum`。
2. `next(val)`：
   - 若 `len(queue) == size`，说明窗口已满，弹出队首元素并从 `sum` 中减去；
   - 将 `val` 入队，`sum += val`；
   - 返回 `float64(sum) / float64(len(queue))`（窗口未满时按当前实际元素个数取平均）。

## 复杂度分析

- **时间复杂度**：`next` 为 O(1)，只有一次入队、至多一次出队和常数次算术运算。
- **空间复杂度**：O(size)，队列最多保存 `size` 个元素。
