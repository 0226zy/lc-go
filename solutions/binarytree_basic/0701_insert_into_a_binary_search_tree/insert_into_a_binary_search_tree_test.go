package insertintobst

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestInsertIntoBST(t *testing.T) {
	min := math.MinInt32
	tests := []struct {
		name string
		vals []int // 层序构建二叉树，math.MinInt32 表示 nil 空节点
		val  int
		want []int // 插入后期望的层序结果（math.MinInt32 表示 nil 空节点）
	}{
		// LeetCode 官方示例 1: root = [4,2,7,1,3], val = 5
		// 插入后可为 [4,2,7,1,3,5]
		{"示例1: root=[4,2,7,1,3],val=5",
			[]int{4, 2, 7, 1, 3}, 5,
			[]int{4, 2, 7, 1, 3, 5}},

		// LeetCode 官方示例 2: root = [40,20,60,10,30,50,70], val = 25
		// 插入后可为 [40,20,60,10,30,50,70,nil,nil,25]
		{"示例2: root=[40,20,60,10,30,50,70],val=25",
			[]int{40, 20, 60, 10, 30, 50, 70}, 25,
			[]int{40, 20, 60, 10, 30, 50, 70, min, min, 25}},

		// LeetCode 官方示例 3: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
		// 与示例 1 同一棵树，插入后 [4,2,7,1,3,5]
		{"示例3: root=[4,2,7,1,3],val=5",
			[]int{4, 2, 7, 1, 3, min, min, min, min, min, min}, 5,
			[]int{4, 2, 7, 1, 3, 5}},

		// 边界：空树，插入后成为单节点树
		{"空树插入成为根节点", nil, 1, []int{1}},

		// 边界：单节点，新值小于根，插入到左子树
		{"单节点插入左子树", []int{5}, 3, []int{5, 3}},

		// 边界：单节点，新值大于根，插入到右子树
		{"单节点插入右子树", []int{5}, 8, []int{5, min, 8}},

		// 边界：链式树（全部走右），新值插到最深处
		{"右链树最深插入", []int{1, min, 2, min, 3}, 4, []int{1, min, 2, min, 3, min, 4}},

		// 边界：链式树（全部走左），新值插到最深处
		{"左链树最深插入", []int{3, 2, min, 1}, 0, []int{3, 2, min, 1, min, 0}},

		// 边界：含负数，插入负值（-4 < -3，挂到 -3 的左孩子）
		{"含负数树插入负值", []int{0, -3, 9, min, -1}, -4, []int{0, -3, 9, -4, -1}},

		// 边界：新值比所有节点都大，沿最右侧一路插入
		{"插入最大值", []int{4, 2, 7}, 100, []int{4, 2, 7, min, min, min, 100}},

		// 边界：新值比所有节点都小，沿最左侧一路插入
		{"插入最小值", []int{4, 2, 7}, -100, []int{4, 2, 7, -100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := InsertIntoBST(root, tt.val)
			wantTree := datastructures.NewTreeFromSlice(tt.want)
			gotLevel, wantLevel := got.LevelOrder(), wantTree.LevelOrder()
			if !reflect.DeepEqual(gotLevel, wantLevel) {
				t.Errorf("InsertIntoBST(vals=%v, val=%d) 层序 = %v, 期望 %v",
					tt.vals, tt.val, gotLevel, wantLevel)
			}
		})
	}
}

func BenchmarkInsertIntoBST(b *testing.B) {
	// 构造一棵近似平衡的二叉搜索树，重复插入最深处的新值
	benchmarks := []struct {
		name string
		vals []int
		val  int
	}{
		{"平衡树深度4", []int{8, 4, 12, 2, 6, 10, 14}, 16},
		{"右链树深度4", []int{1, math.MinInt32, 2, math.MinInt32, 3, math.MinInt32, 4}, 5},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				root := datastructures.NewTreeFromSlice(bm.vals)
				InsertIntoBST(root, bm.val)
			}
		})
	}
}
