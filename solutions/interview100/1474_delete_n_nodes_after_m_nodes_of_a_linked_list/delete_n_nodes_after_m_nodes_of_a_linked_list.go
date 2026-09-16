package deletennodesaftermnodesofalinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// DeleteNodes 删除链表 M 个节点之后的 N 个节点
// 遍历链表，保留 m 个节点后删除接下来的 n 个节点，如此反复直到链表末尾，返回修改后的头节点。
// 时间复杂度: O(len)，len 为链表长度  空间复杂度: O(1)
func DeleteNodes(head *datastructures.ListNode, m, n int) *datastructures.ListNode {
	curr := head
	for curr != nil {
		// 跳过要保留的 m 个节点（curr 已是第 1 个），结束时 curr 为保留段尾节点
		for i := 1; i < m && curr != nil; i++ {
			curr = curr.Next
		}
		if curr == nil {
			break // 链表剩余节点不足 m 个，全部保留
		}
		// 从保留段尾节点之后跨过要删除的 n 个节点
		prev := curr
		curr = curr.Next
		for i := 0; i < n && curr != nil; i++ {
			curr = curr.Next
		}
		// 保留段直接接到删除段之后，摘除中间 n 个节点
		prev.Next = curr
	}
	return head
}
