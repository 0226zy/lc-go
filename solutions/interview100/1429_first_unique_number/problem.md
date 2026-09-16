# 1429. 第一个唯一数字

> 难度：中等 ｜ 分类：队列 ｜ 尊享面试 100 题 · 第 33 题
> 链接：https://leetcode.cn/problems/first-unique-number/

## 题目描述

给定一系列整数，请你设计一个数据结构，既可以高效地返回序列中的**第一个唯一数字**，也可以向序列中追加数字。

实现 `FirstUnique` 类：

- `FirstUnique(int[] nums)` 用整数数组 `nums` 初始化该对象；
- `int showFirstUnique()` 返回序列中的第一个唯一数字。如果不存在唯一数字，返回 `-1`；
- `void add(int value)` 将 `value` 追加到序列中。

**唯一数字**指在序列中只出现一次的数字；「第一个」按加入序列的先后顺序判断。

### 示例 1

```
输入:
["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique"]
[[[2,3,5]],[],[5],[],[2],[],[3],[]]
输出:
[null,2,null,2,null,3,null,-1]
解释:
FirstUnique firstUnique = new FirstUnique([2,3,5]);
firstUnique.showFirstUnique(); // 返回 2
firstUnique.add(5);            // 此时序列为 [2,3,5,5]
firstUnique.showFirstUnique(); // 返回 2
firstUnique.add(2);            // 此时序列为 [2,3,5,5,2]
firstUnique.showFirstUnique(); // 返回 3
firstUnique.add(3);            // 此时序列为 [2,3,5,5,2,3]
firstUnique.showFirstUnique(); // 返回 -1，所有数字都出现了两次
```

### 示例 2

```
输入:
["FirstUnique","showFirstUnique","add","add","add","add","showFirstUnique"]
[[[7,7,7,7,7,7]],[],[7],[3],[3],[7],[17]]
输出:
[null,-1,null,null,null,null,17]
解释:
FirstUnique firstUnique = new FirstUnique([7,7,7,7,7,7]);
firstUnique.showFirstUnique(); // 返回 -1
firstUnique.add(7);            // 此时序列为 [7,7,7,7,7,7,7]
firstUnique.add(3);            // 此时序列为 [7,7,7,7,7,7,7,3]
firstUnique.add(3);            // 此时序列为 [7,7,7,7,7,7,7,3,3]
firstUnique.add(7);            // 此时序列为 [7,7,7,7,7,7,7,3,3,7]
firstUnique.add(17);           // 此时序列为 [7,7,7,7,7,7,7,3,3,7,17]
firstUnique.showFirstUnique(); // 返回 17
```

### 提示

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^8`
- `1 <= value <= 10^8`
- `showFirstUnique` 和 `add` 的调用总数最多为 `50000` 次

## 思路解析

### 核心思路

「第一个唯一数字」需要同时维护**出现次数**和**加入顺序**，用「**哈希表计数 + 队列维护候选顺序**」的经典组合：

- 哈希表 `count` 记录每个数字出现的次数；
- 队列 `queue` 保存「目前仍唯一」的候选数字，按加入顺序排列；
- `add(v)` 时：若 `v` 是首次出现，加入队列；`count[v]` 加一；
- `showFirstUnique()` 时：**惰性删除**——只要队首数字的次数已大于 1 就弹出，直到队首唯一或队列为空。

每个数字最多进队一次、出队一次，因此每个操作是**均摊 O(1)** 的，避免了每次查询都从头扫描。

### 算法步骤

1. `Constructor`：初始化哈希表与空队列，依次对每个 `nums[i]` 调用 `Add`。
2. `Add(v)`：若 `count[v] == 0` 则把 `v` 追加到队尾；然后 `count[v]++`。
3. `ShowFirstUnique()`：
   - 循环弹出所有 `count > 1` 的队首元素；
   - 若队列为空返回 `-1`，否则返回队首元素。

## 复杂度分析

- **时间复杂度**: `Add` 为 O(1)，`ShowFirstUnique` 为均摊 O(1)（每个数字最多被弹出一次，均摊到所有操作上）。
- **空间复杂度**: O(n)，哈希表和队列最多存下所有不同的数字。
