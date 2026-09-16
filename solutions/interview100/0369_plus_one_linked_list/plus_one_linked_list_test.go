package plusonelinkedlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, []int{1, 2, 4}},
		{"示例2: [0]", []int{0}, []int{1}},

		// 边界：单个 9，进位后变两位数
		{"单个9", []int{9}, []int{1, 0}},

		// 边界：整条链表全是 9，需要新增头节点
		{"全为9", []int{9, 9, 9}, []int{1, 0, 0, 0}},

		// 边界：末尾连续多个 9
		{"末尾连续9", []int{1, 9, 9, 9}, []int{2, 0, 0, 0}},
		{"中间分段9", []int{4, 5, 9, 9}, []int{4, 6, 0, 0}},

		// 典型场景：末位不是 9，只改末位
		{"末位非9", []int{1, 0, 8}, []int{1, 0, 9}},

		// 典型场景：包含 0 的多位整数
		{"含0的整数", []int{1, 0, 0, 0}, []int{1, 0, 0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewLinkedList(tt.vals)
			got := PlusOne(head).ToSlice()
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("PlusOne(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	// 构造最长规模（100 个节点）的链表
	vals := make([]int, 100)
	for i := range vals {
		vals[i] = (i % 9) + 1 // 1~9 循环，避免全 9
	}

	benchmarks := []struct {
		name string
		vals []int
	}{
		{"len=3", []int{1, 2, 3}},
		{"len=10", vals[:10]},
		{"len=100", vals},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 每轮重新构造，避免上一轮的加一影响输入
				head := datastructures.NewLinkedList(bm.vals)
				PlusOne(head)
			}
		})
	}
}
