# 25. K 个一组翻转链表 (Reverse Nodes in k-Group)

## 题目描述

给你链表的头节点 `head` ，每 `k` 个节点一组进行翻转，请你返回修改后的链表。

`k` 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 `k` 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

### 示例 1

```
输入: head = [1,2,3,4,5], k = 2
输出: [2,1,4,3,5]
```

### 示例 2

```
输入: head = [1,2,3,4,5], k = 3
输出: [3,2,1,4,5]
```

## 提示

- 链表中的节点数目为 `n`
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

## 题目解析

### 核心思路

这是 92 题「反转链表 II」的进阶版：92 题只反转固定一段，本题要**反复反转长度为 k 的一段，直到链表末尾**。

把问题拆成两层就清楚了：

1. **外层**：逐组推进。每组反转前，先确认剩余节点够不够 k 个——不够就整体保持原样（这是题目"不足 k 个不翻转"的要求）；够 k 个就反转这一组，然后推进到下一组。
2. **内层**：反转一组 k 个节点。用第 206 题的标准迭代反转：用 `pre` 和 `cur` 两个指针，`cur` 逐个把节点头插到 `pre` 前面。

关键仍然是**边界接线**。每组反转前后涉及 4 个角色：

- `groupPrev`：上一组的尾节点（也就是本组的前驱），它的 `Next` 要改指本组反转后的头
- `groupHead`：本组的头节点，反转后变成本组的尾，它的 `Next` 要改指下一组的头
- 反转完成后，新的组头就是 `pre`（反转结束时 `pre` 指向本组最后一个节点），下一组的头是 `cur`

外层推进用 `groupPrev = groupHead`（反转后的组尾）即可继续处理下一组。

**为什么先检查剩余长度？** 题目要求不足 k 个的尾段保持原序。如果不管长度直接反转，最后一段被"残缺翻转"就错了。检查方法：从本组头开始数 k 个节点，中途遇到 `nil` 说明不足，直接结束。

画个图（`[1,2,3,4,5]`, `k=2`）：

```
初始:  dummy -> [1 -> 2] -> [3 -> 4] -> 5
              groupPrev   groupHead

第1组反转 [1,2]:
  反转后: dummy -> 2 -> 1 -> [3 -> 4] -> 5
         groupPrev=1(组尾) groupHead=3
第2组反转 [3,4]:
  反转后: dummy -> 2 -> 1 -> 4 -> 3 -> 5
剩余 5 不足 2 个，保持原样
结果: [2,1,4,3,5] ✓
```

### 算法步骤

1. 创建虚拟头节点 `dummy`，`dummy.Next = head`；`groupPrev` 从 `dummy` 出发
2. 循环：每次处理一组
   - 检查剩余节点是否够 k 个：从 `groupHead = groupPrev.Next` 开始数 k 个，不足则 break
   - 执行标准反转 k 个节点：`pre` 从 `nil` 开始，`cur` 从 `groupHead` 开始，逐个把 `cur` 头插到 `pre` 前，同时用 `nextNode` 缓存防止断链
   - 接线：`groupPrev.Next = pre`（接反转后的新头），`groupHead.Next = cur`（组尾接下一组的头，尾段不足 k 时 `cur` 就是剩余原序节点）
   - 推进：`groupPrev = groupHead`（移动到本组尾，即下一组的前驱）
3. 返回 `dummy.Next`

### 复杂度分析

- **时间复杂度**: O(n)。每组的检查最多看 k 个节点，反转 k 个节点，每个节点被常数次访问，总计 O(n)
- **空间复杂度**: O(1)。只使用常数个指针

## 代码实现

### 虚拟头节点 + 逐组反转

```go
func reverseKGroup(head *datastructures.ListNode, k int) *datastructures.ListNode {
    dummy := &datastructures.ListNode{Next: head}
    groupPrev := dummy // 上一组的尾节点（本组的前驱）
    for {
        // 第 1 步：检查剩余节点是否够 k 个
        groupHead := groupPrev.Next
        check := groupHead
        for i := 0; i < k; i++ {
            if check == nil {
                return dummy.Next // 不足 k 个，保持原样，结束
            }
            check = check.Next
        }
        // 第 2 步：标准迭代反转 k 个节点
        // 结束后 pre 指向本组新头，cur 指向下一组的头
        var pre *datastructures.ListNode
        cur := groupHead
        for i := 0; i < k; i++ {
            nextNode := cur.Next
            cur.Next = pre
            pre = cur
            cur = nextNode
        }
        // 第 3 步：接线——本组反转后头尾各接哪里
        groupPrev.Next = pre   // 前驱接本组新头
        groupHead.Next = cur   // 本组新尾（原头）接下一组的头
        // 第 4 步：推进到下一组
        groupPrev = groupHead
    }
}
```

**执行过程示例**（`head = [1,2,3,4,5]`, `k = 3`）：

```
初始:   dummy -> [1 -> 2 -> 3] -> 4 -> 5
检查: 从 1 开始数 3 个，够 → 反转 [1,2,3]:
  反转中: 1 摘下头插到 nil 前: nil <- 1, cur=2
          2 摘下头插到 1 前:   nil <- 2 <- 1, cur=3
          3 摘下头插到 2 前:   nil <- 3 <- 2 <- 1, cur=4
  接线: dummy.Next = 3, groupHead(1).Next = 4
结果: dummy -> [3 -> 2 -> 1] -> 4 -> 5
推进: groupPrev = 1

第 2 组: groupHead = 4
检查: 从 4 开始数 3 个，只有 4、5 两个就遇到 nil → 不足，结束
返回: [3,2,1,4,5] ✓
```
