package swapnodesinpairs

import "github.com/0226zy/lc-go/pkg/datastructures"

// SwapPairsIterative 两两交换链表中的节点（迭代法）
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SwapPairsIterative(head *datastructures.ListNode) *datastructures.ListNode {
	// 创建哑节点，简化头节点的处理
	dummy := &datastructures.ListNode{Next: head}
	prev := dummy

	// 当还有至少两个节点可以交换时
	for prev.Next != nil && prev.Next.Next != nil {
		// 要交换的两个节点
		first := prev.Next
		second := prev.Next.Next

		// 执行交换
		first.Next = second.Next
		second.Next = first
		prev.Next = second

		// 移动 prev 指针，准备下一轮交换
		prev = first
	}

	return dummy.Next
}

// SwapPairsRecursive 两两交换链表中的节点（递归法）
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。
// 时间复杂度: O(n)  空间复杂度: O(n)
func SwapPairsRecursive(head *datastructures.ListNode) *datastructures.ListNode {
	// 递归终止条件：空链表或只有一个节点
	if head == nil || head.Next == nil {
		return head
	}

	// 要交换的两个节点
	first := head
	second := head.Next

	// 递归交换剩余部分
	first.Next = SwapPairsRecursive(second.Next)

	// 完成当前两个节点的交换
	second.Next = first

	// 返回交换后的头节点
	return second
}
