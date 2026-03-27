package addtwonumbers

import (
	"testing"

	ds "github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		// 基本用例
		{"342+465=807", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"0+0=0", []int{0}, []int{0}, []int{0}},
		{"9999999+9999=10009998", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},

		// 边界：单节点
		{"5+5=10（单节点产生进位）", []int{5}, []int{5}, []int{0, 1}},
		{"1+0=1", []int{1}, []int{0}, []int{1}},
		{"0+1=1", []int{0}, []int{1}, []int{1}},
		{"9+1=10", []int{9}, []int{1}, []int{0, 1}},

		// 边界：长度差异大
		{"l1远长于l2", []int{1, 2, 3, 4, 5}, []int{1}, []int{2, 2, 3, 4, 5}},
		{"l2远长于l1", []int{1}, []int{1, 2, 3, 4, 5}, []int{2, 2, 3, 4, 5}},

		// 边界：连续进位传播
		{"999+1=1000", []int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},
		{"1+999=1000", []int{1}, []int{9, 9, 9}, []int{0, 0, 0, 1}},
		{"全9+全9", []int{9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 1}},

		// 边界：无进位
		{"无进位相加", []int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},

		// 边界：结果中间有0
		{"中间位为0", []int{5, 5, 5}, []int{5, 5, 5}, []int{0, 1, 1, 1}},

		// 边界：一方全0
		{"0+12345=12345", []int{0}, []int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := ds.NewLinkedList(tt.l1)
			l2 := ds.NewLinkedList(tt.l2)
			got := AddTwoNumbers(l1, l2).ToSlice()
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	genList := func(n int, val int) []int {
		s := make([]int, n)
		for i := range s {
			s[i] = val
		}
		return s
	}

	benchmarks := []struct {
		name string
		size int
	}{
		{"len=10", 10},
		{"len=100", 100},
		{"len=1000", 1000},
		{"len=10000", 10000},
	}
	for _, bm := range benchmarks {
		vals := genList(bm.size, 9)
		l1 := ds.NewLinkedList(vals)
		l2 := ds.NewLinkedList(vals)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				AddTwoNumbers(l1, l2)
			}
		})
	}
}

func BenchmarkAddTwoNumbers_Asymmetric(b *testing.B) {
	short := ds.NewLinkedList([]int{1})
	longVals := make([]int, 10000)
	for i := range longVals {
		longVals[i] = 9
	}
	long := ds.NewLinkedList(longVals)

	b.Run("1+10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AddTwoNumbers(short, long)
		}
	})
}
