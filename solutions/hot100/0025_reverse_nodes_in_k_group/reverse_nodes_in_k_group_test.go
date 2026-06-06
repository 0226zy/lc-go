package reversenodesinkgroup

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestReverseNodesInKGroupIterative(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		// 示例 1：k = 2
		{"示例1_k等于2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},

		// 示例 2：k = 3
		{"示例2_k等于3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},

		// 边界：k = 1（不翻转）
		{"k等于1_不翻转", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},

		// 边界：k 等于链表长度（全部翻转）
		{"k等于长度_全部翻转", []int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},

		// 边界：k 大于链表长度（不翻转）
		{"k大于长度_不翻转", []int{1, 2, 3}, 5, []int{1, 2, 3}},

		// 边界：不足 k 个节点（不翻转最后一组）
		{"不足k个_保留最后", []int{1, 2, 3, 4}, 3, []int{3, 2, 1, 4}},

		// 边界：空链表
		{"空链表", nil, 2, []int{}},

		// 边界：单个节点
		{"单个节点", []int{1}, 2, []int{1}},

		// 边界：两个节点，k=2
		{"两个节点_k等于2", []int{1, 2}, 2, []int{2, 1}},

		// 正常：偶数个节点，k=2
		{"偶数节点_k等于2", []int{1, 2, 3, 4, 5, 6}, 2, []int{2, 1, 4, 3, 6, 5}},

		// 正常：奇数个节点，k=3
		{"奇数节点_k等于3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{3, 2, 1, 6, 5, 4, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if tt.nums != nil {
				head = datastructures.NewLinkedList(tt.nums)
			}

			result := ReverseNodesInKGroupIterative(head, tt.k)
			resultSlice := []int{}
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expected) {
				t.Errorf("ReverseNodesInKGroupIterative(%v, %d) = %v, expected %v",
					tt.nums, tt.k, resultSlice, tt.expected)
			}
		})
	}
}

func TestReverseNodesInKGroupRecursive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		// 示例 1：k = 2
		{"示例1_k等于2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},

		// 示例 2：k = 3
		{"示例2_k等于3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},

		// 边界：k = 1（不翻转）
		{"k等于1_不翻转", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},

		// 边界：k 等于链表长度（全部翻转）
		{"k等于长度_全部翻转", []int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},

		// 边界：k 大于链表长度（不翻转）
		{"k大于长度_不翻转", []int{1, 2, 3}, 5, []int{1, 2, 3}},

		// 边界：不足 k 个节点（不翻转最后一组）
		{"不足k个_保留最后", []int{1, 2, 3, 4}, 3, []int{3, 2, 1, 4}},

		// 边界：空链表
		{"空链表", nil, 2, []int{}},

		// 边界：单个节点
		{"单个节点", []int{1}, 2, []int{1}},

		// 边界：两个节点，k=2
		{"两个节点_k等于2", []int{1, 2}, 2, []int{2, 1}},

		// 正常：偶数个节点，k=2
		{"偶数节点_k等于2", []int{1, 2, 3, 4, 5, 6}, 2, []int{2, 1, 4, 3, 6, 5}},

		// 正常：奇数个节点，k=3
		{"奇数节点_k等于3", []int{1, 2, 3, 4, 5, 6, 7}, 3, []int{3, 2, 1, 6, 5, 4, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if tt.nums != nil {
				head = datastructures.NewLinkedList(tt.nums)
			}

			result := ReverseNodesInKGroupRecursive(head, tt.k)
			resultSlice := []int{}
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expected) {
				t.Errorf("ReverseNodesInKGroupRecursive(%v, %d) = %v, expected %v",
					tt.nums, tt.k, resultSlice, tt.expected)
			}
		})
	}
}

// BenchmarkReverseNodesInKGroupIterative 迭代法性能测试
func BenchmarkReverseNodesInKGroupIterative(b *testing.B) {
	// 准备测试数据：1000 个节点的链表
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		head := datastructures.NewLinkedList(nums)
		result := ReverseNodesInKGroupIterative(head, 2)
		_ = result // 避免未使用变量警告
	}
}

// BenchmarkReverseNodesInKGroupRecursive 递归法性能测试
func BenchmarkReverseNodesInKGroupRecursive(b *testing.B) {
	// 准备测试数据：1000 个节点的链表
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		head := datastructures.NewLinkedList(nums)
		result := ReverseNodesInKGroupRecursive(head, 2)
		_ = result // 避免未使用变量警告
	}
}
