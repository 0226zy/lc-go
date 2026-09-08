package removeduplicatesii

import "github.com/0226zy/lc-go/pkg/datastructures"

// DeleteDuplicates 删除排序链表中的重复元素 II
// 给定一个已排序的链表的头 head，删除原始链表中所有重复数字的节点，只留下不同的数字。
// 返回已排序的链表。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 只使用常数指针
func DeleteDuplicates(head *datastructures.ListNode) *datastructures.ListNode {
	dummy := &datastructures.ListNode{Next: head}
	prev := dummy // prev 指向已确认「保留」区间的最后一个节点
	curr := head
	for curr != nil {
		// 发现连续重复段：curr 及其后继值相同
		if curr.Next != nil && curr.Next.Val == curr.Val {
			val := curr.Val
			// 把整个重复段跳过去，一个都不留
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			// 保留区间直接接到重复段之后的节点
			prev.Next = curr
		} else {
			// curr 是独一无二的节点，纳入保留区间
			prev = curr
			curr = curr.Next
		}
	}
	return dummy.Next
}
