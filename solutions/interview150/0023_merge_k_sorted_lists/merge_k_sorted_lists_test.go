package mergeksortedlists

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// buildLists 把 [][]int 转成链表数组
func buildLists(listVals [][]int) []*datastructures.ListNode {
	lists := make([]*datastructures.ListNode, len(listVals))
	for i, vals := range listVals {
		lists[i] = datastructures.NewLinkedList(vals)
	}
	return lists
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name     string
		listVals [][]int
		want     []int
	}{
		// LeetCode 官方示例
		{"示例1: 三个链表", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"示例2: 空数组", [][]int{}, nil},
		{"示例3: 一个空链表", [][]int{{}}, nil},

		// 边界：单链表
		{"单链表", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"单链表单元素", [][]int{{7}}, []int{7}},

		// 边界：k 个单元素链表
		{"k个单元素", [][]int{{3}, {1}, {2}}, []int{1, 2, 3}},

		// 边界：含空链表
		{"含空链表", [][]int{{1, 3}, {}, {2}}, []int{1, 2, 3}},
		{"前导空链表", [][]int{{}, {1}, {2, 4}}, []int{1, 2, 4}},

		// 边界：含负数
		{"含负数", [][]int{{-10, -5}, {-8, -3}, {-12}}, []int{-12, -10, -8, -5, -3}},

		// 边界：重复值
		{"大量重复值", [][]int{{1, 1}, {1, 1, 1}, {1}}, []int{1, 1, 1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeKLists(buildLists(tt.listVals)).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKLists(%v) = %v, want %v", tt.listVals, got, tt.want)
			}
		})
	}
}

func TestMergeKListsWithHeap(t *testing.T) {
	tests := []struct {
		name     string
		listVals [][]int
		want     []int
	}{
		{"示例1: 三个链表", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"示例2: 空数组", [][]int{}, nil},
		{"示例3: 一个空链表", [][]int{{}}, nil},
		{"单链表", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"k个单元素", [][]int{{3}, {1}, {2}}, []int{1, 2, 3}},
		{"含空链表", [][]int{{1, 3}, {}, {2}}, []int{1, 2, 3}},
		{"含负数", [][]int{{-10, -5}, {-8, -3}, {-12}}, []int{-12, -10, -8, -5, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeKListsWithHeap(buildLists(tt.listVals)).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeKListsWithHeap(%v) = %v, want %v", tt.listVals, got, tt.want)
			}
		})
	}
}

func BenchmarkMergeKLists(b *testing.B) {
	benchmarks := []struct {
		name string
		k    int // 链表条数
		n    int // 每条链表长度
	}{
		{"k=10,n=100", 10, 100},
		{"k=100,n=100", 100, 100},
		{"k=1000,n=100", 1000, 100},
		{"k=100,n=1000", 100, 1000},
	}

	for _, bm := range benchmarks {
		b.Run("分治/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				lists := generateKLists(bm.k, bm.n)
				b.StartTimer()
				MergeKLists(lists)
			}
		})
		b.Run("堆/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				lists := generateKLists(bm.k, bm.n)
				b.StartTimer()
				MergeKListsWithHeap(lists)
			}
		})
	}
}

// generateKLists 生成 k 条各含 n 个升序节点的链表（值由线性同余伪随机生成）
func generateKLists(k, n int) []*datastructures.ListNode {
	lists := make([]*datastructures.ListNode, k)
	seed := 1
	for i := 0; i < k; i++ {
		vals := make([]int, n)
		prev := 0
		for j := 0; j < n; j++ {
			seed = (seed*1103515245 + 12345) % (1 << 31)
			prev += seed % 100
			vals[j] = prev
		}
		lists[i] = datastructures.NewLinkedList(vals)
	}
	return lists
}
