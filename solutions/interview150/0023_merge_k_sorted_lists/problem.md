# 23. 合并 K 个升序链表 (Merge k Sorted Lists)

## 题目描述

给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。

### 示例 1

```
输入: lists = [[1,4,5],[1,3,4],[2,6]]
输出: [1,1,2,3,4,4,5,6]
解释: 链表数组化为:
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到:
1->1->2->3->4->4->5->6
```

### 示例 2

```
输入: lists = []
输出: []
```

### 示例 3

```
输入: lists = [[]]
输出: []
```

### 提示

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` 按升序排列
- `lists[i].length` 的总和不超过 `10^4`

## 题目解析

### 核心思路

合并两个有序链表很简单（双指针），这道题的关键是**怎么把 k 个两两合并的总代价降下来**。

- **朴素想法**：挨个合并——先把第 1、2 条合并，再把结果和第 3 条合并…… 每轮都要重新遍历已经合并好的长链表，总代价是 O(k·N)，k 大时很慢。这就像把 k 个数相加，一个接一个地加。
- **分治合并**：把 k 个链表两两配对：第 1、2 条合并，第 3、4 条合并…… 一轮后剩 k/2 个链表，长度翻倍；重复 log k 轮后只剩 1 个。每个节点只参与 log k 次合并，总代价 O(N log k)。这就像锦标赛：每一轮每个选手只打一场比赛。

这正是“**分治两两合并**”模板（类似归并排序的合并阶段）。另一种经典解法是**最小堆**：把 k 个头节点放进小根堆，每次弹出最小的接到结果链上，再把它的后继推回堆中，堆中始终最多 k 个元素，复杂度同为 O(N log k)，空间 O(k)。

两种解法都给出：主解法是分治版，堆版作为对照。

### 算法步骤（分治版）

1. 递归函数 `mergeRange(lists, low, high)`：
   - 若 `low == high`，只剩一个链表，直接返回；
   - 取中点 `mid`，左半 = `mergeRange(low, mid)`，右半 = `mergeRange(mid+1, high)`；
   - 返回 `mergeTwoLists(左半, 右半)`。
2. 从 `mergeRange(0, len(lists)-1)` 开始；空数组直接返回 `nil`。

### 复杂度分析

| 解法 | 时间复杂度 | 空间复杂度 |
|------|-----------|-----------|
| 分治两两合并（主解法） | O(N log k)，N 为节点总数 | O(log k) 递归栈 |
| 最小堆 | O(N log k) | O(k) 堆容量 |
| 逐个合并（对照，不推荐） | O(k·N) | O(1) |

## 代码实现

```go
// 分治：把 k 个链表两两配对合并
func MergeKLists(lists []*datastructures.ListNode) *datastructures.ListNode {
    if len(lists) == 0 {
        return nil
    }
    return mergeRange(lists, 0, len(lists)-1)
}

func mergeRange(lists []*datastructures.ListNode, low, high int) *datastructures.ListNode {
    if low == high {
        return lists[low]
    }
    mid := low + (high-low)/2
    left := mergeRange(lists, low, mid)
    right := mergeRange(lists, mid+1, high)
    return mergeTwoLists(left, right)
}
```

**执行过程示例**（`lists = [[1,4,5],[1,3,4],[2,6]]`）：

```
mergeRange(0,2)
├─ mergeRange(0,1)
│   ├─ mergeRange(0,0) → 1->4->5
│   ├─ mergeRange(1,1) → 1->3->4
│   └─ 合并 → 1->1->3->4->4->5
├─ mergeRange(2,2) → 2->6
└─ 合并 1->1->3->4->4->5 与 2->6
   → 1->1->2->3->4->4->5->6
```

可以看到每个节点最多参与 2 次合并（log 3 向上取整），而逐个合并时最后一个链表的节点要参与 3 次合并——k 越大差距越明显。

**堆版对照**：`container/heap` 维护一个以节点值排序的小根堆，初始把 k 个非空头节点入堆；每次 `Pop` 最小节点接到结果链尾，若它有后继则把后继 `Push` 入堆，直到堆空。
