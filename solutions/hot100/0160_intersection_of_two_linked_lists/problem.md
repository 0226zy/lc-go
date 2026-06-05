# 160. 相交链表 (Intersection of Two Linked Lists)

## 题目描述

给你两个单链表的头节点 `headA` 和 `headB` ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 `null` 。

题目数据 **保证** 整个链式结构中不存在环。

**注意**，函数返回结果后，链表必须 **保持其原始结构** 。

### 示例 1

```
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,6,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
```

### 示例 2

```
输入：intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [1,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
```

### 示例 3

```
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：No intersection
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。
```

## 提示

- `listA` 中节点数目为 `m`
- `listB` 中节点数目为 `n`
- `1 <= m, n <= 3 * 10^4`
- `1 <= Node.val <= 10^5`
- `0 <= skipA < m` 和 `0 <= skipB < n`
- 如果 `listA` 和 `listB` 没有交点，`intersectVal` 为 `0`
- 如果 `listA` 和 `listB` 有交点，`intersectVal == listA[skipA] == listB[skipB]`

## 题目解析

### 核心思路

这道题要求找出两个单链表相交的起始节点。

**关键观察**：如果两个链表相交，那么从相交节点开始，后面的节点都是共享的。

**最优解法：双指针法**
1. 使用两个指针 `pA` 和 `pB`，分别指向 `headA` 和 `headB`
2. 遍历链表：
   - 如果 `pA` 到达末尾，则将其指向 `headB`
   - 如果 `pB` 到达末尾，则将其指向 `headA`
   - 如果 `pA` 和 `pB` 相遇，则相遇点就是相交节点（或者都是 `nil`，表示不相交）
3. 这个方法的原理是：两个指针走过的路径长度相同（都是 `m + n`），如果有相交节点，一定会在某点相遇

### 算法步骤（双指针法）

1. 初始化 `pA = headA`，`pB = headB`
2. 当 `pA != pB` 时：
   - 如果 `pA == nil`，则 `pA = headB`；否则 `pA = pA.Next`
   - 如果 `pB == nil`，则 `pB = headA`；否则 `pB = pB.Next`
3. 返回 `pA`（如果相交，则是相交节点；如果不相交，则是 `nil`）

### 复杂度分析

- **时间复杂度**: O(m+n)，两个指针最多遍历两个链表各一次
- **空间复杂度**: O(1)，只使用常数额外空间

## 代码实现

### 最优解：双指针法

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    
    pA, pB := headA, headB
    for pA != pB {
        if pA == nil {
            pA = headB
        } else {
            pA = pA.Next
        }
        if pB == nil {
            pB = headA
        } else {
            pB = pB.Next
        }
    }
    return pA
}
```

**执行过程示例**（`headA = [4,1,8,4,5]`, `headB = [5,6,1,8,4,5]`）：

```
初始: pA指向4, pB指向5

第1次循环: pA指向1, pB指向6
第2次循环: pA指向8, pB指向1
第3次循环: pA指向4, pB指向8
第4次循环: pA指向5, pB指向4
第5次循环: pA=nil→指向5, pB指向5
第6次循环: pA指向6, pB指向1
第7次循环: pA指向1, pB指向8
第8次循环: pA指向8, pB指向4
第9次循环: pA指向4, pB指向5
第10次循环: pA指向5, pB=nil→指向4
第11次循环: pA指向6, pB指向1
...

实际上，当 pA 和 pB 都走了 m+n 步后，它们会同时到达相交节点（或同时到达 nil）

简化理解:
pA 的路径: 4→1→8→4→5→nil→5→6→1→8 (相遇在8)
pB 的路径: 5→6→1→8→4→5→nil→4→1→8 (相遇在8)

结果: 相交节点值为 8
```

**性能优势**：
- O(m+n) 时间复杂度，一次遍历完成
- O(1) 空间复杂度，不使用额外数据结构
- 代码简洁，只需几行代码
