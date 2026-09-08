package removenthnode

import "github.com/0226zy/lc-go/pkg/datastructures"

// RemoveNthFromEnd 删除链表的倒数第 N 个结点
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 时间复杂度: O(L) 一次遍历，L 为链表长度  空间复杂度: O(1) 只使用常数指针
func RemoveNthFromEnd(head *datastructures.ListNode, n int) *datastructures.ListNode {
	// 虚拟头节点，简化对头节点的删除（删除头节点时无需特判）
	dummy := &datastructures.ListNode{Next: head}
	slow, fast := dummy, dummy
	// fast 先走 n+1 步，使 slow 与 fast 之间保持 n+1 的间隔
	// 这样当 fast 走到 nil（末尾之后）时，slow 正好指向倒数第 n 个结点的前驱
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	// 快慢指针同时前进，直到 fast 到达 nil
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// 跳过待删除的倒数第 n 个结点
	slow.Next = slow.Next.Next
	return dummy.Next
}
