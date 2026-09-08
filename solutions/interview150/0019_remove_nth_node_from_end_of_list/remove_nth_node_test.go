package removenthnode

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 删除倒数第2个", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"示例2: 单节点删除唯一节点", []int{1}, 1, nil},
		{"示例3: 删除倒数第1个", []int{1, 2}, 1, []int{1}},

		// 边界：删除头节点（n 等于链表长度）
		{"删除头节点", []int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
		{"删除尾节点", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{"两个节点删除头节点", []int{1, 2}, 2, []int{2}},

		// 边界：重复值
		{"链表中存在重复值", []int{1, 2, 1}, 2, []int{1, 1}},
		{"全部相同值", []int{5, 5, 5, 5}, 3, []int{5, 5, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := RemoveNthFromEnd(head, tt.n).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd(%v, %d) = %v, want %v", tt.vals, tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		b.Run("len="+itoa(n), func(b *testing.B) {
			vals := make([]int, n)
			for i := range vals {
				vals[i] = i
			}
			for i := 0; i < b.N; i++ {
				RemoveNthFromEnd(datastructures.NewLinkedList(vals), n/2)
			}
		})
	}
}

// itoa 简单的整数转字符串，避免引入 fmt
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	return string(buf[i:])
}
