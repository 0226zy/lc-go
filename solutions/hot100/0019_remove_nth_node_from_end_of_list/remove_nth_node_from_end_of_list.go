package removenthnodefromendoflist

import "github.com/0226zy/lc-go/pkg/datastructures"

// RemoveNthNodeFromEndOfList 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 时间复杂度:   空间复杂度:
func RemoveNthNodeFromEndOfList(
	head *datastructures.ListNode, n int,
) *datastructures.ListNode {
	if head == nil {
		return nil
	}
	dummy := &datastructures.ListNode{0, head}
	left, right := head, dummy
	for i := 0; i < n; i++ {
		left = left.Next
	}
	for ; left != nil; left = left.Next {
		right = right.Next
	}
	right.Next = right.Next.Next
	return dummy.Next

}
