package sortlist

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [4,2,1,3]", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{"示例2: [-1,5,3,4,0]", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},

		// 边界：空链表
		{"空链表", nil, nil},

		// 边界：单节点
		{"单节点", []int{1}, []int{1}},
		{"单节点负数", []int{-5}, []int{-5}},

		// 边界：已有序（正序/逆序）
		{"已有序", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"完全逆序", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},

		// 边界：全部相同
		{"全部相同", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},

		// 边界：含负数
		{"含负数", []int{-1, -5, 3, -2, 0}, []int{-5, -2, -1, 0, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := SortList(head).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortList(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func TestSortListBottomUp(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1: [4,2,1,3]", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{"示例2: [-1,5,3,4,0]", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{"空链表", nil, nil},
		{"单节点", []int{1}, []int{1}},
		{"完全逆序", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"全部相同", []int{3, 3, 3, 3}, []int{3, 3, 3, 3}},
		{"含负数", []int{-1, -5, 3, -2, 0}, []int{-5, -2, -1, 0, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := SortListBottomUp(head).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortListBottomUp(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkSortList(b *testing.B) {
	benchmarks := []struct {
		name string
		vals []int
	}{
		{"len=100", generateRandomList(100)},
		{"len=1000", generateRandomList(1000)},
		{"len=10000", generateRandomList(10000)},
		{"len=100000", generateRandomList(100000)},
	}

	for _, bm := range benchmarks {
		b.Run("递归归并/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := datastructures.NewLinkedList(bm.vals)
				SortList(head)
			}
		})
		b.Run("自底向上/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := datastructures.NewLinkedList(bm.vals)
				SortListBottomUp(head)
			}
		})
	}
}

// generateRandomList 用简单线性同余法生成伪随机链表数组，避免测试依赖随机源
func generateRandomList(n int) []int {
	vals := make([]int, n)
	seed := 1
	for i := 0; i < n; i++ {
		seed = (seed*1103515245 + 12345) % (1 << 31)
		vals[i] = seed % 100000
	}
	return vals
}
