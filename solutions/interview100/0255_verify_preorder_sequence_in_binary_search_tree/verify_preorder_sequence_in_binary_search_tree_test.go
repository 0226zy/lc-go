package verifypreordersequenceinbinarysearchtree

import "testing"

func TestVerifyPreorder(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		want     bool
	}{
		// LeetCode 官方示例
		{"示例1: 合法序列 [5,2,1,3,6]", []int{5, 2, 1, 3, 6}, true},
		{"示例2: 非法序列 [5,2,6,1,3]", []int{5, 2, 6, 1, 3}, false},

		// 边界：单元素与两元素
		{"单元素", []int{1}, true},
		{"两元素递减", []int{2, 1}, true},
		{"两元素递增", []int{1, 2}, true},

		// 链式结构
		{"一路左孩子", []int{5, 4, 3, 2, 1}, true},
		{"一路右孩子", []int{1, 2, 3, 4, 5}, true},

		// 典型场景
		{"右子树中出现小于祖先的值", []int{3, 4, 2}, false},
		{"多层右拐后回到更上层右子树", []int{10, 5, 1, 7, 6, 8, 15, 20}, true},
		{"大值之后又出现越界小值", []int{10, 15, 9}, false},
		{"右孩子的左孩子仍合法", []int{10, 15, 12}, true},
		{"右子树中混入小于根的值", []int{5, 8, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPreorder(tt.preorder); got != tt.want {
				t.Errorf("VerifyPreorder(%v) = %v, want %v", tt.preorder, got, tt.want)
			}
			if got := VerifyPreorderRecursive(tt.preorder); got != tt.want {
				t.Errorf("VerifyPreorderRecursive(%v) = %v, want %v", tt.preorder, got, tt.want)
			}
		})
	}
}

// buildBalancedPreorder 构造一棵平衡 BST 的前序遍历序列（取中点为根递归）
func buildBalancedPreorder(vals []int) []int {
	if len(vals) == 0 {
		return nil
	}
	mid := len(vals) / 2
	pre := []int{vals[mid]}
	pre = append(pre, buildBalancedPreorder(vals[:mid])...)
	pre = append(pre, buildBalancedPreorder(vals[mid+1:])...)
	return pre
}

func BenchmarkVerifyPreorder(b *testing.B) {
	// 构造规模为 n 的升序数组对应的平衡 BST 前序遍历
	sorted := make([]int, 10000)
	for i := range sorted {
		sorted[i] = i + 1
	}
	preorder := buildBalancedPreorder(sorted)

	b.Run("单调栈_平衡树_万级", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			VerifyPreorder(preorder)
		}
	})
	b.Run("递归分治_平衡树_千级", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			VerifyPreorderRecursive(preorder[:1000])
		}
	})
}
