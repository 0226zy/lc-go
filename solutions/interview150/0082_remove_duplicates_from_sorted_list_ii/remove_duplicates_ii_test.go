package removeduplicatesii

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 中间两段重复", []int{1, 2, 3, 3, 4, 4, 5}, []int{1, 2, 5}},
		{"示例2: 开头连续重复", []int{1, 1, 1, 2, 3}, []int{2, 3}},

		// 边界：空链表
		{"空链表", []int{}, nil},
		// 边界：单节点
		{"单节点", []int{1}, []int{1}},
		// 边界：全部重复
		{"全部重复", []int{1, 1, 1, 1}, nil},
		// 边界：无重复
		{"无重复", []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		// 边界：头尾重复
		{"头尾重复", []int{1, 1, 2, 3, 3}, []int{2}},
		// 边界：只有两个且重复
		{"两个重复节点", []int{1, 1}, nil},
		// 边界：相邻多段重复
		{"多段连续重复", []int{1, 1, 2, 2, 3, 3}, nil},
		{"重复段之间有单个节点", []int{1, 2, 2, 3, 3, 4}, []int{1, 4}},
		// 边界：负数
		{"包含负数", []int{-3, -1, -1, 0, 0, 2}, []int{-3, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := DeleteDuplicates(head).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteDuplicates(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkDeleteDuplicates(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		b.Run("len="+itoa(n), func(b *testing.B) {
			// 构造一半节点成对重复的链表
			vals := make([]int, n)
			for i := 0; i < n; i += 2 {
				v := i
				vals[i] = v
				if i+1 < n {
					vals[i+1] = v
				}
			}
			for i := 0; i < b.N; i++ {
				DeleteDuplicates(datastructures.NewLinkedList(vals))
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
