package reversenodesinkgroup

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5] k=2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"示例2: [1,2,3,4,5] k=3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},

		// 边界：k=1，每个节点自成一组，相当于不翻转
		{"k=1 不翻转", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},

		// 边界：k 等于链表长度，整体反转
		{"k=5 整体反转", []int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},

		// 边界：长度不是 k 的倍数，尾段保持原样
		{"[1,2,3,4,5] k=4 尾段剩1个", []int{1, 2, 3, 4, 5}, 4, []int{4, 3, 2, 1, 5}},
		{"[1,2,3,4,5] k=6 不足一组", []int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},

		// 边界：单节点
		{"单节点 k=1", []int{9}, 1, []int{9}},

		// 边界：两节点
		{"两节点 k=2", []int{1, 2}, 2, []int{2, 1}},

		// 边界：两节点 k=1
		{"两节点 k=1", []int{1, 2}, 1, []int{1, 2}},

		// 边界：恰好分完多组
		{"[1,2,3,4] k=2", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"[1,2,3,4,5,6] k=3", []int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},

		// 边界：含重复值
		{"含重复值 [1,1,2,2] k=2", []int{1, 1, 2, 2}, 2, []int{1, 1, 2, 2}},

		// 边界：含 0
		{"含0 [0,1,0,1] k=2", []int{0, 1, 0, 1}, 2, []int{1, 0, 1, 0}},

		// 边界：较长链表
		{"[1..10] k=4 尾段剩2个", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := ReverseKGroup(head, tt.k)
			var gotSlice []int
			if got != nil {
				gotSlice = got.ToSlice()
			}
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("ReverseKGroup(%v, %d) = %v, want %v", tt.vals, tt.k, gotSlice, tt.want)
			}
		})
	}
}

func BenchmarkReverseKGroup(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"n=100 k=2", 100, 2},
		{"n=1000 k=2", 1000, 2},
		{"n=1000 k=10", 1000, 10},
		{"n=5000 k=100", 5000, 100}, // 达到提示中的最大规模
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			vals := make([]int, bm.n)
			for i := 0; i < bm.n; i++ {
				vals[i] = i + 1
			}
			for i := 0; i < b.N; i++ {
				head := datastructures.NewLinkedList(vals)
				ReverseKGroup(head, bm.k)
			}
		})
	}
}
