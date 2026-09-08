package lrucache

// lruNode 双向链表节点，同时记录 key 以便淘汰时从哈希表中删除
type lruNode struct {
	key, val   int
	prev, next *lruNode
}

// LRUCache LRU 缓存
// 运用你掌握的数据结构，设计和实现一个 LRU（最近最少使用）缓存。
// get 和 put 必须以 O(1) 的平均时间复杂度运行。
type LRUCache struct {
	capacity int
	size     int
	cache    map[int]*lruNode
	head     *lruNode // 哨兵头，紧邻 head 的节点是「最久未使用」的
	tail     *lruNode // 哨兵尾，紧邻 tail 的节点是「最近使用」的
}

// Constructor 创建指定容量的 LRU 缓存
// 时间复杂度: O(1)  空间复杂度: O(capacity)
func Constructor(capacity int) LRUCache {
	head := &lruNode{}
	tail := &lruNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*lruNode),
		head:     head,
		tail:     tail,
	}
}

// Get 获取 key 对应的值，不存在返回 -1，命中后将节点移到链表尾部
// 时间复杂度: O(1)  空间复杂度: O(1)
func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		c.moveToTail(node)
		return node.val
	}
	return -1
}

// Put 写入或更新 key-value，超出容量时淘汰最久未使用的节点
// 时间复杂度: O(1)  空间复杂度: O(1)
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
		// 淘汰头节点后的第一个真实节点（最久未使用）
		removed := c.head.next
		c.removeNode(removed)
		delete(c.cache, removed.key)
		c.size--
	}
}

// moveToTail 将已存在的节点摘下并插到链表尾部（标记为最近使用）
func (c *LRUCache) moveToTail(node *lruNode) {
	c.removeNode(node)
	c.addToTail(node)
}

// addToTail 把节点插到哨兵尾之前
func (c *LRUCache) addToTail(node *lruNode) {
	prev := c.tail.prev
	prev.next = node
	node.prev = prev
	node.next = c.tail
	c.tail.prev = node
}

// removeNode 从链表中摘下节点
func (c *LRUCache) removeNode(node *lruNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
