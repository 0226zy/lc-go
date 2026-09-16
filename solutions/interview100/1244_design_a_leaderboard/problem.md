# 1244. 力扣排行榜

> 难度：中等 ｜ 分类：设计 ｜ 尊享面试 100 题 · 第 86 题
> 链接：https://leetcode.cn/problems/design-a-leaderboard/

## 题目描述

请你设计一个「力扣排行榜」系统，需要支持以下三种操作：

- `AddScore(playerId, score)`：把参赛者 `playerId` 的分数增加 `score` 分。如果该参赛者不在排行榜中，则把他加入排行榜，并把分数初始化为 `score`。
- `Top(K)`：返回当前排行榜上前 `K` 名参赛者的分数之和。
- `Reset(playerId)`：重置参赛者 `playerId` 的分数为 0，并把他移出排行榜。题目保证调用前该参赛者一定在排行榜中。

### 示例 1

```
输入:
["Leaderboard","addScore","addScore","addScore","addScore","addScore","top","reset","reset","addScore","top"]
[[],[1,73],[2,56],[3,39],[4,51],[5,4],[1],[1],[2],[2,51],[3]]
输出:
[null,null,null,null,null,null,73,null,null,null,141]

解释:
Leaderboard leaderboard = new Leaderboard();
leaderboard.addScore(1,73);   // 排行榜为 [[1,73]]
leaderboard.addScore(2,56);   // 排行榜为 [[1,73],[2,56]]
leaderboard.addScore(3,39);   // 排行榜为 [[1,73],[2,56],[3,39]]
leaderboard.addScore(4,51);   // 排行榜为 [[1,73],[2,56],[3,39],[4,51]]
leaderboard.addScore(5,4);    // 排行榜为 [[1,73],[2,56],[3,39],[4,51],[5,4]]
leaderboard.top(1);           // 返回 73
leaderboard.reset(1);         // 排行榜为 [[2,56],[3,39],[4,51],[5,4]]
leaderboard.reset(2);         // 排行榜为 [[3,39],[4,51],[5,4]]
leaderboard.addScore(2,51);   // 排行榜为 [[2,51],[3,39],[4,51],[5,4]]
leaderboard.top(3);           // 返回 141 = 51 + 51 + 39
```

### 示例 2

```
输入:
["Leaderboard","addScore","top","top","addScore","top"]
[[],[10,20],[1],[2],[10,30],[1]]
输出:
[null,null,20,20,null,50]

解释: 只有一个参赛者 10 号，分数先为 20，top(2) 也只会累加现有参赛者的分数；
     再加 30 分后 top(1) 返回 50。
```

### 提示

- `1 <= playerId, K <= 10000`
- 调用 `top` 时，`K` 不超过当前排行榜中的参赛者人数（示例 2 仅为说明行为，测试以保证 `K` 合法的数据为准）；实际实现中即使 `K` 超过人数也只需累加所有人的分数
- `1 <= score <= 100`
- `AddScore`、`Top`、`Reset` 的总调用次数不超过 1000 次

## 思路解析

### 核心思路

排行榜本质上维护「参赛者 → 当前总分」的映射，外加一个「求前 K 大之和」的查询。

- `AddScore` 和 `Reset` 都是针对单个参赛者的操作，用**哈希表** `scores`（playerId → 总分）即可 O(1) 完成：`AddScore` 累加，`Reset` 直接删除。
- `Top(K)` 需要前 K 大的分数之和。最直接的实现是每次把哈希表中的分数取出来**排序**，再累加前 K 个。本题总调用次数不超过 1000，参赛者不超过 10000 人，排序的代价完全可接受，且代码最简单可靠。

进一步优化的思路（供参考）：如果 `Top` 调用频繁，可以用「哈希表 + 大小为 K 的最小堆」把单次查询降到 O(n log K)，或用平衡树/有序映射维护「分数 → 人数」做到 O(log n) 查询；这里采用最简洁的排序方案。

### 算法步骤

1. `Constructor`：初始化空哈希表 `scores`。
2. `AddScore(playerId, score)`：`scores[playerId] += score`（不存在时自动从 0 开始累加，天然处理了「新加入」的情况）。
3. `Reset(playerId)`：`delete(scores, playerId)`。
4. `Top(K)`：
   - 把 `scores` 中所有分数收集到切片；
   - 按从大到小排序；
   - 累加前 `K` 个（若切片长度不足 `K` 则全部累加）并返回。

## 复杂度分析

设当前排行榜中的参赛者人数为 `n`：

- **AddScore**：时间 O(1)，空间 O(1)。
- **Reset**：时间 O(1)，空间 O(1)。
- **Top(K)**：时间 O(n log n)（排序主导），空间 O(n)（收集分数的切片）。
- 整体空间复杂度为 O(n)，哈希表最多存下所有参赛者。
