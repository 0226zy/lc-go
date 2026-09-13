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
		want [][]int // 期望的层序遍历结果（取下中点构造出的唯一树形）
	}{
		// LeetCode 官方示例：[-10,-3,0,5,9] 的合法答案之一为 [0,-10,5,null,-3,null,9]
		{"示例1: [-10,-3,0,5,9]", []int{-10, -3, 0, 5, 9}, [][]int{{0}, {-10, 5}, {-3, 9}}},
		{"示例2: [1,3]", []int{1, 3}, [][]int{{1}, {3}}},

		// 边界：空数组
		{"空数组", []int{}, nil},

		// 边界：单元素
		{"单元素", []int{42}, [][]int{{42}}},
		{"单元素负数", []int{-7}, [][]int{{-7}}},

		// 边界：两个元素（取下中点为根，右子树挂右孩子）
		{"两个元素", []int{1, 2}, [][]int{{1}, {2}}},

		// 含负数与零的混合数组
		{"含负数与零", []int{-5, -1, 0, 2, 6}, [][]int{{0}, {-5, 2}, {-1, 6}}},

		// 边界：全相同元素
		{"全相同元素", []int{5, 5, 5, 5}, [][]int{{5}, {5, 5}, {5}}},

		// 2^k-1 个元素：构造出满二叉树
		{"满二叉树规模", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{4}, {2, 6}, {1, 3, 5, 7}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := SortedArrayToBST(tt.nums)
			got := root.LevelOrder()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST(%v) 层序遍历 = %v, want %v", tt.nums, got, tt.want)
			}
			// 结果必须同时满足：二叉搜索树性质（中序遍历升序）且高度平衡
			if len(tt.nums) > 0 && !isValidBST(root) {
				t.Errorf("SortedArrayToBST(%v) 结果不是合法二叉搜索树", tt.nums)
			}
			if len(tt.nums) > 0 && !isBalanced(root) {
				t.Errorf("SortedArrayToBST(%v) 结果不是高度平衡树", tt.nums)
			}
		})
	}
}

// isValidBST 检查二叉搜索树性质：中序遍历非递减（数组升序，允许相等元素）
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

// isBalanced 检查高度平衡：每个节点的左右子树高度差不超过 1
func isBalanced(root *datastructures.TreeNode) bool {
	balanced := true
	var height func(node *datastructures.TreeNode) int
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
		{"长度10", makeRange(10)},
		{"长度100", makeRange(100)},
		{"长度1000", makeRange(1000)},
		{"长度10000", makeRange(10000)},
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
