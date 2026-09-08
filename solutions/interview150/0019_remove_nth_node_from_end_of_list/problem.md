# 19. 删除链表的倒数第 N 个结点 (Remove Nth Node From End of List)

## 题目描述

给你一个链表，删除链表的倒数第 `n` 个结点，并且返回链表的头结点。

### 示例 1

```
输入: head = [1,2,3,4,5], n = 2
输出: [1,2,3,5]
```

### 示例 2

```
输入: head = [1], n = 1
输出: []
```

### 示例 3

```
输入: head = [1,2], n = 1
输出: [1]
```

### 提示

- 链表中结点的数目为 `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## 题目解析

### 核心思路

最直觉的做法是先遍历一遍求出链表长度 `L`，再从头走 `L - n` 步找到待删节点的前驱，删掉它。这要做两次遍历。

更快的做法是**快慢指针，一次遍历搞定**。想象两个人站在链表起点，一个人（fast）先走 `n + 1` 步，然后两人一起走。因为两人始终保持着 `n + 1` 的间隔，当 fast 走到链表末尾（nil）时，slow 恰好停在**倒数第 n 个节点的前驱**上——这时直接让 `slow.Next = slow.Next.Next` 就删掉了目标节点。

为什么慢 `n + 1` 步而不是 `n` 步？因为要删一个节点，必须拿到它的**前驱**。fast 领先 slow 共 `n + 1` 个间隔，从 slow 的位置往后数第 `n + 1` 个位置就是 nil，说明 slow 后面恰好还有 n 个节点，slow.Next 正是倒数第 n 个。

另外引入一个**虚拟头节点**（dummy），把 `head` 接在 dummy 后面。这样即使要删的是头节点（比如 `n = sz`），slow 也能安全地停在 dummy 上统一处理，不需要特判。

### 算法步骤

1. 创建虚拟头节点 `dummy`，`dummy.Next = head`。
2. `slow` 和 `fast` 都指向 `dummy`。
3. `fast` 先向前移动 `n + 1` 步。
4. `slow`、`fast` 同时前进，直到 `fast == nil`。
5. 此时 `slow.Next` 就是倒数第 n 个节点，执行 `slow.Next = slow.Next.Next`。
6. 返回 `dummy.Next`。

### 复杂度分析

- **时间复杂度**: O(L)，一次遍历链表，L 为链表长度。
- **空间复杂度**: O(1)，只使用了两个指针。

## 代码实现

```go
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slow, fast := dummy, dummy
    // fast 先走 n+1 步，与 slow 保持 n+1 的间隔
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }
    // 同时前进，fast 到 nil 时 slow 指向倒数第 n 个节点的前驱
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

**执行过程示例**（`head = [1,2,3,4,5], n = 2`）：

```
初始: dummy -> 1 -> 2 -> 3 -> 4 -> 5, slow=fast=dummy
fast 先走 3 步: fast 到达节点 3, slow 仍在 dummy
同时前进:
  第1轮: slow=1, fast=4
  第2轮: slow=2, fast=5
  第3轮: slow=3, fast=nil
slow=3, slow.Next=4 即倒数第 2 个节点, 删除它
结果: 1 -> 2 -> 3 -> 5
```
