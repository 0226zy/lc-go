package sortlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// SortListTopDown 排序链表 - 自顶向下归并排序
// 给定链表的头节点，返回排序后的链表（升序）。
// 时间复杂度: O(n log n)  空间复杂度: O(log n)
func SortListTopDown(head *datastructures.ListNode) *datastructures.ListNode {
	// 边界条件：空链表或单节点链表
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表中点
	slow, fast := head, head
	var prev *datastructures.ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 断开链表
	prev.Next = nil
	mid := slow

	// 递归排序左右两部分
	left := SortListTopDown(head)
	right := SortListTopDown(mid)

	// 合并两个有序链表
	return mergeTwoLists(left, right)
}

// SortListBottomUp 排序链表 - 自底向上归并排序（最优解）
// 给定链表的头节点，返回排序后的链表（升序）。
// 时间复杂度: O(n log n)  空间复杂度: O(1)
func SortListBottomUp(head *datastructures.ListNode) *datastructures.ListNode {
	// 边界条件：空链表或单节点链表
	if head == nil || head.Next == nil {
		return head
	}

	// 计算链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	// 创建哑节点
	dummy := &datastructures.ListNode{Next: head}

	// 自底向上归并
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev := dummy
		curr := dummy.Next

		// 每次处理两个长度为 subLength 的子链表
		for curr != nil {
			// 获取第一个子链表
			head1 := curr
			for i := 1; i < subLength && curr.Next != nil; i++ {
				curr = curr.Next
			}

			// 获取第二个子链表
			head2 := curr.Next
			curr.Next = nil // 断开第一个子链表
			curr = head2
			for i := 1; i < subLength && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}

			// 保存下一个子链表的头节点
			var next *datastructures.ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil // 断开第二个子链表
			}

			// 合并两个子链表
			merged := mergeTwoLists(head1, head2)
			prev.Next = merged

			// 移动到合并后链表的尾部
			for prev.Next != nil {
				prev = prev.Next
			}

			curr = next
		}
	}

	return dummy.Next
}

// mergeTwoLists 合并两个有序链表（辅助函数）
func mergeTwoLists(l1, l2 *datastructures.ListNode) *datastructures.ListNode {
	dummy := &datastructures.ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
