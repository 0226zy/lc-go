package partitionlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// Partition 分隔链表
// 给你一个链表的头节点 head 和一个特定值 x，请你对链表进行分隔，
// 使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
// 你应当保留两个分区中每个节点的初始相对位置。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 只是重排已有节点，不新建节点
func Partition(head *datastructures.ListNode, x int) *datastructures.ListNode {
	smallDummy := &datastructures.ListNode{} // 小于 x 的节点挂在这条链上
	bigDummy := &datastructures.ListNode{}   // 大于等于 x 的节点挂在这条链上
	small, big := smallDummy, bigDummy
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val < x {
			small.Next = curr
			small = small.Next
		} else {
			big.Next = curr
			big = big.Next
		}
	}
	// 切断 big 链的尾部，防止残留旧链表中的后续节点造成环
	big.Next = nil
	// 拼接两条链：小值链在前，大值链在后
	small.Next = bigDummy.Next
	return smallDummy.Next
}
