package mergetwosortedlists

import "github.com/0226zy/lc-go/pkg/datastructures"

// MergeTwoListsIterative 合并两个有序链表（迭代法）
// 将两个升序链表合并为一个新的升序链表并返回。
// 时间复杂度: O(m+n)  空间复杂度: O(1)
func MergeTwoListsIterative(list1 *datastructures.ListNode, list2 *datastructures.ListNode) *datastructures.ListNode {
	// 创建哑节点作为新链表的头部
	dummy := &datastructures.ListNode{}
	current := dummy

	// 遍历两个链表，将较小值的节点接入新链表
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// 将剩余的非空链表直接接入
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}

// MergeTwoListsRecursive 合并两个有序链表（递归法）
// 将两个升序链表合并为一个新的升序链表并返回。
// 时间复杂度: O(m+n)  空间复杂度: O(m+n)
func MergeTwoListsRecursive(list1 *datastructures.ListNode, list2 *datastructures.ListNode) *datastructures.ListNode {
	// 递归终止条件：如果其中一个链表为空，返回另一个链表
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// 比较两个链表头节点的值，递归合并剩余部分
	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}
