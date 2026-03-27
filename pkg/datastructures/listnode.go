package datastructures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
