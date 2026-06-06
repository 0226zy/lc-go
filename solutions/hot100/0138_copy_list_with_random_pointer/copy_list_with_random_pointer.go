package copyrandomlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// CopyRandomListHashTable 随机链表的复制 - 哈希表法
// 给定一个随机链表，每个节点包含 Next 和 Random 指针，返回该链表的深拷贝。
// 时间复杂度: O(n)  空间复杂度: O(n)
func CopyRandomListHashTable(head *datastructures.RandomListNode) *datastructures.RandomListNode {
	if head == nil {
		return nil
	}

	// 哈希表：原节点 -> 新节点
	hashTable := make(map[*datastructures.RandomListNode]*datastructures.RandomListNode)

	// 第一次遍历：创建所有新节点，建立映射关系
	curr := head
	for curr != nil {
		hashTable[curr] = &datastructures.RandomListNode{Val: curr.Val}
		curr = curr.Next
	}

	// 第二次遍历：设置新节点的 Next 和 Random 指针
	curr = head
	for curr != nil {
		newNode := hashTable[curr]
		// 设置 Next 指针
		if curr.Next != nil {
			newNode.Next = hashTable[curr.Next]
		}
		// 设置 Random 指针
		if curr.Random != nil {
			newNode.Random = hashTable[curr.Random]
		}
		curr = curr.Next
	}

	return hashTable[head]
}

// CopyRandomListInPlace 随机链表的复制 - 原地复制法（最优解）
// 给定一个随机链表，每个节点包含 Next 和 Random 指针，返回该链表的深拷贝。
// 时间复杂度: O(n)  空间复杂度: O(1)
func CopyRandomListInPlace(head *datastructures.RandomListNode) *datastructures.RandomListNode {
	if head == nil {
		return nil
	}

	// 第一步：在原节点后插入复制节点
	// 原链表：A -> B -> C -> null
	// 复制后：A -> A' -> B -> B' -> C -> C' -> null
	curr := head
	for curr != nil {
		newNode := &datastructures.RandomListNode{Val: curr.Val}
		newNode.Next = curr.Next
		curr.Next = newNode
		curr = newNode.Next
	}

	// 第二步：设置复制节点的 Random 指针
	// 对于原节点 A，其复制节点 A' 的 Random 应该是 A.Random 的复制节点
	curr = head
	for curr != nil {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
		curr = curr.Next.Next
	}

	// 第三步：拆分链表，恢复原链表并提取复制链表
	curr = head
	newHead := head.Next
	for curr != nil {
		newNode := curr.Next
		curr.Next = newNode.Next // 恢复原链表
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next // 连接复制链表
		}
		curr = curr.Next
	}

	return newHead
}
