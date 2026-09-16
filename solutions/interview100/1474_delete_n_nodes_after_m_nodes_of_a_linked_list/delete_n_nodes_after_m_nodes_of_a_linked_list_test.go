package deletennodesaftermnodesofalinkedlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestDeleteNodes(t *testing.T) {
	tests := []struct {
		name string
		head []int
		m    int
		n    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: m=2 n=3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 2, 3, []int{1, 2, 6, 7, 11, 12}},
		{"示例2: m=1 n=3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 3, []int{1, 5, 9}},
		{"示例3: m=3 n=1", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 3, 1, []int{1, 2, 3, 5, 6, 7, 9, 10, 11}},

		// 边界：删除段不足 n 个节点就到尾部
		{"尾部删除段不足n个", []int{1, 2, 3, 4}, 1, 10, []int{1}},

		// 边界：链表长度恰好为 m+n 的整数倍
		{"长度恰为m加n整数倍", []int{1, 2, 3, 4, 5, 6}, 1, 1, []int{1, 3, 5}},

		// 边界：单节点链表，保留段占满
		{"单节点链表", []int{1}, 1, 1, []int{1}},

		// 边界：剩余节点不足 m 个，全部保留
		{"剩余不足m个全保留", []int{1, 2, 3, 4, 5}, 10, 3, []int{1, 2, 3, 4, 5}},

		// 边界：m=1 n=1 交替保留删除
		{"m等于1且n等于1", []int{1, 2, 3, 4, 5}, 1, 1, []int{1, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.head)
			got := DeleteNodes(head, tt.m, tt.n)
			if !utils.EqualIntSlice(got.ToSlice(), tt.want) {
				t.Errorf("DeleteNodes(%v, %d, %d) = %v, want %v", tt.head, tt.m, tt.n, got.ToSlice(), tt.want)
			}
		})
	}
}

func BenchmarkDeleteNodes(b *testing.B) {
	// 构造长度为 10000 的链表
	vals := make([]int, 10000)
	for i := range vals {
		vals[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		head := datastructures.NewLinkedList(vals)
		b.StartTimer()
		DeleteNodes(head, 5, 3)
	}
}
