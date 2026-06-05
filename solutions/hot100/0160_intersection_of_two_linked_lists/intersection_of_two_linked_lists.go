package intersectionoftwolinkedlists

import "github.com/0226zy/lc-go/pkg/datastructures"

// GetIntersectionNode 相交链表
// 给你两个单链表的头节点 headA 和 headB，请你找出并返回两个单链表相交的起始节点。
// 如果两个链表不存在相交节点，返回 nil。
// 时间复杂度: O(m+n) 两个指针最多遍历两个链表各一次  空间复杂度: O(1) 常数空间
func GetIntersectionNode(headA, headB *datastructures.ListNode) *datastructures.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
