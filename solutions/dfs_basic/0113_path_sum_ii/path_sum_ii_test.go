package pathsumii

import (
	"math"
	"reflect"
	"sort"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// normalize 对二维切片排序，用于忽略路径结果的顺序差异
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, path := range result {
		p := make([]int, len(path))
		copy(p, path)
		out[i] = p
	}
	sort.Slice(out, func(i, j int) bool {
		if len(out[i]) != len(out[j]) {
			return len(out[i]) < len(out[j])
		}
		for k := range out[i] {
			if out[i][k] != out[j][k] {
				return out[i][k] < out[j][k]
			}
		}
		return false
	})
	return out
}

func TestPathSum(t *testing.T) {
	tests := []struct {
		name      string
		vals      []int
		targetSum int
		want      [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 存在两条路径和为22",
			[]int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, 5, 1}, 22,
			[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}},
		{"示例2: 不存在路径和为5", []int{1, 2, 3}, 5, nil},
		{"示例3: 空树", []int{}, 5, nil},

		// 边界：单节点，值恰好匹配 / 不匹配
		{"单节点值匹配", []int{1}, 1, [][]int{{1}}},
		{"单节点值不匹配", []int{1}, 2, nil},
		// 边界：路径必须在叶节点结束，不能在内部节点提前终止
		{"必须在叶节点结束", []int{1, 2}, 1, nil},
		// 边界：负数，目标和可以为负
		{"负数路径和", []int{-2, math.MinInt32, -3}, -5, [][]int{{-2, -3}}},
		// 边界：正负抵消，路径和为 0
		{"正负抵消", []int{1, -1, math.MinInt32, math.MinInt32, 1}, 1, [][]int{{1, -1, 1}}},
		// 边界：多条路径共享前缀
		{"多条路径共享前缀", []int{1, 2, 3}, 3, [][]int{{1, 2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			got := normalize(PathSum(root, tt.targetSum))
			want := normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("PathSum(%v, %d) = %v, want %v", tt.vals, tt.targetSum, got, want)
			}
		})
	}
}

func BenchmarkPathSum(b *testing.B) {
	// 构建一个值全为 1 的满二叉树（约 4095 个节点），
	// 目标和设为树高 12，每条根到叶路径都满足条件，迫使收集所有路径
	var vals []int
	for size := 1; len(vals)+size <= 4095; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, 1)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root := datastructures.NewTreeFromSlice(vals)
		PathSum(root, 12)
	}
}
