package binarysearchtreeiterator

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// newBST 从层序切片构建二叉搜索树，MinInt32 表示 nil
func newBST(vals []int) *datastructures.TreeNode {
	return datastructures.NewTreeFromSlice(vals)
}

func TestBSTIteratorOfficialExample(t *testing.T) {
	// LeetCode 官方示例：[7,3,15,null,null,9,20] 的中序遍历为 [3,7,9,15,20]
	it := Constructor(newBST([]int{7, 3, 15, math.MinInt32, math.MinInt32, 9, 20}))

	wantNext := []int{3, 7, 9, 15, 20}
	// 每次 Next 之前 HasNext 都应为 true
	wantHas := []bool{true, true, true, true, true}

	for i, want := range wantNext {
		if got := it.HasNext(); got != wantHas[i] {
			t.Fatalf("第 %d 次 HasNext() = %v, want %v", i, got, wantHas[i])
		}
		if got := it.Next(); got != want {
			t.Fatalf("第 %d 次 Next() = %v, want %v", i, got, want)
		}
	}
	if it.HasNext() {
		t.Fatal("遍历结束后 HasNext() 应为 false")
	}
}

func TestBSTIterator(t *testing.T) {
	tests := []struct {
		name string
		vals []int // 层序表示，MinInt32 为 nil
		want []int // 期望的中序遍历结果
	}{
		// LeetCode 官方示例
		{"示例1: [7,3,15,null,null,9,20]", []int{7, 3, 15, math.MinInt32, math.MinInt32, 9, 20}, []int{3, 7, 9, 15, 20}},
		{"示例2: [5,3,7]", []int{5, 3, 7}, []int{3, 5, 7}},

		// 边界：单节点
		{"单节点", []int{1}, []int{1}},

		// 边界：空树
		{"空树", nil, nil},

		// 边界：只有左子树（退化为链表 5→4→3）
		{"只有左子树", []int{5, 4, math.MinInt32, 3}, []int{3, 4, 5}},

		// 边界：只有右子树（退化为链表 1→2→3）
		{"只有右子树", []int{1, math.MinInt32, 2, math.MinInt32, 3}, []int{1, 2, 3}},

		// 边界：0 值与极值
		{"含0和较大值", []int{10, 0, 1000000}, []int{0, 10, 1000000}},

		// 复杂树：13 是 14 的左孩子
		{"多层BST", []int{8, 3, 10, 1, 6, math.MinInt32, 14, math.MinInt32, math.MinInt32, 4, 7, 13},
			[]int{1, 3, 4, 6, 7, 8, 10, 13, 14}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := Constructor(newBST(tt.vals))
			var got []int
			for it.HasNext() {
				got = append(got, it.Next())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("中序遍历结果 = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIteratorInterleaved(t *testing.T) {
	// 交错调用 next / hasNext，模拟官方示例的调用序列
	it := Constructor(newBST([]int{7, 3, 15, math.MinInt32, math.MinInt32, 9, 20}))
	type call struct {
		op   string // "next" 或 "hasNext"
		want int    // next 的期望值；hasNext 用 1/0 表示 true/false
	}
	calls := []call{
		{"next", 3}, {"next", 7}, {"hasNext", 1}, {"next", 9},
		{"hasNext", 1}, {"next", 15}, {"hasNext", 1}, {"next", 20}, {"hasNext", 0},
	}
	for i, c := range calls {
		if c.op == "next" {
			if got := it.Next(); got != c.want {
				t.Fatalf("第 %d 步 next = %d, want %d", i, got, c.want)
			}
		} else {
			if got := it.HasNext(); got != (c.want == 1) {
				t.Fatalf("第 %d 步 hasNext = %v, want %v", i, got, c.want == 1)
			}
		}
	}
}

func BenchmarkBSTIterator(b *testing.B) {
	// 构建一棵满二叉树，高度分别为 10 和 15
	benchmarks := []struct {
		name  string
		root  *datastructures.TreeNode
		nodes int
	}{
		{"满树高度10", newBST(buildFullTreeVals(10)), (1 << 10) - 1},
		{"满树高度15", newBST(buildFullTreeVals(15)), (1 << 15) - 1},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := Constructor(bm.root)
				sum := 0
				for it.HasNext() {
					sum += it.Next()
				}
				_ = sum
			}
		})
	}
}

// buildFullTreeVals 生成高度为 h 的满二叉树的层序切片
func buildFullTreeVals(h int) []int {
	n := (1 << h) - 1
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i + 1
	}
	return vals
}
