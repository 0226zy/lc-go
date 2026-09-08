package partitionlist

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		x    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 按3分隔", []int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{"示例2: 头节点不小于x", []int{2, 1}, 2, []int{1, 2}},

		// 边界：空链表
		{"空链表", []int{}, 3, nil},
		// 边界：单节点
		{"单节点小于x", []int{1}, 3, []int{1}},
		{"单节点大于等于x", []int{5}, 3, []int{5}},
		// 边界：全部小于 x，顺序不变
		{"全部小于x", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		// 边界：全部大于等于 x，顺序不变
		{"全部大于等于x", []int{5, 6, 7}, 3, []int{5, 6, 7}},
		// 边界：x 大于所有值 / 小于所有值
		{"x大于所有值", []int{3, 1, 2}, 10, []int{3, 1, 2}},
		{"x小于所有值", []int{3, 1, 2}, 1, []int{3, 1, 2}},
		// 边界：含相等值
		{"含与x相等的值", []int{1, 4, 3, 0, 2, 5, 2}, 3, []int{1, 0, 2, 2, 4, 3, 5}},
		// 边界：负数
		{"包含负数", []int{-1, 4, -2, 3, 0}, 0, []int{-1, -2, 4, 3, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := Partition(head, tt.x).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition(%v, %d) = %v, want %v", tt.vals, tt.x, got, tt.want)
			}
		})
	}
}

func BenchmarkPartition(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		b.Run("len="+itoa(n), func(b *testing.B) {
			vals := make([]int, n)
			for i := range vals {
				vals[i] = (i * 37) % 100
			}
			for i := 0; i < b.N; i++ {
				Partition(datastructures.NewLinkedList(vals), 50)
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
