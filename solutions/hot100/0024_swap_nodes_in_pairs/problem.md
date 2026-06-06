# 24. 两两交换链表中的节点

## 题目描述

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。

**必须在不修改节点内部的值的情况下完成**（即，只能进行节点交换）。

### 示例 1
```
输入：head = [1,2,3,4]
输出：[2,1,4,3]
```

### 示例 2
```
输入：head = []
输出：[]
```

### 示例 3
```
输入：head = [1]
输出：[1]
```

### 提示
- 链表中节点的数目在范围 `[0, 100]` 内
- `0 <= Node.val <= 100`

---

## 解题思路

### 方法一：迭代法（使用哑节点）

**核心思想**：
使用哑节点（dummy node）简化头节点的处理，然后迭代地交换每两个相邻节点。

**算法步骤**：
1. 创建哑节点 `dummy`，令 `dummy.Next = head`
2. 初始化 `prev = dummy`，用于指向要交换的两个节点之前的位置
3. 当 `prev.Next` 和 `prev.Next.Next` 都不为空时，进行交换：
   - 设 `first = prev.Next`（第一个节点）
   - 设 `second = prev.Next.Next`（第二个节点）
   - 执行交换：
     - `first.Next = second.Next`
     - `second.Next = first`
     - `prev.Next = second`
   - 移动 `prev` 到 `first` 的位置（即交换后的第一个节点）
4. 返回 `dummy.Next`

**交换过程图示**：
```
初始: prev -> 1 -> 2 -> 3 -> 4
交换前: prev -> first(1) -> second(2) -> 3 -> 4
交换后: prev -> second(2) -> first(1) -> 3 -> 4
然后: prev 移动到 first(1) 的位置
```

**复杂度分析**：
- **时间复杂度**：O(n)，其中 n 是链表的节点数。需要遍历整个链表一次。
- **空间复杂度**：O(1)，只使用了常数级的额外空间。

---

### 方法二：递归法

**核心思想**：
递归地交换链表的前两个节点，然后递归地处理剩余部分。

**算法步骤**：
1. 递归终止条件：如果链表为空或只有一个节点，直接返回该链表
2. 设 `first = head`，`second = head.Next`
3. 递归交换从 `second.Next` 开始的剩余链表：`first.Next = swapPairs(second.Next)`
4. 将 `second.Next` 指向 `first`：`second.Next = first`
5. 返回 `second` 作为新的头节点

**递归过程图示**：
```
swapPairs([1,2,3,4])
  first = 1, second = 2
  first.Next = swapPairs([3,4])
    swapPairs([3,4])
      first = 3, second = 4
      first.Next = swapPairs([])
        swapPairs([]) 返回 nil
      first.Next = nil
      second.Next = first  =>  4 -> 3 -> nil
      返回 4
  first.Next = 4 -> 3 -> nil  =>  1 -> 4 -> 3 -> nil
  second.Next = first  =>  2 -> 1 -> 4 -> 3 -> nil
  返回 2
```

**复杂度分析**：
- **时间复杂度**：O(n)，其中 n 是链表的节点数。每个节点都会被处理一次。
- **空间复杂度**：O(n)，递归调用栈的深度为 n/2，即 O(n)。

---

## 知识点总结

1. **链表操作**：理解链表的节点连接操作和指针操作
2. **哑节点技巧**：简化链表头部的边界情况处理
3. **递归思想**：将大问题分解为相似的子问题
4. **指针操作**：注意指针赋值的顺序，避免丢失节点引用

---

## 相关题目

- [25. K 个一组翻转链表](https://leetcode.cn/problems/reverse-nodes-in-k-group/)（Hard）
- [206. 反转链表](https://leetcode.cn/problems/reverse-linked-list/)（Easy）
- [92. 反转链表 II](https://leetcode.cn/problems/reverse-linked-list-ii/)（Medium）
