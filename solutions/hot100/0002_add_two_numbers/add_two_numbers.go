package addtwonumbers

import "github.com/0226zy/lc-go/pkg/datastructures"

// AddTwoNumbers 两数相加
// 给定两个非空链表，表示两个非负整数（逆序存储），将两数相加并以链表形式返回。
// 时间复杂度: O(max(m,n))  空间复杂度: O(max(m,n))
func AddTwoNumbers(l1 *datastructures.ListNode, l2 *datastructures.ListNode) *datastructures.ListNode {
	var newHead, newLast *datastructures.ListNode
	flag := 0
	h1, h2 := l1, l2
	for h1 != nil || h2 != nil {
		sum := 0
		if h1 != nil {
			sum += h1.Val
			h1 = h1.Next
		}
		if h2 != nil {
			sum += h2.Val
			h2 = h2.Next
		}

		sum += flag

		newNode := &datastructures.ListNode{Val: sum % 10}
		flag = sum / 10
		if newHead == nil {
			newHead = newNode
		} else {
			newLast.Next = newNode
		}
		newLast = newNode
	}
	if flag > 0 {
		newLast.Next = &datastructures.ListNode{Val: flag}
	}
	return newHead
}
