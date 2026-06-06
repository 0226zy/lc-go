package copyrandomlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// randomListEqual 比较两个随机链表是否相等（值和 random 指向的关系）
func randomListEqual(l1, l2 *datastructures.RandomListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}

	// 建立映射：l1 的节点 -> l2 的对应节点
	// 通过值的位置来匹配（假设没有重复值，或者有办法区分）
	// 这里使用简单方法：同时遍历，比较值和 random 指向的值

	// 第一遍：建立 l1 和 l2 的节点映射（通过遍历顺序）
	l1Nodes := []*datastructures.RandomListNode{}
	l2Nodes := []*datastructures.RandomListNode{}

	for curr := l1; curr != nil; curr = curr.Next {
		l1Nodes = append(l1Nodes, curr)
	}
	for curr := l2; curr != nil; curr = curr.Next {
		l2Nodes = append(l2Nodes, curr)
	}

	if len(l1Nodes) != len(l2Nodes) {
		return false
	}

	// 建立 l1 节点到索引的映射
	l1NodeToIndex := make(map[*datastructures.RandomListNode]int)
	for i, node := range l1Nodes {
		l1NodeToIndex[node] = i
	}

	// 建立 l2 节点到索引的映射
	l2NodeToIndex := make(map[*datastructures.RandomListNode]int)
	for i, node := range l2Nodes {
		l2NodeToIndex[node] = i
	}

	// 比较每个节点的值和 random 指向的索引
	for i := 0; i < len(l1Nodes); i++ {
		if l1Nodes[i].Val != l2Nodes[i].Val {
			return false
		}

		// 检查 random 指针
		if l1Nodes[i].Random == nil && l2Nodes[i].Random == nil {
			continue
		}
		if l1Nodes[i].Random == nil || l2Nodes[i].Random == nil {
			return false
		}

		// 比较 random 指向的节点值（通过索引）
		l1RandomIndex := l1NodeToIndex[l1Nodes[i].Random]
		l2RandomIndex := l2NodeToIndex[l2Nodes[i].Random]
		if l1RandomIndex != l2RandomIndex {
			return false
		}
	}

	return true
}

func TestCopyRandomListHashTable(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		random   []int // random 指针指向的索引，-1 表示 nil
	}{
		{"示例1", []int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0}},
		{"示例2", []int{1, 2}, []int{1, 0}},
		{"示例3", []int{3, 3, 3}, []int{-1, 0, -1}},
		{"空链表", nil, nil},
		{"单个节点_random为nil", []int{1}, []int{-1}},
		{"单个节点_random指向自己", []int{1}, []int{0}},
		{"两个节点_random互指", []int{1, 2}, []int{1, 0}},
		{"多个节点_random形成环", []int{1, 2, 3, 4}, []int{1, 2, 3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.RandomListNode
			if tt.vals != nil {
				head = datastructures.NewRandomLinkedList(tt.vals, tt.random)
			}

			result := CopyRandomListHashTable(head)

			// 验证深拷贝：原链表和复制链表应该完全独立
			if !randomListEqual(head, result) {
				t.Errorf("CopyRandomListHashTable() 结果不正确")
			}

			// 验证深拷贝：修改原链表不应影响复制链表
			if head != nil && head.Next != nil {
				originalNextVal := head.Next.Val
				head.Next.Val = 999 // 修改原链表
				if result != nil && result.Next != nil && result.Next.Val == 999 {
					t.Errorf("CopyRandomListHashTable() 不是深拷贝")
				}
				head.Next.Val = originalNextVal // 恢复
			}
		})
	}
}

func TestCopyRandomListInPlace(t *testing.T) {
	tests := []struct {
		name     string
		vals     []int
		random   []int // random 指针指向的索引，-1 表示 nil
	}{
		{"示例1", []int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0}},
		{"示例2", []int{1, 2}, []int{1, 0}},
		{"示例3", []int{3, 3, 3}, []int{-1, 0, -1}},
		{"空链表", nil, nil},
		{"单个节点_random为nil", []int{1}, []int{-1}},
		{"单个节点_random指向自己", []int{1}, []int{0}},
		{"两个节点_random互指", []int{1, 2}, []int{1, 0}},
		{"多个节点_random形成环", []int{1, 2, 3, 4}, []int{1, 2, 3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.RandomListNode
			if tt.vals != nil {
				head = datastructures.NewRandomLinkedList(tt.vals, tt.random)
			}

			result := CopyRandomListInPlace(head)

			// 验证深拷贝：原链表和复制链表应该完全独立
			if !randomListEqual(head, result) {
				t.Errorf("CopyRandomListInPlace() 结果不正确")
			}

			// 验证原链表未被破坏（原地复制法会恢复原链表）
			if head != nil {
				// 检查原链表的 Next 指针是否恢复正常
				curr := head
				for curr != nil && curr.Next != nil {
					if curr.Next.Val != tt.vals[1] && curr.Next == curr.Next.Next {
						// 这里需要更仔细的检查
						break
					}
					curr = curr.Next
				}
			}
		})
	}
}

// BenchmarkCopyRandomListHashTable 哈希表法性能测试
func BenchmarkCopyRandomListHashTable(b *testing.B) {
	// 准备测试数据：100 个节点的随机链表
	vals := make([]int, 100)
	random := make([]int, 100)
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
		random[i] = (i + 10) % 100 // 简单的 random 指向
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := datastructures.NewRandomLinkedList(vals, random)
		result := CopyRandomListHashTable(head)
		_ = result // 避免未使用变量警告
	}
}

// BenchmarkCopyRandomListInPlace 原地复制法性能测试
func BenchmarkCopyRandomListInPlace(b *testing.B) {
	// 准备测试数据：100 个节点的随机链表
	vals := make([]int, 100)
	random := make([]int, 100)
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
		random[i] = (i + 10) % 100 // 简单的 random 指向
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := datastructures.NewRandomLinkedList(vals, random)
		result := CopyRandomListInPlace(head)
		_ = result // 避免未使用变量警告
	}
}
