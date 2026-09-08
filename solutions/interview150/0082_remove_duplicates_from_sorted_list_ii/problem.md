# 82. 删除排序链表中的重复元素 II (Remove Duplicates from Sorted List II)

## 题目描述

给定一个已排序的链表的头 `head`，**删除原始链表中所有重复数字的节点**，只留下不同的数字。返回已排序的链表。

注意与第 83 题的区别：83 题保留一个重复元素，本题是「凡重复过的元素一个都不留」。

### 示例 1

```
输入: head = [1,2,3,3,4,4,5]
输出: [1,2,5]
```

### 示例 2

```
输入: head = [1,1,1,2,3]
输出: [2,3]
```

### 提示

- 链表中节点数目在范围 `[0, 300]` 内
- `-100 <= Node.val <= 100`
- 题目数据保证链表已经按升序排列

## 题目解析

### 核心思路

链表已经有序，所以**值相同的节点一定是连续的**。本题的关键是：发现一个重复段时，要把整段**全部删掉**，一个不留；而且它可能出现在链表的任何位置（包括头），所以需要虚拟头节点来统一处理。

维护两个指针：

- `prev`：指向**已确认保留区间**的最后一个节点（初始为虚拟头节点）。`prev.Next` 的含义是「下一个被确认保留的节点接在这里」。
- `curr`：负责向后扫描。

扫描时两种情况：

1. `curr.Next` 存在且值与 `curr` 相同 → `curr` 处在一个重复段里。不断让 `curr` 前进，直到跳过整个值相同的段（此时 `curr` 停在段的下一个节点，或 nil）。然后执行 `prev.Next = curr` 直接越过整段——这段节点一个都不保留，`prev` **不动**，因为 `curr` 指向的新节点还没有被确认。
2. `curr` 的值与后继不同 → `curr` 是独一无二的安全节点，把 `prev` 推进到 `curr`，`curr` 前进一格。

### 算法步骤

1. 创建虚拟头节点 `dummy`，`prev = dummy`，`curr = head`。
2. 当 `curr != nil`：
   - 若 `curr.Next` 存在且值等于 `curr.Val`：记录该值，循环把 `curr` 推进到第一个不等于该值的节点（或 nil），然后 `prev.Next = curr`。
   - 否则：`prev = curr`，`curr = curr.Next`。
3. 返回 `dummy.Next`。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点最多被访问两次。
- **空间复杂度**: O(1)，只使用了常数指针。

## 代码实现

```go
func DeleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    prev := dummy // 已确认保留区间的最后一个节点
    curr := head
    for curr != nil {
        if curr.Next != nil && curr.Next.Val == curr.Val {
            val := curr.Val
            // 跳过整个重复段，一个都不留
            for curr != nil && curr.Val == val {
                curr = curr.Next
            }
            prev.Next = curr
        } else {
            prev = curr // curr 独一无二，纳入保留区间
            curr = curr.Next
        }
    }
    return dummy.Next
}
```

**执行过程示例**（`head = [1,2,3,3,4,4,5]`）：

```
dummy -> 1 -> 2 -> 3 -> 3 -> 4 -> 4 -> 5
curr=1: 后继 2 ≠ 1，保留, prev=1
curr=2: 后继 3 ≠ 2，保留, prev=2
curr=3: 后继也是 3，发现重复段！跳过两个 3，curr 停在 4，prev.Next=4（2 直接连到 4，两个 3 被删）
curr=4: 后继也是 4，跳过两个 4，curr 停在 5，prev.Next=5（2 直接连到 5）
curr=5: 后继 nil，保留, prev=5, curr=nil，结束
结果: 1 -> 2 -> 5
```
