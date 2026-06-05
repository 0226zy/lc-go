package palindromelinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsPalindrome 判断链表是否为回文链表
// 给你一个单链表的头节点 head，请你判断该链表是否为回文链表。
// 时间复杂度: O(n)  空间复杂度: O(1)
func IsPalindrome(head *datastructures.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 步骤1: 使用快慢指针找到链表中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 步骤2: 反转后半部分链表
	secondHalf := reverseList(slow.Next)
	slow.Next = nil // 切断前半部分和后半部分的连接

	// 步骤3: 比较前半部分和反转后的后半部分
	firstHalf := head
	result := compareLists(firstHalf, secondHalf)

	// 步骤4: 恢复链表（可选，但推荐保持输入不变）
	slow.Next = reverseList(secondHalf)

	return result
}

// reverseList 反转链表（迭代法）
func reverseList(head *datastructures.ListNode) *datastructures.ListNode {
	var prev *datastructures.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// compareLists 比较两个链表是否相同
func compareLists(l1, l2 *datastructures.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
