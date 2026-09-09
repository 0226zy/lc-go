package triangle

import (
	"reflect"
	"testing"
)

var minimumTotalCases = []struct {
	name     string
	triangle [][]int
	want     int
}{
	// LeetCode 官方示例
	{name: "示例1：四层三角形", triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, want: 11},
	{name: "示例2：单元素负数", triangle: [][]int{{-10}}, want: -10},

	// 边界：单元素
	{name: "单元素零", triangle: [][]int{{0}}, want: 0},
	{name: "单元素正数", triangle: [][]int{{7}}, want: 7},

	// 边界：两行
	{name: "两行选较小子节点", triangle: [][]int{{1}, {2, 3}}, want: 3},

	// 全负数：最小路径和是最“负”的那条路
	{name: "全负数三角形", triangle: [][]int{{-1}, {-2, -3}, {-4, -5, -6}}, want: -10},

	// 全正数：贪心地走数值小的分支
	{name: "全正数走左边界", triangle: [][]int{{1}, {1, 9}, {1, 9, 9}}, want: 3},
	{name: "全正数走右边界", triangle: [][]int{{1}, {9, 1}, {9, 9, 1}}, want: 3},

	// 中间凹陷：最小值藏在中间
	{name: "最小值在中间", triangle: [][]int{{2}, {9, 1}, {9, 9, 1}}, want: 4},

	// 含零与正负混合
	{name: "正负混合", triangle: [][]int{{-1}, {2, 3}, {1, -1, -3}}, want: -1},

	// 更深的三角形，含负数混合
	{name: "五层含负数混合", triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}, {-1, 5, 2, 0, 9}}, want: 13},
}

func TestMinimumTotal(t *testing.T) {
	for _, tt := range minimumTotalCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotal(tt.triangle); got != tt.want {
				t.Errorf("MinimumTotal(%v) = %d, want %d", tt.triangle, got, tt.want)
			}
		})
	}
}

func TestMinimumTotalOptimized(t *testing.T) {
	for _, tt := range minimumTotalCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTotalOptimized(tt.triangle); got != tt.want {
				t.Errorf("MinimumTotalOptimized(%v) = %d, want %d", tt.triangle, got, tt.want)
			}
		})
	}
}

// cloneTriangle 深拷贝三角形，用于构造期望的快照
func cloneTriangle(t [][]int) [][]int {
	c := make([][]int, len(t))
	for i := range t {
		c[i] = make([]int, len(t[i]))
		copy(c[i], t[i])
	}
	return c
}

// TestMinimumTotalNotModifyInput 验证两个版本实现都不会修改输入的 triangle
func TestMinimumTotalNotModifyInput(t *testing.T) {
	for _, fn := range []struct {
		name string
		f    func([][]int) int
	}{
		{"DP数组版", MinimumTotal},
		{"滚动数组版", MinimumTotalOptimized},
	} {
		t.Run(fn.name, func(t *testing.T) {
			triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
			snapshot := cloneTriangle(triangle)

			if got := fn.f(triangle); got != 11 {
				t.Errorf("got %d, want %d", got, 11)
			}
			if !reflect.DeepEqual(triangle, snapshot) {
				t.Errorf("输入被修改了: got %v, want %v", triangle, snapshot)
			}
		})
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
	versions := []struct {
		name string
		f    func([][]int) int
	}{
		{"DP数组版", MinimumTotal},
		{"滚动数组版", MinimumTotalOptimized},
	}

	// 200 行（题目约束上限），元素全为 -10000
	// 每行都取 -10000，最小路径和 = 200 * -10000 = -2000000
	t.Run("200行全为-10000", func(t *testing.T) {
		for _, v := range versions {
			tri := buildTriangle(200, func(i, j int) int { return -10000 })
			if got := v.f(tri); got != -2000000 {
				t.Errorf("%s = %d, want %d", v.name, got, -2000000)
			}
		}
	})

	// 200 行只有左边界为 0、其余为 1：可以全程走 0，最小路径和 = 0
	t.Run("200行锯齿交替存在全零路径", func(t *testing.T) {
		for _, v := range versions {
			tri := buildTriangle(200, func(i, j int) int {
				if j == 0 {
					return 0
				}
				return 1
			})
			if got := v.f(tri); got != 0 {
				t.Errorf("%s = %d, want %d", v.name, got, 0)
			}
		}
	})

	// 200 行伪随机确定值（可复现），两个版本互相印证
	t.Run("200行伪随机值两版本结果一致", func(t *testing.T) {
		tri := buildTriangle(200, func(i, j int) int { return (i*7+j*13)%20001 - 10000 })
		gotDP := MinimumTotal(cloneTriangle(tri))
		gotOpt := MinimumTotalOptimized(cloneTriangle(tri))
		if gotDP != gotOpt {
			t.Errorf("两版本结果不一致: DP数组版 = %d, 滚动数组版 = %d", gotDP, gotOpt)
		}
	})
}

var minimumTotalBenchmarks = []struct {
	name string
	tri  [][]int
}{
	{"官方示例4层", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
	{"50行伪随机", buildTriangle(50, func(i, j int) int { return (i*7 + j*13) % 20001 - 10000 })},
	{"200行伪随机(约束上限)", buildTriangle(200, func(i, j int) int { return (i*7 + j*13) % 20001 - 10000 })},
}

func BenchmarkMinimumTotal(b *testing.B) {
	for _, bm := range minimumTotalBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinimumTotal(bm.tri)
			}
		})
	}
}

func BenchmarkMinimumTotalOptimized(b *testing.B) {
	for _, bm := range minimumTotalBenchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinimumTotalOptimized(bm.tri)
			}
		})
	}
}
