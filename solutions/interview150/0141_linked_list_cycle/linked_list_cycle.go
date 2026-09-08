package linkedlistcycle

import "github.com/0226zy/lc-go/pkg/datastructures"

// HasCycle 环形链表
// 给定一个链表的头节点 head，判断链表中是否存在环。
// 使用 Floyd 快慢指针：快指针每次走两步，慢指针每次走一步，
// 若链表有环则两指针一定相遇，无环则快指针先到 nil。
// 时间复杂度: O(n)  空间复杂度: O(1)
func HasCycle(head *datastructures.ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 慢指针走一步
		fast = fast.Next.Next // 快指针走两步
		if slow == fast {     // 相遇说明有环
			return true
		}
	}
	return false // 快指针到达末尾，无环
}
