package binarytreelevelordertraversal

import (
	"math"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestLevelOrder_OfficialExample(t *testing.T) {
	// 树: [3,9,20,null,null,15,7] -> 层序: [[3],[9,20],[15,7]]
	root := datastructures.NewTreeFromSlice([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7})
	want := [][]int{{3}, {9, 20}, {15, 7}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_EmptyTree(t *testing.T) {
	want := [][]int{}

	got := LevelOrder(nil)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_SingleNode(t *testing.T) {
	root := datastructures.NewTreeFromSlice([]int{42})
	want := [][]int{{42}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_LeftSkewed(t *testing.T) {
	// 左斜树: [3,2,null,1]
	root := datastructures.NewTreeFromSlice([]int{3, 2, math.MinInt32, 1})
	want := [][]int{{3}, {2}, {1}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_RightSkewed(t *testing.T) {
	// 右斜树: [1,null,2,null,3]
	root := datastructures.NewTreeFromSlice([]int{1, math.MinInt32, 2, math.MinInt32, 3})
	want := [][]int{{1}, {2}, {3}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_CompleteTree(t *testing.T) {
	// 完全二叉树: [1,2,3,4,5,6,7]
	root := datastructures.NewTreeFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	want := [][]int{{1}, {2, 3}, {4, 5, 6, 7}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLevelOrder_OnlyLeftChildren(t *testing.T) {
	// 只有左子节点: [1,2,3,4]（4 是 2 的左子）
	root := datastructures.NewTreeFromSlice([]int{1, 2, 3, 4})
	want := [][]int{{1}, {2, 3}, {4}}

	got := LevelOrder(root)
	if !utils.Equal2DIntSlice(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func BenchmarkLevelOrder_CompleteTree(b *testing.B) {
	root := datastructures.NewTreeFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = LevelOrder(root)
	}
}
