# 92. 反转链表 II (Reverse Linked List II)

## 题目描述

给你单链表的头节点 `head` ，和两个整数 `left` 和 `right` ，其中 `left <= right` 。请你反转从位置 `left` 到位置 `right` 的链表节点，返回 **反转后的链表** 。

### 示例 1

```
输入: head = [1,2,3,4,5], left = 2, right = 4
输出: [1,4,3,2,5]
```

### 示例 2

```
输入: head = [5], left = 1, right = 1
输出: [5]
```

## 提示

- 链表中节点数目为 `n`
- `1 <= n <= 500`
- `-500 <= Node.val <= 500`
- `1 <= left <= right <= n`

## 题目解析

### 核心思路

第 206 题「反转链表」是反转整个链表，这题是它的升级版：**只反转中间某一段** `[left, right]`。

难点不在反转本身，而在**边界处理**——反转完之后，如何把这一段的"前后邻居"接回去。想象一根绳子，你把中间一截剪下来翻转后，必须把两端的线头重新接好：

- 第 `left - 1` 个节点（反转段的前驱）原来指着第 `left` 个节点，现在要改指第 `right` 个节点
- 第 `right + 1` 个节点（反转段的后继）位置不动，但它要成为反转后第 `left` 个节点的 `Next`

**穿针引线法**就是先把这三样东西固定住，再动中间的段：

1. 用**虚拟头节点**（dummy）统一处理 `left = 1`（即反转从头开始，前驱不存在）的边界；
2. 让 `prev` 走到第 `left - 1` 个节点，此时 `prev.Next` 是反转段的头（记为 `segmentHead`），`prev.Next.Next ...` 一路下去是反转段的剩余部分；
3. 对反转段做**头插法**：依次把 `segmentHead` 后面的节点摘下，插到 `prev` 的后面。重复 `right - left` 次后，这一段就反过来了，期间需要缓存反转段的后继节点 `nextNode`，防止断链后找不到；
4. 全部完成后 `prev.Next` 就是反转后的段头（原来的第 `right` 个节点），天然接好了前驱；头插过程中最后一个被移动的节点的 `Next` 指向原来的后继，也天然接好了后面。

画个图看最清楚（`[1,2,3,4,5]`, `left=2`, `right=4`）：

```
初始:   dummy -> 1 -> [2 -> 3 -> 4] -> 5
                prev  segmentHead

第1次头插(把3插到prev后面):
        dummy -> 1 -> 3 -> 2 -> 4 -> 5
                   prev       (4 被 2 指着，没有丢)

第2次头插(把4插到prev后面):
        dummy -> 1 -> 4 -> 3 -> 2 -> 5
                   prev
完成:  2 是反转段新尾，它的 Next 仍指向 5，后继接好了
       prev.Next = 4，前驱接好了
结果: [1,4,3,2,5] ✓
```

### 算法步骤

1. 创建虚拟头节点 `dummy`，`dummy.Next = head`；`prev` 从 `dummy` 出发
2. 让 `prev` 前进 `left - 1` 步，停在第 `left - 1` 个节点处
3. 记 `segmentHead = prev.Next`（反转段的头），执行 `right - left` 次头插：
   - `moved := segmentHead.Next`（待移动节点）
   - `segmentHead.Next = moved.Next`（从原位置摘下）
   - `moved.Next = prev.Next`（插到 prev 后面）
   - `prev.Next = moved`
4. 返回 `dummy.Next`

### 复杂度分析

- **时间复杂度**: O(n)，定位前驱 O(left)，头插 O(right - left)，合计 O(n)
- **空间复杂度**: O(1)，只使用常数个指针

## 代码实现

### 穿针引线（虚拟头节点 + 头插法）

```go
func reverseBetween(head *datastructures.ListNode, left, right int) *datastructures.ListNode {
    dummy := &datastructures.ListNode{Next: head} // 虚拟头节点，统一处理 left=1
    prev := dummy
    // prev 走到反转段的前驱（第 left-1 个节点）
    for i := 0; i < left-1; i++ {
        prev = prev.Next
    }
    // segmentHead 是反转段的头，反转后它将变成反转段的尾
    segmentHead := prev.Next
    // 头插 right-left 次，把 segmentHead 后面的节点逐个提到 prev 后面
    for i := 0; i < right-left; i++ {
        moved := segmentHead.Next   // 待移动节点
        segmentHead.Next = moved.Next // 摘下，保持不断链
        moved.Next = prev.Next      // 插到 prev 后面
        prev.Next = moved
    }
    return dummy.Next
}
```

**为什么只需 `right - left` 次而不是 `right - left + 1` 次？** 段内有 `right - left + 1` 个节点，但头插法以段头为基准移动后面的节点，第一个节点无需移动——移动 `right - left` 个节点后整段即完成反转。

**边界为什么不会出错？**
- `left = 1`：`prev` 停在 dummy 上，反转后 dummy.Next 自动指向新的头，无需特判；
- `left = right`：循环执行 0 次，原样返回；
- `right = n`：头插时 `moved.Next` 最终落到 `nil`，链表正确收尾。
