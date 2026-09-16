package plusonelinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// PlusOne 给单链表加一
// 链表按高位到低位存储一个非负整数，每位一个节点，将其加一并返回新链表的头节点。
// 核心思路：加一的进位只影响末尾连续的 9，找到最靠右的非 9 节点加一，
// 其后所有节点置 0；借助哨兵节点统一处理整条链表全是 9 时需要新增头节点的情况。
// 时间复杂度: O(n) 链表线性扫描  空间复杂度: O(1) 只用常数个指针
func PlusOne(head *datastructures.ListNode) *datastructures.ListNode {
	// 哨兵节点：当原链表全是 9 时，哨兵进位为 1 成为新的头节点
	dummy := &datastructures.ListNode{Val: 0, Next: head}
	// 记录最靠右的非 9 节点，初始为哨兵
	notNine := dummy
	for node := head; node != nil; node = node.Next {
		if node.Val != 9 {
			notNine = node
		}
	}
	// 最靠右的非 9 节点加一，其后的 9 全部进位为 0
	notNine.Val++
	for node := notNine.Next; node != nil; node = node.Next {
		node.Val = 0
	}
	// 哨兵被加一说明原链表全是 9，哨兵成为新头节点
	if dummy.Val > 0 {
		return dummy
	}
	return dummy.Next
}
