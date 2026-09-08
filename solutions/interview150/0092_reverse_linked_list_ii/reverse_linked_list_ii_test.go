package reverselinkedlistii

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		vals  []int
		left  int
		right int
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,4,5] 反转[2,4]", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"示例2: [5] 反转[1,1]", []int{5}, 1, 1, []int{5}},

		// 边界：left == right，不反转
		{"left等于right", []int{1, 2, 3, 4, 5}, 3, 3, []int{1, 2, 3, 4, 5}},

		// 边界：反转整个链表（left=1, right=n）
		{"反转整个链表", []int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
		{"单节点反转自身", []int{7}, 1, 1, []int{7}},

		// 边界：反转前两个节点
		{"反转开头两个", []int{1, 2, 3, 4, 5}, 1, 2, []int{2, 1, 3, 4, 5}},

		// 边界：反转最后两个节点
		{"反转结尾两个", []int{1, 2, 3, 4, 5}, 4, 5, []int{1, 2, 3, 5, 4}},

		// 边界：只反转第一个节点（等价于不动）
		{"只反转第一个", []int{1, 2, 3}, 1, 1, []int{1, 2, 3}},

		// 边界：只反转最后一个节点（等价于不动）
		{"只反转最后一个", []int{1, 2, 3}, 3, 3, []int{1, 2, 3}},

		// 边界：两个节点整体反转
		{"两节点整体反转", []int{1, 2}, 1, 2, []int{2, 1}},

		// 边界：负数节点
		{"负数节点反转中间段", []int{-5, -4, -3, -2, -1}, 2, 4, []int{-5, -2, -3, -4, -1}},

		// 边界：较长链表反转中间大段
		{"长链表反转中间段", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 8, []int{1, 2, 8, 7, 6, 5, 4, 3, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := ReverseBetween(head, tt.left, tt.right)
			var gotSlice []int
			if got != nil {
				gotSlice = got.ToSlice()
			}
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("ReverseBetween(%v, %d, %d) = %v, want %v",
					tt.vals, tt.left, tt.right, gotSlice, tt.want)
			}
		})
	}
}

func BenchmarkReverseBetween(b *testing.B) {
	vals := make([]int, 500)
	for i := 0; i < 500; i++ {
		vals[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := datastructures.NewLinkedList(vals)
		ReverseBetween(head, 100, 400) // 反转中间 301 个节点
	}
}
