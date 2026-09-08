package triangle

import (
	"reflect"
	"testing"
)

// cloneTriangle 深拷贝三角形，用于构造期望的快照
func cloneTriangle(t [][]int) [][]int {
	c := make([][]int, len(t))
	for i := range t {
		c[i] = make([]int, len(t[i]))
		copy(c[i], t[i])
	}
	return c
}

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		// LeetCode 官方示例
		{"示例1: 四层三角形", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"示例2: 单元素负数", [][]int{{-10}}, -10},

		// 边界：单元素
		{"单元素零", [][]int{{0}}, 0},
		{"单元素正数", [][]int{{7}}, 7},

		// 边界：两行
		{"两行选较小子节点", [][]int{{1}, {2, 3}}, 3},

		// 全负数：最小路径和是最“负”的那条路
		{"全负数三角形", [][]int{{-1}, {-2, -3}, {-4, -5, -6}}, -10},

		// 全正数：贪心地走数值小的分支
		{"全正数走左边界", [][]int{{1}, {1, 9}, {1, 9, 9}}, 3},
		{"全正数走右边界", [][]int{{1}, {9, 1}, {9, 9, 1}}, 3},

		// 中间凹陷：最小值藏在中间
		{"最小值在中间", [][]int{{2}, {9, 1}, {9, 9, 1}}, 4},

		// 含零与正负混合
		{"正负混合", [][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotal(tt.triangle); got != tt.want {
				t.Errorf("MinimumTotal(%v) = %d, want %d", tt.triangle, got, tt.want)
			}
		})
	}
}

// TestMinimumTotalNotModifyInput 验证实现不会修改输入的 triangle
func TestMinimumTotalNotModifyInput(t *testing.T) {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	snapshot := cloneTriangle(triangle)

	if got := MinimumTotal(triangle); got != 11 {
		t.Errorf("MinimumTotal(%v) = %d, want %d", triangle, got, 11)
	}
	if !reflect.DeepEqual(triangle, snapshot) {
		t.Errorf("输入被修改了: got %v, want %v", triangle, snapshot)
	}
}

// buildTriangle 构造 n 行三角形，元素值由 f(i,j) 决定，用于压力场景
func buildTriangle(n int, f func(i, j int) int) [][]int {
	tri := make([][]int, n)
	for i := 0; i < n; i++ {
		tri[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			tri[i][j] = f(i, j)
		}
	}
	return tri
}

func TestMinimumTotalStress(t *testing.T) {
	// 200 行（题目约束上限），元素全为 -10000
	// 每行都取 -10000，最小路径和 = 200 * -10000 = -2000000
	t.Run("200行全为-10000", func(t *testing.T) {
		tri := buildTriangle(200, func(i, j int) int { return -10000 })
		if got := MinimumTotal(tri); got != -2000000 {
			t.Errorf("MinimumTotal = %d, want %d", got, -2000000)
		}
	})

	// 200 行锯齿值 (i+j)%2 交替 0/1：可以全程走 0，最小路径和 = 0
	t.Run("200行锯齿交替存在全零路径", func(t *testing.T) {
		tri := buildTriangle(200, func(i, j int) int {
			if j == 0 {
				return 0
			}
			return 1
		})
		if got := MinimumTotal(tri); got != 0 {
			t.Errorf("MinimumTotal = %d, want %d", got, 0)
		}
	})

	// 200 行随机化确定值（伪随机但可复现），与朴素 O(n^2) 二维 DP 对照
	t.Run("200行伪随机值与二维DP对照", func(t *testing.T) {
		seed := 42
		rand := func() int {
			// 简单线性同余伪随机数，保证可复现
			seed = (seed*1103515245 + 12345) & 0x7fffffff
			return seed%20001 - 10000 // 落在 [-10000, 10000]
		}
		tri := buildTriangle(200, func(i, j int) int { return rand() })

		// 朴素二维 DP 作为参照实现
		n := len(tri)
		dp2 := make([][]int, n)
		for i := range dp2 {
			dp2[i] = make([]int, i+1)
		}
		copy(dp2[n-1], tri[n-1])
		for i := n - 2; i >= 0; i-- {
			for j := 0; j <= i; j++ {
				m := dp2[i+1][j]
				if dp2[i+1][j+1] < m {
					m = dp2[i+1][j+1]
				}
				dp2[i][j] = tri[i][j] + m
			}
		}

		if got := MinimumTotal(tri); got != dp2[0][0] {
			t.Errorf("MinimumTotal = %d, want %d (二维DP参照)", got, dp2[0][0])
		}
	})
}

func BenchmarkMinimumTotal(b *testing.B) {
	benchmarks := []struct {
		name string
		tri  [][]int
	}{
		{"官方示例4层", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
		{"50行伪随机", buildTriangle(50, func(i, j int) int { return (i*7 + j*13) % 20001 - 10000 })},
		{"200行伪随机(约束上限)", buildTriangle(200, func(i, j int) int { return (i*7 + j*13) % 20001 - 10000 })},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinimumTotal(bm.tri)
			}
		})
	}
}
