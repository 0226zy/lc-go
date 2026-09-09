package triangle

import "testing"

var minimumTotalCases = []struct {
	name     string
	triangle [][]int
	want     int
}{
	{name: "单个元素负数（示例2）", triangle: [][]int{{-10}}, want: -10},
	{name: "单个元素正数", triangle: [][]int{{5}}, want: 5},
	{
		name:     "两行",
		triangle: [][]int{{1}, {2, 3}},
		want:     3,
	},
	{
		name:     "四层三角形（示例1）",
		triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
		want:     11,
	},
	{
		name:     "全负数",
		triangle: [][]int{{-1}, {-2, -3}, {-4, -5, -6}},
		want:     -10, // -1 → -3 → -6
	},
	{
		name:     "含负数混合",
		triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}, {-1, 5, 2, 0, 9}},
		want:     13, // 2 → 3 → 5 → 1 → 2
	},
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

var benchTriangle = func() [][]int {
	// 构造 200 行的三角形用于基准测试
	tri := make([][]int, 200)
	for i := range tri {
		tri[i] = make([]int, i+1)
		for j := range tri[i] {
			tri[i][j] = (i*7+j*13)%200 - 100
		}
	}
	return tri
}()

func BenchmarkMinimumTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinimumTotal(benchTriangle)
	}
}

func BenchmarkMinimumTotalOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinimumTotalOptimized(benchTriangle)
	}
}
