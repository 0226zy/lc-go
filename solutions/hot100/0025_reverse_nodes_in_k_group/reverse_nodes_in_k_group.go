package reversenodesinkgroup

import "github.com/0226zy/lc-go/pkg/datastructures"

// ReverseNodesInKGroupIterative K个一组翻转链表 - 迭代法
// 给你一个链表，每 k 个节点一组进行翻转，请你返回修改后的链表。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ReverseNodesInKGroupIterative(head *datastructures.ListNode, k int) *datastructures.ListNode {
	// 边界条件：k 不大于 1 时不需要翻转
	if head == nil || k <= 1 {
		return head
	}

	// 创建哑节点
	dummy := &datastructures.ListNode{Next: head}
	prevGroupEnd := dummy // 上一组翻转后的尾节点

	for {
		// 检查从 prevGroupEnd.Next 开始是否有 k 个节点
		kthNode := prevGroupEnd
		for i := 0; i < k; i++ {
			kthNode = kthNode.Next
			if kthNode == nil {
				// 不足 k 个节点，直接返回结果
				return dummy.Next
			}
		}

		// 记录这一组的第一个节点（翻转后会变成尾节点）
		groupStart := prevGroupEnd.Next
		// 记录下一组的开始节点
		nextGroupStart := kthNode.Next

		// 翻转这一组节点（从 groupStart 到 kthNode）
		prev := nextGroupStart
		cur := groupStart
		for cur != nextGroupStart {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		// 将翻转后的组连接到原链表
		prevGroupEnd.Next = kthNode // kthNode 现在是这一组的新头节点
		groupStart.Next = nextGroupStart

		// 更新 prevGroupEnd 为这一组的尾节点（即原来的 groupStart）
		prevGroupEnd = groupStart
	}
}

// ReverseNodesInKGroupRecursive K个一组翻转链表 - 递归法
// 给你一个链表，每 k 个节点一组进行翻转，请你返回修改后的链表。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 时间复杂度: O(n)  空间复杂度: O(n/k)
func ReverseNodesInKGroupRecursive(head *datastructures.ListNode, k int) *datastructures.ListNode {
	// 边界条件：k 不大于 1 时不需要翻转
	if head == nil || k <= 1 {
		return head
	}

	// 检查是否有 k 个节点
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			// 不足 k 个节点，直接返回 head（不翻转）
			return head
		}
		cur = cur.Next
	}

	// 递归翻转前 k 个节点
	// cur 现在是第 k+1 个节点
	prev := cur
	node := head
	for i := 0; i < k; i++ {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	// head 现在是这一组的尾节点，连接到后面递归翻转的结果
	head.Next = ReverseNodesInKGroupRecursive(cur, k)

	// prev 是这一组的新头节点
	return prev
}
