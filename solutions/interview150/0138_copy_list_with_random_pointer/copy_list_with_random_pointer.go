package copylistwithrandompointer

import "github.com/0226zy/lc-go/pkg/datastructures"

// CopyRandomList 随机链表的复制
// 给定一个每个节点带 random 指针（可指向任意节点或空）的链表，返回其深拷贝。
// 第一遍遍历为每个原节点创建拷贝节点并存入哈希表（原节点 -> 拷贝节点），
// 第二遍遍历按原节点的 next/random 关系为拷贝节点填指针。
// 哈希表查不到 nil 键时返回 nil，因此空指针无需特判。
// 时间复杂度: O(n)  空间复杂度: O(n)
func CopyRandomList(head *datastructures.RandomListNode) *datastructures.RandomListNode {
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
		copyNode := nodeMap[curr]
		copyNode.Next = nodeMap[curr.Next]     // curr.Next 为 nil 时映射结果为 nil
		copyNode.Random = nodeMap[curr.Random] // curr.Random 为 nil 时映射结果为 nil
	}
	return nodeMap[head]
}
