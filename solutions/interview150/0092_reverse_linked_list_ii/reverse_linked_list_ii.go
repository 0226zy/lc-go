package reverselinkedlistii

import "github.com/0226zy/lc-go/pkg/datastructures"

// ReverseBetween 反转链表 II
// 反转单链表中从位置 left 到 right 的节点，返回反转后的链表。
// 使用虚拟头节点 + 穿针引线法：prev 定位到反转段前驱，
// 然后对反转段做头插法，逐个把后续节点提到 prev 后面，
// 反转段的尾节点天然保持与原后继的连接，无需特判。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ReverseBetween(head *datastructures.ListNode, left, right int) *datastructures.ListNode {
	dummy := &datastructures.ListNode{Next: head} // 虚拟头节点，统一处理 left=1
	prev := dummy
	// prev 走到反转段的前驱（第 left-1 个节点）
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	// segmentHead 是反转段的头，反转完成后它将变成反转段的尾
	segmentHead := prev.Next
	// 头插 right-left 次：把 segmentHead 后面的节点逐个插到 prev 后面
	for i := 0; i < right-left; i++ {
		moved := segmentHead.Next     // 待移动节点
		segmentHead.Next = moved.Next // 从原位置摘下，不断链
		moved.Next = prev.Next        // 插到 prev 后面
		prev.Next = moved
	}
	return dummy.Next
}
