package mergetwosortedlists

import (
	"fmt"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// TestMergeTwoListsIterative 测试迭代法
func TestMergeTwoListsIterative(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		expect []int
	}{
		{"示例1_两个链表都非空", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"示例2_两个链表都为空", []int{}, []int{}, []int{}},
		{"示例3_一个链表为空", []int{}, []int{0}, []int{0}},
		{"只有一个元素_链表1", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"只有一个元素_链表2", []int{2, 3, 4}, []int{1}, []int{1, 2, 3, 4}},
		{"所有元素都小于另一个链表", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"所有元素都大于另一个链表", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{"有相同值的元素", []int{1, 2, 2, 3}, []int{2, 2, 4}, []int{1, 2, 2, 2, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head1, head2 *datastructures.ListNode
			if len(tt.list1) > 0 {
				head1 = datastructures.NewLinkedList(tt.list1)
			}
			if len(tt.list2) > 0 {
				head2 = datastructures.NewLinkedList(tt.list2)
			}

			result := MergeTwoListsIterative(head1, head2)

			// 转换为切片进行比较
			var resultSlice []int
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("MergeTwoListsIterative(%v, %v) = %v, want %v", tt.list1, tt.list2, resultSlice, tt.expect)
			}
		})
	}
}

// TestMergeTwoListsRecursive 测试递归法
func TestMergeTwoListsRecursive(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		expect []int
	}{
		{"示例1_两个链表都非空", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"示例2_两个链表都为空", []int{}, []int{}, []int{}},
		{"示例3_一个链表为空", []int{}, []int{0}, []int{0}},
		{"只有一个元素_链表1", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"只有一个元素_链表2", []int{2, 3, 4}, []int{1}, []int{1, 2, 3, 4}},
		{"所有元素都小于另一个链表", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"所有元素都大于另一个链表", []int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{"有相同值的元素", []int{1, 2, 2, 3}, []int{2, 2, 4}, []int{1, 2, 2, 2, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head1, head2 *datastructures.ListNode
			if len(tt.list1) > 0 {
				head1 = datastructures.NewLinkedList(tt.list1)
			}
			if len(tt.list2) > 0 {
				head2 = datastructures.NewLinkedList(tt.list2)
			}

			result := MergeTwoListsRecursive(head1, head2)

			// 转换为切片进行比较
			var resultSlice []int
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("MergeTwoListsRecursive(%v, %v) = %v, want %v", tt.list1, tt.list2, resultSlice, tt.expect)
			}
		})
	}
}

// BenchmarkMergeTwoListsIterative 迭代法性能测试
func BenchmarkMergeTwoListsIterative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		l1 := datastructures.NewLinkedList([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100})
		l2 := datastructures.NewLinkedList([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99})
		result := MergeTwoListsIterative(l1, l2)
		_ = result // 避免未使用变量警告
	}
}

// BenchmarkMergeTwoListsRecursive 递归法性能测试
func BenchmarkMergeTwoListsRecursive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		l1 := datastructures.NewLinkedList([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100})
		l2 := datastructures.NewLinkedList([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 41, 43, 45, 47, 49, 51, 53, 55, 57, 59, 61, 63, 65, 67, 69, 71, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99})
		result := MergeTwoListsRecursive(l1, l2)
		_ = result // 避免未使用变量警告
	}
}

// ttName 根据节点数生成测试名称
func ttName(n int) string {
	return fmt.Sprintf("链表长度_%d", n)
}
