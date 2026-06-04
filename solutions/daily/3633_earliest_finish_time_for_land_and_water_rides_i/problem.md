# 3633. 陆地和水上项目最早完成时间 I

## 题目描述

你正在参观一个主题公园，公园中有两类游乐项目：**陆地项目**和**水上项目**。

对于每类项目，你都知道它们的开放时间和持续时间：

**陆地项目：**
- `landStartTime[i]` — 第 `i` 个陆地项目最早可以登上的时间
- `landDuration[i]` — 第 `i` 个陆地项目的持续时间

**水上项目：**
- `waterStartTime[j]` — 第 `j` 个水上项目最早可以登上的时间
- `waterDuration[j]` — 第 `j` 个水上项目的持续时间

作为游客，你必须从**每类**项目中**恰好体验一个**。你可以按**任意顺序**进行（先陆地后水上，或者先水上后陆地）。

游玩规则如下：
- 你可以在项目开放时间或之后的任意时刻开始
- 如果在时间 `t` 开始一个项目，则会在时间 `t + duration` 结束
- 结束一个项目后，你可以立即登上下一个项目（如果已经开放），或者等待它开放

返回你完成两个项目的**最早可能时间**。

### 示例 1

```
输入: landStartTime = [2,8], landDuration = [4,1], waterStartTime = [6], waterDuration = [3]
输出: 9
```

### 示例 2

```
输入: landStartTime = [5], landDuration = [3], waterStartTime = [1], waterDuration = [10]
输出: 14
```

### 约束条件

- `1 <= landStartTime.length == landDuration.length <= 100`
- `1 <= waterStartTime.length == waterDuration.length <= 100`
- `1 <= landStartTime[i], landDuration[i], waterStartTime[j], waterDuration[j] <= 1000`

## 解题思路

### 贪心 + 枚举

由于只需要从每类项目中各选一个，且顺序有两种（陆地→水上 或 水上→陆地），我们可以分别计算两种顺序的最早完成时间，然后取最小值。

对于一个固定的顺序（如先第一类，后第二类）：
1. **第一类的最早结束时间**：对于第一类的每个项目，最早结束时间 = 开始时间 + 持续时间。我们希望第一类结束得越早越好，因此取所有项目中最早结束时间的**最小值**：
   ```
   minEnd = min(start[i] + duration[i])
   ```
2. **第二类的最早完成时间**：对于第二类的每个项目 `j`，我们必须等第一类结束后才能开始，同时也不能早于该项目的开放时间。因此实际开始时间为 `max(start[j], minEnd)`，完成时间为 `max(start[j], minEnd) + duration[j]`。取所有第二类项目完成时间的**最小值**。

### 算法步骤

1. 定义辅助函数 `calc(s1, d1, s2, d2)`，表示先玩第一类、再玩第二类的最早完成时间。
2. 计算 `minEnd = min(s1[i] + d1[i])`。
3. 计算 `minFinish = min(max(s2[j], minEnd) + d2[j])`。
4. 返回两种顺序的最小值：`min(calc(land, water), calc(water, land))`。

## 复杂度分析

- **时间复杂度**: `O(n + m)`，其中 `n` 为陆地项目数量，`m` 为水上项目数量。辅助函数分别遍历两类项目各一次，调用两次。
- **空间复杂度**: `O(1)`，只使用了若干变量。
