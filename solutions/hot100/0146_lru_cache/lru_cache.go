package lrucache

// LRUCache LRU 缓存
// 使用哈希表 + 双向链表实现，支持 O(1) 的 get 和 put 操作。
// 时间复杂度: Get O(1), Put O(1)  空间复杂度: O(capacity)
type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

// Constructor 初始化一个容量为 capacity 的 LRUCache
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node, capacity),
		head:     &node{},
		tail:     &node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get 获取 key 对应的值，若不存在返回 -1
func (l *LRUCache) Get(key int) int {
	if n, ok := l.cache[key]; ok {
		l.moveToHead(n)
		return n.value
	}
	return -1
}

// Put 插入或更新 key-value，若超出容量则淘汰最久未使用的节点
func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.cache[key]; ok {
		n.value = value
		l.moveToHead(n)
		return
	}

	n := &node{key: key, value: value}
	l.cache[key] = n
	l.addToHead(n)

	if len(l.cache) > l.capacity {
		removed := l.removeTail()
		delete(l.cache, removed.key)
	}
}

// addToHead 将节点插入到链表头部（head 之后）
func (l *LRUCache) addToHead(n *node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

// removeNode 从链表中移除指定节点
func (l *LRUCache) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// moveToHead 将已有节点移到链表头部
func (l *LRUCache) moveToHead(n *node) {
	l.removeNode(n)
	l.addToHead(n)
}

// removeTail 移除链表尾部节点（tail 之前），并返回该节点
func (l *LRUCache) removeTail() *node {
	n := l.tail.prev
	l.removeNode(n)
	return n
}
