package rotatelist

import "github.com/0226zy/lc-go/pkg/datastructures"

// RotateRight 旋转链表
// 给你一个链表的头节点 head，将链表每个节点向右移动 k 个位置。
// 时间复杂度: O(n) 一次遍历求长度 + 常数次指针移动  空间复杂度: O(1)
func RotateRight(head *datastructures.ListNode, k int) *datastructures.ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	// 第一次遍历：求链表长度 n，并找到尾节点
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	// k 可能远大于链表长度，先取模
	k %= n
	if k == 0 {
		return head
	}
	// 把链表首尾相接成环
	tail.Next = head
	// 向右移动 k 位后，新的尾节点是倒数第 k+1 个节点，
	// 即从头走 n-k-1 步
	steps := n - k - 1
	newTail := head
	for i := 0; i < steps; i++ {
		newTail = newTail.Next
	}
	// 新头节点是 newTail 的下一个节点，断环返回
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}
