# 146. LRU 缓存 (LRU Cache)

## 题目描述

请你设计并实现一个满足 LRU（最近最少使用）缓存约束的数据结构。

实现 `LRUCache` 类：

- `LRUCache(int capacity)` 以正整数作为容量 `capacity` 初始化 LRU 缓存
- `int get(int key)` 如果关键字 `key` 存在于缓存中，则返回关键字的值，否则返回 `-1`。
- `void put(int key, int value)` 如果关键字 `key` 已经存在，则变更其数据值 `value`；如果不存在，则向缓存中插入该组 `key-value`。如果插入操作导致关键字数量超过 `capacity`，则应该逐出**最久未使用**的关键字。

函数 `get` 和 `put` 必须以 `O(1)` 的平均时间复杂度运行。

### 示例

```
输入:
["LRUCache", "put", "put", "get", "put", "put", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [4, 4], [1]]

输出:
[null, null, null, 1, null, null, -1]

解释:
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1（未找到）
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1（未找到）
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
```

### 提示

- `1 <= capacity <= 3000`
- `0 <= key <= 10000`
- `0 <= value <= 10^5`
- 最多调用 `2 * 10^5` 次 `get` 和 `put`

## 题目解析

### 核心思路

LRU 要支持两个看似矛盾的操作：

1. **按 key 快速定位**（O(1) 查找）→ 自然想到**哈希表**。
2. **知道「谁最久没用」并把它踢掉**（O(1) 淘汰）→ 需要维护一个**按使用先后排列的顺序结构**。

二者结合的经典方案：**哈希表 + 双向链表**。

- 哈希表 `map[int]*lruNode`：key 直接映射到链表节点，实现 O(1) 定位。
- 双向链表：把所有缓存项串起来，**表头方向是最久未使用，表尾方向是最近使用**。每次 `get` 命中或 `put` 已存在的 key，就把对应节点摘下来插到表尾；`put` 新 key 且超容量时，直接删掉头节点后第一个真实节点即可——它必然是最久未使用的。

**为什么必须用「双向」链表？** 淘汰一个节点时，需要修改它前驱的 `next` 和后继的 `prev`。单向链表没有 `prev` 指针，要从头找前驱，淘汰操作就退化成了 O(n)。双向链表拿到节点就能 O(1) 摘除自己，这是它存在的唯一理由。

实现上再加两个**哨兵节点**（dummy head / dummy tail），让「摘除」「插入表尾」「淘汰头节点」都不需要为边界（空链表、单节点）写特判，代码更不易错。节点上还要**额外记一个 key 字段**：淘汰节点时得知道它是哪个 key，才能顺手从哈希表里删掉。

### 算法步骤

`get(key)`：
1. 哈希表查不到 → 返回 -1。
2. 查到节点 → 把它摘除并插到链表尾部（标记为最近使用），返回值。

`put(key, value)`：
1. key 已存在 → 更新值，并移到链表尾部，结束。
2. key 不存在 → 新建节点，入哈希表、插到链表尾部；若数量超过容量，删掉头节点后的第一个真实节点（最久未使用），同时从哈希表删除其 key。

### 复杂度分析

- **时间复杂度**: `get` / `put` 均为 O(1)（哈希表定位 O(1)，双向链表摘除/插入 O(1)）。
- **空间复杂度**: O(capacity)，缓存项数不超过容量。

## 代码实现

```go
type lruNode struct {
    key, val   int
    prev, next *lruNode
}

type LRUCache struct {
    capacity int
    size     int
    cache    map[int]*lruNode
    head     *lruNode // 哨兵头：紧邻的是最久未使用的节点
    tail     *lruNode // 哨兵尾：紧邻的是最近使用的节点
}

func Constructor(capacity int) LRUCache {
    head, tail := &lruNode{}, &lruNode{}
    head.next, tail.prev = tail, head
    return LRUCache{capacity: capacity, cache: make(map[int]*lruNode), head: head, tail: tail}
}

func (c *LRUCache) Get(key int) int {
    if node, ok := c.cache[key]; ok {
        c.moveToTail(node) // 命中后挪到表尾
        return node.val
    }
    return -1
}

func (c *LRUCache) Put(key, value int) {
    if node, ok := c.cache[key]; ok {
        node.val = value
        c.moveToTail(node)
        return
    }
    node := &lruNode{key: key, val: value}
    c.cache[key] = node
    c.addToTail(node)
    c.size++
    if c.size > c.capacity {
        removed := c.head.next // 最久未使用
        c.removeNode(removed)
        delete(c.cache, removed.key)
        c.size--
    }
}
```

**执行过程示例**（容量 2，`put(1,1) put(2,2) get(1) put(3,3)`，左侧为最久未使用）：

```
put(1,1): 链表: head <-> [1] <-> tail
put(2,2): 链表: head <-> [1] <-> [2] <-> tail
get(1):   [1] 挪到表尾: head <-> [2] <-> [1] <-> tail, 返回 1
put(3,3): 先插入: head <-> [2] <-> [1] <-> [3] <-> tail, 超出容量
          淘汰头后第一个节点 [2]: head <-> [1] <-> [3] <-> tail
之后 get(2) 返回 -1，符合预期
```
