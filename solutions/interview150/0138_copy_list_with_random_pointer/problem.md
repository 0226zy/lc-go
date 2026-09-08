# 138. 随机链表的复制 (Copy List with Random Pointer)

## 题目描述

给你一个长度为 `n` 的链表，每个节点包含一个额外增加的随机指针 `random` ，该指针可以指向链表中的任何节点或空节点。

构造这个链表的 **深拷贝**。深拷贝应该正好由 `n` 个 **全新** 节点组成，其中每个新节点的值设为其对应的原节点的值。新节点的 `next` 指针和 `random` 指针也都应指向复制链表中的新节点，并使上述指针在复制链表中也保持原来的结构。复制链表中的指针均不应指向原始链表中的节点 。

例如，如果原始链表中有 `X` 和 `Y` 两个节点，其中 `X.random --> Y` 。那么在复制链表中对应的两个节点 `x` 和 `y` ，同样有 `x.random --> y` 。

用一个由 `n` 个节点组成的链表来表示输入/输出。每个节点用一个 `[val, random_index]` 表示：

- `val`：一个表示 `Node.val` 的整数。
- `random_index`：随机指针指向的节点索引（**从 0 开始**）；如果不指向任何节点，则为 `null` 。

你的代码 **只** 接受原始链表的头节点 `head` 作为传入参数。

### 示例 1

```
输入: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

### 示例 2

```
输入: head = [[1,1],[2,1]]
输出: [[1,1],[2,1]]
```

### 示例 3

```
输入: head = [[3,null],[3,0],[3,null]]
输出: [[3,null],[3,0],[3,null]]
```

## 提示

- `0 <= n <= 1000`
- `-10^4 <= Node.val <= 10^4`
- `Node.random` 为 `null` 或指向链表中的节点

## 题目解析

### 核心思路

这是「**复杂链表的深拷贝**」问题。难点在于 `next` 指针还好办，但 `random` 指针可以指向**任意**节点——你在拷贝过程中遍历时，random 指向的节点可能还没被创建，就算已创建也需要快速找到"原节点的拷贝节点"是哪个。

核心矛盾一句话概括：**需要在"原节点"和"它的拷贝节点"之间建立快速的对应关系**。想到这一层，解法就呼之欲出了——用**哈希表**记录 `原节点 -> 拷贝节点` 的映射：

1. 第一遍遍历：为每个原节点创建拷贝节点（只填值，`next`/`random` 先留空），存入哈希表。
2. 第二遍遍历：对每个原节点，从哈希表取出它的拷贝节点，把 `next` 和 `random` 指针按照"原节点指向谁，拷贝节点就指向谁的拷贝"填上。

**为什么这个解法是对的？** 深拷贝的本质要求是：复制出的所有节点都是全新的，且节点间的相对引用关系与原链表完全一致。哈希表保证了每个原节点有且只有一个全新拷贝；第二遍填指针时，`copy.Next = map[curr.Next]`、`copy.Random = map[curr.Random]` 保证拷贝链表的指针全部指向拷贝节点，且结构与原来一一对应。

进阶技巧（面试加分项）：可以先把每个拷贝节点插到对应原节点后面（A -> A' -> B -> B'），形成交错链表，然后用 `curr.Next` 直接取到拷贝节点填 random，最后拆开成两个链表。这个方法不用哈希表，空间 O(1)，但写起来易错，工程上首选哈希表法，清晰稳妥。

### 算法步骤（哈希表法）

1. 若 `head` 为 `nil`，直接返回 `nil`
2. 第一遍遍历原链表：为每个节点 `curr` 创建拷贝 `&RandomListNode{Val: curr.Val}`，存入 `map[curr]copy`
3. 第二遍遍历原链表：对 `map[curr]` 填指针：
   - `copy.Next = map[curr.Next]`（`curr.Next` 为 `nil` 时映射为 `nil`）
   - `copy.Random = map[curr.Random]`（同上）
4. 返回 `map[head]`

### 复杂度分析

- **时间复杂度**: O(n)，两遍遍历
- **空间复杂度**: O(n)，哈希表存储 n 个原节点到拷贝节点的映射

## 代码实现

### 哈希表法

```go
func copyRandomList(head *datastructures.RandomListNode) *datastructures.RandomListNode {
    if head == nil {
        return nil
    }
    // 第一遍：为每个原节点创建拷贝节点，建立 原节点 -> 拷贝节点 的映射
    nodeMap := make(map[*datastructures.RandomListNode]*datastructures.RandomListNode)
    for curr := head; curr != nil; curr = curr.Next {
        nodeMap[curr] = &datastructures.RandomListNode{Val: curr.Val}
    }
    // 第二遍：按原节点的指针关系，为拷贝节点填 next 和 random
    for curr := head; curr != nil; curr = curr.Next {
        copy := nodeMap[curr]
        copy.Next = nodeMap[curr.Next]     // nil 键查不到，值为 nil，天然正确
        copy.Random = nodeMap[curr.Random]
    }
    return nodeMap[head]
}
```

**执行过程示例**（`[[7,null],[13,0],[11,4]]`）：

```
第一遍遍历后，nodeMap 中保存了映射关系（原节点 -> 拷贝节点）:
  7 -> 7', 13 -> 13', 11 -> 11'

第二遍遍历:
  原节点 7:  7'.Next = 13', 7'.Random = nil     （原 random 为 null）
  原节点 13: 13'.Next = 11', 13'.Random = 7'    （原 random 指向索引 0，即节点 7）
  原节点 11: 11'.Next = nil,  11'.Random = 11'   （原 random 指向索引 4，即自身）

结果: [[7,null],[13,0],[11,4]]，结构与输入完全一致 ✓
```

**注意**：`map` 查询不存在的键（包括 `nil` 键）会返回零值 `nil`，所以 `curr.Next == nil` 时 `nodeMap[curr.Next]` 直接得到 `nil`，无需特判，代码更简洁。
