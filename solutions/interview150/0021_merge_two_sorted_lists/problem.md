# 21. 合并两个有序链表 (Merge Two Sorted Lists)

## 题目描述

将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

### 示例 1

```
输入: list1 = [1,2,4], list2 = [1,3,4]
输出: [1,1,2,3,4,4]
```

### 示例 2

```
输入: list1 = [], list2 = []
输出: []
```

### 示例 3

```
输入: list1 = [], list2 = [0]
输出: [0]
```

## 提示

- 两个链表的节点数目范围是 `[0, 50]`
- `-100 <= Node.val <= 100`
- `list1` 和 `list2` 均按 **非递减顺序** 排列

## 题目解析

### 核心思路

这是「**有序链表双指针归并**」模板题，本质上是归并排序中"归并"那一步的链表版本。

两个链表都已经是升序的，就像两排已经各自排好队的人，要合成一队：每次只需比较两队的队首，谁小谁就出列站到新队伍里。由于两队内部有序，这个贪心选择（每次取当前最小）一定正确。

有两种写法：

1. **迭代法**：用一个**虚拟头节点**（dummy）统一处理新链表的拼接，每次把较小的节点接到 `curr` 后面。循环结束后，哪个链表还有剩余，直接整体接上去即可（因为剩余部分天然有序，不需要再逐个搬运）。
2. **递归法**：合并 `list1` 和 `list2` 的结果 = 较小头节点 + 递归合并剩余部分。天然贴合"每次取较小者"的定义，代码更短。

### 算法步骤（迭代法）

1. 创建虚拟头节点 `dummy` 和尾指针 `curr`，指向 `dummy`
2. 当 `list1` 和 `list2` 都不为空：
   - 若 `list1.Val <= list2.Val`，把 `list1` 接到 `curr.Next`，`list1` 前进一步；否则接 `list2`
   - `curr` 前进一步
3. 循环结束后，至多一个链表非空，把剩余部分直接接到 `curr.Next`
4. 返回 `dummy.Next`

### 复杂度分析

- **时间复杂度**: O(m + n)，m、n 为两条链表长度，每个节点最多被访问一次
- **空间复杂度**:
  - 迭代法：O(1)，只占用若干指针
  - 递归法：O(m + n)，递归调用栈的深度

## 代码实现

### 迭代法（虚拟头节点）

```go
func mergeTwoLists(list1, list2 *datastructures.ListNode) *datastructures.ListNode {
    dummy := &datastructures.ListNode{} // 虚拟头节点
    curr := dummy
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }
    // 剩余部分天然有序，整体接上
    if list1 != nil {
        curr.Next = list1
    } else {
        curr.Next = list2
    }
    return dummy.Next
}
```

### 递归法

```go
func mergeTwoListsRec(list1, list2 *datastructures.ListNode) *datastructures.ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    if list1.Val <= list2.Val {
        list1.Next = mergeTwoListsRec(list1.Next, list2)
        return list1
    }
    list2.Next = mergeTwoListsRec(list1, list2.Next)
    return list2
}
```

**执行过程示例**（`list1 = [1,2,4]`，`list2 = [1,3,4]`，迭代法）：

```
初始: dummy -> nil, curr=dummy
比较 1<=1: 接 list1 的 1, list1 -> [2,4]   结果: [1]
比较 2>1:  接 list2 的 1, list2 -> [3,4]   结果: [1,1]
比较 2<=3: 接 list1 的 2, list1 -> [4]      结果: [1,1,2]
比较 4>3:  接 list2 的 3, list2 -> [4]      结果: [1,1,2,3]
比较 4<=4: 接 list1 的 4, list1 -> nil      结果: [1,1,2,3,4]
list2 剩余 [4]，整体接上                  结果: [1,1,2,3,4,4]
```
