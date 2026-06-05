package reverselinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// ReverseList 反转链表
// 给你单链表的头节点 head，请你反转链表，并返回反转后的链表。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ReverseList(head *datastructures.ListNode) *datastructures.ListNode {
	var prev *datastructures.ListNode
	curr := head

	for curr != nil {
		next := curr.Next // 保存下一个节点
		curr.Next = prev  // 反转指针
		prev = curr       // 移动 prev
		curr = next       // 移动 curr
	}

	return prev
}
