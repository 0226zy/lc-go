package addtwonumbers

import "github.com/0226zy/lc-go/pkg/datastructures"

// AddTwoNumbers 两数相加
// 给定两个非空链表，逆序存储两个非负整数（个位在表头），返回两数之和的链表。
// 从头模拟竖式加法：逐位求和并处理进位，任一链表为空时其位按 0 计算，
// 遍历结束后若仍有进位需补一个节点。
// 时间复杂度: O(max(m,n))  空间复杂度: O(max(m,n))
func AddTwoNumbers(l1, l2 *datastructures.ListNode) *datastructures.ListNode {
	dummy := &datastructures.ListNode{} // 虚拟头节点，统一处理首节点插入
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		curr.Next = &datastructures.ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}
	if carry > 0 { // 最高位进位，补一个节点（如 99+1=100）
		curr.Next = &datastructures.ListNode{Val: carry}
	}
	return dummy.Next
}
