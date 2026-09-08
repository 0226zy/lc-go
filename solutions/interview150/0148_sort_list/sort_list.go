package sortlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// SortList 排序链表
// 给你链表的头节点 head，请将其按升序排列后返回。
// 核心思路：归并排序的链表版——用快慢指针找到链表中点，将链表切成两半，
// 递归排序左右两半，再像「合并两个有序链表」一样合并。
// 时间复杂度: O(n log n)  空间复杂度: O(log n) 递归栈深度
func SortList(head *datastructures.ListNode) *datastructures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找中点：slow 最终指向后半段的头节点的前一个节点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil // 断开，分成两个子链表

	left := SortList(head)
	right := SortList(mid)
	return mergeTwoLists(left, right)
}

// mergeTwoLists 合并两个升序链表
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
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// SortListBottomUp 排序链表（自底向上归并，O(1) 额外空间）
// 从长度为 1 的子链表开始，逐轮将相邻的、长度相等的两个有序段两两合并，
// 段长度每轮翻倍：1, 2, 4, 8, ...，直到整表有序。全程只修改指针，不用递归。
// 时间复杂度: O(n log n)  空间复杂度: O(1)
func SortListBottomUp(head *datastructures.ListNode) *datastructures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &datastructures.ListNode{Next: head}
	// 先统计链表长度
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}
	for step := 1; step < length; step *= 2 {
		curr := dummy.Next
		tail := dummy // tail 指向已合并部分的末尾
		for curr != nil {
			// 切出两个长度最多为 step 的有序段 left、right
			left := curr
			right := cut(left, step)
			curr = cut(right, step)
			// 合并 left、right，接到 tail 后面
			tail.Next = mergeTwoLists(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}
	return dummy.Next
}

// cut 从 head 开始保留 n 个节点，返回第 n+1 个节点（即下一段的头），并断开连接
func cut(head *datastructures.ListNode, n int) *datastructures.ListNode {
	curr := head
	for i := 1; i < n && curr != nil; i++ {
		curr = curr.Next
	}
	if curr == nil {
		return nil
	}
	next := curr.Next
	curr.Next = nil
	return next
}
