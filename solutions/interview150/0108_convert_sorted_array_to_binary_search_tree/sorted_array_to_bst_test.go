package sortedarraytobst

import (
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int // 层序遍历结果
	}{
		// LeetCode 官方示例：[-10,-3,0,5,9] 合法答案之一为 [0,-3,9,-10,null,5]
		{"示例1: [-10,-3,0,5,9]", []int{-10, -3, 0, 5, 9}, [][]int{{0}, {-3, 9}, {-10, 5}}},
		{"示例2: [1,3]", []int{1, 3}, [][]int{{3}, {1}}},

		// 边界：空数组
		{"空数组", []int{}, nil},

		// 边界：单元素
		{"单元素", []int{42}, [][]int{{42}}},
		{"单元素负数", []int{-7}, [][]int{{-7}}},

		// 边界：两个元素（取右中点为根）
		{"两个元素", []int{1, 2}, [][]int{{2}, {1}}},

		// 边界：全相同元素
		{"全相同元素", []int{5, 5, 5, 5}, [][]int{{5}, {5, 5}, {5}}},

		// 较大数组
		{"升序长数组", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{4}, {2, 6}, {1, 3, 5, 7}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := SortedArrayToBST(tt.nums)
			got := root.LevelOrder()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST(%v) 层序遍历 = %v, want %v", tt.nums, got, tt.want)
			}
			// 结果必须满足：二叉搜索树性质（中序遍历为升序）且高度平衡
			if tt.nums != nil && !isValidBST(root) {
				t.Errorf("SortedArrayToBST(%v) 结果不是合法二叉搜索树", tt.nums)
			}
			if tt.nums != nil && !isBalanced(root) {
				t.Errorf("SortedArrayToBST(%v) 结果不是高度平衡树", tt.nums)
			}
		})
	}
}

// isValidBST 检查二叉搜索树性质：中序遍历严格升序（元素允许重复时题目无此约束，这里数组升序允许相等）
func isValidBST(root *datastructures.TreeNode) bool {
	var prev *int
	ok := true
	var inorder func(node *datastructures.TreeNode)
	inorder = func(node *datastructures.TreeNode) {
		if node == nil || !ok {
			return
		}
		inorder(node.Left)
		if prev != nil && node.Val < *prev {
			ok = false
			return
		}
		v := node.Val
		prev = &v
		inorder(node.Right)
	}
	inorder(root)
	return ok
}

// isBalanced 检查高度平衡：左右子树高度差不超过 1
func isBalanced(root *datastructures.TreeNode) bool {
	var height func(node *datastructures.TreeNode) int
	balanced := true
	height = func(node *datastructures.TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := height(node.Left), height(node.Right)
		if l-r > 1 || r-l > 1 {
			balanced = false
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	height(root)
	return balanced
}

func BenchmarkSortedArrayToBST(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=10", makeRange(10)},
		{"len=100", makeRange(100)},
		{"len=1000", makeRange(1000)},
		{"len=10000", makeRange(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SortedArrayToBST(bm.nums)
			}
		})
	}
}

// makeRange 生成 [0, n) 的升序数组
func makeRange(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	return nums
}
