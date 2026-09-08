package mergetwosortedlists

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,4]+[1,3,4]", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"示例2: 空+空", []int{}, []int{}, []int{}},
		{"示例3: 空+[0]", []int{}, []int{0}, []int{0}},

		// 边界：一方为空
		{"空+[1,2,3]", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"[1,2,3]+空", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},

		// 边界：单节点
		{"[1]+[2]", []int{1}, []int{2}, []int{1, 2}},
		{"[2]+[1]", []int{2}, []int{1}, []int{1, 2}},
		{"[1]+[1]", []int{1}, []int{1}, []int{1, 1}},

		// 边界：完全相等
		{"[1,2,3]+[1,2,3]", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},

		// 边界：负数
		{"[-5,-1]+[-3,0]", []int{-5, -1}, []int{-3, 0}, []int{-5, -3, -1, 0}},

		// 边界：一方全小于另一方
		{"[1,2]+[3,4,5]", []int{1, 2}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"[4,5,6]+[1,2]", []int{4, 5, 6}, []int{1, 2}, []int{1, 2, 4, 5, 6}},

		// 边界：长度差较大
		{"[1]+[1,2,3,4,5,6,7,8,9,10]", []int{1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := datastructures.NewLinkedList(tt.list1)
			list2 := datastructures.NewLinkedList(tt.list2)
			got := MergeTwoLists(list1, list2)
			var gotSlice []int
			if got != nil {
				gotSlice = got.ToSlice()
			}
			want := tt.want
			if len(want) == 0 {
				want = nil // 空链表归一化为 nil，与 gotSlice 对齐
			}
			if !reflect.DeepEqual(gotSlice, want) {
				t.Errorf("MergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, gotSlice, tt.want)
			}
		})
	}
}

// TestMergeTwoListsRec 验证递归法与迭代法结果一致
func TestMergeTwoListsRec(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{"示例1: [1,2,4]+[1,3,4]", []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{"示例2: 空+空", []int{}, []int{}, []int{}},
		{"示例3: 空+[0]", []int{}, []int{0}, []int{0}},
		{"负数混合", []int{-5, -1}, []int{-3, 0}, []int{-5, -3, -1, 0}},
		{"一方全小于另一方", []int{4, 5, 6}, []int{1, 2}, []int{1, 2, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := datastructures.NewLinkedList(tt.list1)
			list2 := datastructures.NewLinkedList(tt.list2)
			got := MergeTwoListsRec(list1, list2)
			var gotSlice []int
			if got != nil {
				gotSlice = got.ToSlice()
			}
			want := tt.want
			if len(want) == 0 {
				want = nil // 空链表归一化为 nil，与 gotSlice 对齐
			}
			if !reflect.DeepEqual(gotSlice, want) {
				t.Errorf("MergeTwoListsRec(%v, %v) = %v, want %v", tt.list1, tt.list2, gotSlice, tt.want)
			}
		})
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	benchmarks := []struct {
		name  string
		list1 []int
		list2 []int
	}{
		{"len=10", makeSortedRange(10), makeSortedRange(10)},
		{"len=50", makeSortedRange(50), makeSortedRange(50)},
		{"len=500", makeSortedRange(500), makeSortedRange(500)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				l1 := datastructures.NewLinkedList(bm.list1)
				l2 := datastructures.NewLinkedList(bm.list2)
				MergeTwoLists(l1, l2)
			}
		})
	}
}

func BenchmarkMergeTwoListsRec(b *testing.B) {
	list1 := makeSortedRange(500)
	list2 := makeSortedRange(500)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l1 := datastructures.NewLinkedList(list1)
		l2 := datastructures.NewLinkedList(list2)
		MergeTwoListsRec(l1, l2)
	}
}

// makeSortedRange 生成 [0, 2, 4, ..., 2(n-1)] 的升序切片
func makeSortedRange(n int) []int {
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i * 2
	}
	return vals
}
