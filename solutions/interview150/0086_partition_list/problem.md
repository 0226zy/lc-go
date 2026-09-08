# 86. 分隔链表 (Partition List)

## 题目描述

给你一个链表的头节点 `head` 和一个特定值 `x`，请你对链表进行分隔，使得所有**小于 x** 的节点都出现在**大于或等于 x** 的节点之前。你应当**保留两个分区中每个节点的初始相对位置**。

### 示例 1

```
输入: head = [1,4,3,2,5,2], x = 3
输出: [1,2,2,4,3,5]
```

### 示例 2

```
输入: head = [2,1], x = 2
输出: [1,2]
```

### 提示

- 链表中节点的数目在范围 `[0, 200]` 内
- `-100 <= Node.val <= 100`
- `-200 <= x <= 200`

## 题目解析

### 核心思路

题目只要求「小的在前、大的在后」，**并不要求排序**。如果试图在原链表上交换节点，要处理大量指针回退，很容易写错。

干净的做法是「**双哑节点、拆成两条链再拼接**」：准备两个虚拟头节点——`smallDummy` 挂所有 `< x` 的节点，`bigDummy` 挂所有 `>= x` 的节点。从头到尾扫一遍原链表，每个节点按大小**依次摘下来挂到对应链尾**。因为我们是按原顺序收集的，每条链内部的相对顺序天然不变。

扫描结束后做两件事：

1. `big.Next = nil`：把大值链的尾巴切断。这一步**不能省**——big 可能是原链表的中间某个节点，它原来的 `Next` 还指着原链表后面的节点，不断开会出错甚至成环。
2. `small.Next = bigDummy.Next`：把小值链的头接到大值链的头。

整个过程不新建数据节点，只是把原节点重新串起来。

### 算法步骤

1. 创建 `smallDummy`、`bigDummy` 两个虚拟头节点，以及各自的尾指针 `small`、`big`。
2. 遍历原链表每个节点 `curr`：
   - `curr.Val < x`：`small.Next = curr`，`small = curr`。
   - 否则：`big.Next = curr`，`big = curr`。
3. 执行 `big.Next = nil`，切断大值链尾部。
4. 执行 `small.Next = bigDummy.Next`，拼接两条链。
5. 返回 `smallDummy.Next`。

### 复杂度分析

- **时间复杂度**: O(n)，一次遍历。
- **空间复杂度**: O(1)，没有新建数据节点（只用了两个虚拟头），纯属重排。

## 代码实现

```go
func Partition(head *ListNode, x int) *ListNode {
    smallDummy := &ListNode{} // 小于 x 的链
    bigDummy := &ListNode{}   // 大于等于 x 的链
    small, big := smallDummy, bigDummy
    for curr := head; curr != nil; curr = curr.Next {
        if curr.Val < x {
            small.Next = curr
            small = small.Next
        } else {
            big.Next = curr
            big = big.Next
        }
    }
    big.Next = nil          // 切断大值链尾部，防止残留旧链造成环
    small.Next = bigDummy.Next // 小值链在前，大值链在后
    return smallDummy.Next
}
```

**执行过程示例**（`head = [1,4,3,2,5,2], x = 3`）：

```
small 链: 依次收集 1, 2, 2  -> 1 -> 2 -> 2
big 链:   依次收集 4, 3, 5  -> 4 -> 3 -> 5
切断 big.Next = nil
拼接: small 尾(2) 的 Next 指向 big 链头(4)
结果: 1 -> 2 -> 2 -> 4 -> 3 -> 5
```
