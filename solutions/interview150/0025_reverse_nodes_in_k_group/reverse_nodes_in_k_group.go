package reversenodesinkgroup

import "github.com/0226zy/lc-go/pkg/datastructures"

// ReverseKGroup K 个一组翻转链表
// 每 k 个节点一组进行翻转，不足 k 个的尾段保持原有顺序，返回修改后的链表。
// 外层逐组推进：先检查剩余节点是否够 k 个，够则做标准迭代反转并接线
// （前驱接新头、组尾接下一组的头），不够则结束；内层复用单链表反转模板。
// 时间复杂度: O(n)  空间复杂度: O(1)
func ReverseKGroup(head *datastructures.ListNode, k int) *datastructures.ListNode {
	dummy := &datastructures.ListNode{Next: head}
	groupPrev := dummy // 上一组的尾节点，即本组的前驱
	for {
		// 第 1 步：检查剩余节点是否够 k 个
		groupHead := groupPrev.Next
		check := groupHead
		for i := 0; i < k; i++ {
			if check == nil {
				return dummy.Next // 不足 k 个，尾段保持原样
			}
			check = check.Next
		}

		// 第 2 步：标准迭代反转 k 个节点
		// 结束后 pre 指向本组新头，cur 指向下一组的头
		var pre *datastructures.ListNode
		cur := groupHead
		for i := 0; i < k; i++ {
			nextNode := cur.Next
			cur.Next = pre
			pre = cur
			cur = nextNode
		}

		// 第 3 步：接线
		groupPrev.Next = pre // 前驱接本组反转后的新头
		groupHead.Next = cur // 本组新尾（原头）接下一组的头

		// 第 4 步：推进到下一组
		groupPrev = groupHead
	}
}
