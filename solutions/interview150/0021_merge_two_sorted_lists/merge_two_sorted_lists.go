package mergetwosortedlists

import "github.com/0226zy/lc-go/pkg/datastructures"

// MergeTwoLists 合并两个有序链表（迭代法）
// 将两个升序链表合并为一个升序链表并返回。
// 每次比较两个链表的当前头节点，取较小者接到结果链表尾部；
// 循环结束后剩余部分天然有序，整体接上即可。
// 时间复杂度: O(m+n)  空间复杂度: O(1)
func MergeTwoLists(list1, list2 *datastructures.ListNode) *datastructures.ListNode {
	dummy := &datastructures.ListNode{} // 虚拟头节点，统一处理首节点插入
	curr := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	// 至多一个链表非空，剩余部分天然有序，整体接上
	if list1 != nil {
		curr.Next = list1
	} else {
		curr.Next = list2
	}
	return dummy.Next
}

// MergeTwoListsRec 合并两个有序链表（递归法）
// 每次取较小的头节点，剩余部分递归合并后接到其后面。
// 时间复杂度: O(m+n)  空间复杂度: O(m+n) 递归调用栈深度
func MergeTwoListsRec(list1, list2 *datastructures.ListNode) *datastructures.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRec(list1.Next, list2)
		return list1
	}
	list2.Next = MergeTwoListsRec(list1, list2.Next)
	return list2
}
