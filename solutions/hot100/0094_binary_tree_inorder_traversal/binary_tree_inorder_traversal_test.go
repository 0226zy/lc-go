package binarytreeinordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestInorderTraversal_OfficialExample(t *testing.T) {
	// 树: [1,null,2,3] -> 中序: [1,3,2]
	root := datastructures.NewTreeFromSlice([]int{1, math.MinInt32, 2, 3})
	want := []int{1, 3, 2}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestInorderTraversal_EmptyTree(t *testing.T) {
	want := []int{}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(nil)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(nil)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(nil)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestInorderTraversal_SingleNode(t *testing.T) {
	root := datastructures.NewTreeFromSlice([]int{42})
	want := []int{42}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestInorderTraversal_LeftSkewed(t *testing.T) {
	// 左斜树: [3,2,1]（3 为根，2 为左子，1 为 2 的左子）
	root := datastructures.NewTreeFromSlice([]int{3, 2, math.MinInt32, 1})
	want := []int{1, 2, 3}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestInorderTraversal_RightSkewed(t *testing.T) {
	// 右斜树: [1,2,3]（1 为根，2 为右子，3 为 2 的右子）
	root := datastructures.NewTreeFromSlice([]int{1, math.MinInt32, 2, math.MinInt32, 3})
	want := []int{1, 2, 3}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestInorderTraversal_CompleteTree(t *testing.T) {
	// 完全二叉树: [1,2,3,4,5,6,7] -> 中序: [4,2,5,1,6,3,7]
	root := datastructures.NewTreeFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	want := []int{4, 2, 5, 1, 6, 3, 7}

	t.Run("递归法", func(t *testing.T) {
		got := InorderTraversalRecursive(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("迭代法", func(t *testing.T) {
		got := InorderTraversalIterative(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Morris", func(t *testing.T) {
		got := InorderTraversalMorris(root)
		if !utils.EqualIntSlice(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
