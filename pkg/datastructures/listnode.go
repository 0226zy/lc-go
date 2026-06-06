package datastructures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// RandomListNode 随机链表节点，包含随机指针
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

// NewLinkedList 从切片创建链表
func NewLinkedList(vals []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}

// ToSlice 将链表转为切片
func (head *ListNode) ToSlice() []int {
	var result []int
	for curr := head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}

// String 打印链表
func (head *ListNode) String() string {
	if head == nil {
		return "nil"
	}
	result := ""
	for curr := head; curr != nil; curr = curr.Next {
		if result != "" {
			result += " -> "
		}
		result += fmt.Sprintf("%d", curr.Val)
	}
	return result
}

// NewCycleLinkedList 创建带环链表，pos 为环入口索引，-1 表示无环
func NewCycleLinkedList(vals []int, pos int) *ListNode {
	head := NewLinkedList(vals)
	if pos < 0 || head == nil {
		return head
	}
	var cycleEntry, tail *ListNode
	curr := head
	for i := 0; curr != nil; i++ {
		if i == pos {
			cycleEntry = curr
		}
		if curr.Next == nil {
			tail = curr
		}
		curr = curr.Next
	}
	if tail != nil && cycleEntry != nil {
		tail.Next = cycleEntry
	}
	return head
}

// NewRandomLinkedList 从值和随机索引创建随机链表
// vals: 节点值列表
// randomIndices: 每个节点的 random 指针指向的索引，-1 表示 nil
func NewRandomLinkedList(vals []int, randomIndices []int) *RandomListNode {
	if len(vals) == 0 {
		return nil
	}

	// 创建所有节点
	nodes := make([]*RandomListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &RandomListNode{Val: v}
	}

	// 连接 Next 指针
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// 连接 Random 指针
	for i, randIdx := range randomIndices {
		if randIdx >= 0 && randIdx < len(nodes) {
			nodes[i].Random = nodes[randIdx]
		}
	}

	return nodes[0]
}

// RandomListToSlice 将随机链表的值转为切片（用于测试比较）
func RandomListToSlice(head *RandomListNode) []int {
	var result []int
	for curr := head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}

// RandomListToString 打印随机链表（值和random指向的值）
func RandomListToString(head *RandomListNode) string {
	if head == nil {
		return "nil"
	}
	result := ""
	for curr := head; curr != nil; curr = curr.Next {
		if result != "" {
			result += " -> "
		}
		randomVal := "nil"
		if curr.Random != nil {
			randomVal = fmt.Sprintf("%d", curr.Random.Val)
		}
		result += fmt.Sprintf("(%d,random=%s)", curr.Val, randomVal)
	}
	return result
}
