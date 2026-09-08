package rotatelist

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 右移2位", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"示例2: 右移4位取模", []int{0, 1, 2}, 4, []int{2, 0, 1}},

		// 边界：k 为 0 或取模后为 0，不旋转
		{"k为0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"k是长度的倍数", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"k是长度的倍数且很大", []int{1, 2, 3}, 3000000, []int{1, 2, 3}},

		// 边界：空链表与单节点
		{"空链表", []int{}, 2, nil},
		{"单节点", []int{7}, 10, []int{7}},

		// 边界：右移 1 位
		{"右移1位", []int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
		// 边界：k 大于长度但取模后不整除
		{"k大于长度", []int{1, 2, 3, 4, 5}, 7, []int{4, 5, 1, 2, 3}},
		// 边界：全部相同值
		{"全部相同值", []int{6, 6, 6}, 1, []int{6, 6, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := RotateRight(head, tt.k).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateRight(%v, %d) = %v, want %v", tt.vals, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkRotateRight(b *testing.B) {
	for _, n := range []int{10, 100, 1000, 10000} {
		b.Run("len="+itoa(n), func(b *testing.B) {
			vals := make([]int, n)
			for i := range vals {
				vals[i] = i
			}
			for i := 0; i < b.N; i++ {
				RotateRight(datastructures.NewLinkedList(vals), n/3)
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
