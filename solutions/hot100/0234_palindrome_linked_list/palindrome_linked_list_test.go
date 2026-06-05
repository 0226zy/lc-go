package palindromelinkedlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect bool
	}{
		// LeetCode 官方示例
		{"示例1_回文链表", []int{1, 2, 2, 1}, true},
		{"示例2_非回文链表", []int{1, 2}, false},

		// 边界情况
		{"单个节点_是回文", []int{1}, true},
		{"两个节点_回文", []int{1, 1}, true},
		{"两个节点_非回文", []int{1, 2}, false},

		// 奇数长度
		{"奇数长度_回文", []int{1, 2, 1}, true},
		{"奇数长度_非回文", []int{1, 2, 3}, false},

		// 偶数长度
		{"偶数长度_回文", []int{1, 2, 2, 1}, true},
		{"偶数长度_非回文", []int{1, 2, 3, 4}, false},

		// 特殊值
		{"包含零_回文", []int{0, 1, 0}, true},
		{"包含负数_回文", []int{-1, 2, -1}, true},
		{"全部相同值_回文", []int{1, 1, 1, 1}, true},

		// 长链表
		{"长链表_回文", []int{1, 2, 3, 4, 5, 4, 3, 2, 1}, true},
		{"长链表_非回文", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.input)
			result := IsPalindrome(head)

			if result != tt.expect {
				t.Errorf("IsPalindrome(%v) = %v, want %v", tt.input, result, tt.expect)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// 准备测试数据
	testCases := []struct {
		name   string
		input  []int
	}{
		{"短链表_5个节点", []int{1, 2, 3, 2, 1}},
		{"中等链表_50个节点", makePalindromeList(50)},
		{"长链表_500个节点", makePalindromeList(500)},
		{"超长链表_5000个节点", makePalindromeList(5000)},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := datastructures.NewLinkedList(tc.input)
				IsPalindrome(head)
			}
		})
	}
}

// makePalindromeList 生成回文链表的值数组
func makePalindromeList(n int) []int {
	result := make([]int, n)
	for i := 0; i < n/2; i++ {
		result[i] = i + 1
		result[n-1-i] = i + 1
	}
	if n%2 == 1 {
		result[n/2] = n/2 + 1
	}
	return result
}
