package addtwonumbers

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 342+465=807", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"示例2: 0+0=0", []int{0}, []int{0}, []int{0}},
		{"示例3: 9999999+9999=10009998", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},

		// 边界：最终进位（99+1=100）
		{"最终进位 99+1", []int{9, 9}, []int{1}, []int{0, 0, 1}},

		// 边界：单节点（无进位）
		{"单节点 2+3", []int{2}, []int{3}, []int{5}},

		// 边界：单节点（有进位）
		{"单节点 5+5", []int{5}, []int{5}, []int{0, 1}},

		// 边界：两链表等长全 9（进位链贯穿始终）
		{"999+999=1998", []int{9, 9, 9}, []int{9, 9, 9}, []int{8, 9, 9, 1}},

		// 边界：一长一短，长链表多出的高位照抄
		{"1+999=1000", []int{1}, []int{9, 9, 9}, []int{0, 0, 0, 1}},
		{"999+1=1000", []int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},

		// 边界：一长一短无进位
		{"123+45=168", []int{3, 2, 1}, []int{5, 4}, []int{8, 6, 1}},

		// 边界：含 0 节点的多位数
		{"0+123=123", []int{0}, []int{3, 2, 1}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := datastructures.NewLinkedList(tt.l1)
			l2 := datastructures.NewLinkedList(tt.l2)
			got := AddTwoNumbers(l1, l2).ToSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	benchmarks := []struct {
		name string
		l1   []int
		l2   []int
	}{
		{"len=10", makeAllNines(10), makeAllNines(10)},
		{"len=100", makeAllNines(100), makeAllNines(100)},
		{"len=1000", makeAllNines(1000), makeAllNines(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l1 := datastructures.NewLinkedList(bm.l1)
				l2 := datastructures.NewLinkedList(bm.l2)
				AddTwoNumbers(l1, l2)
			}
		})
	}
}

// makeAllNines 生成 n 个 9 的切片（全进位最坏情况）
func makeAllNines(n int) []int {
	vals := make([]int, n)
	for i := range vals {
		vals[i] = 9
	}
	return vals
}
