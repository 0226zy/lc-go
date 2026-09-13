package maxpointsontaline

import "testing"

func TestMaxPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		// LeetCode 官方示例
		{
			name:   "示例1：三点一线",
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			want:   3,
		},
		{
			name:   "示例2：最多四点",
			points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			want:   4,
		},

		// 边界：点数不足
		{name: "单点", points: [][]int{{0, 0}}, want: 1},
		{name: "两点", points: [][]int{{0, 0}, {1, 1}}, want: 2},

		// 垂直线 / 水平线
		{name: "垂直线", points: [][]int{{1, 1}, {1, 2}, {1, 3}, {2, 2}}, want: 3},
		{name: "水平线", points: [][]int{{1, 1}, {2, 1}, {3, 1}, {2, 2}}, want: 3},

		// 重合点
		{name: "全部重合", points: [][]int{{0, 0}, {0, 0}, {0, 0}}, want: 3},
		{name: "部分重合共线", points: [][]int{{1, 1}, {1, 1}, {2, 2}, {3, 3}}, want: 4},

		// 斜率约分：不同刻度同一斜率
		{name: "斜率约分", points: [][]int{{0, 0}, {1, 1}, {2, 2}, {3, 6}}, want: 3},

		// 负斜率 / 符号归一化
		{name: "负斜率", points: [][]int{{0, 0}, {1, -1}, {2, -2}, {0, 1}}, want: 3},
		{name: "反向向量同一直线", points: [][]int{{0, 0}, {1, 1}, {-1, -1}, {2, 3}}, want: 3},

		// 无三点共线
		{name: "无三点共线", points: [][]int{{0, 0}, {1, 1}, {0, 1}, {1, 0}}, want: 2},

		// 原点与象限混合
		{name: "多象限共线", points: [][]int{{-1, -1}, {0, 0}, {1, 1}, {2, 3}}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxPoints(tt.points); got != tt.want {
				t.Errorf("MaxPoints(%v) = %d, want %d", tt.points, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxPoints(b *testing.B) {
	// 构造 n 个点：一半在 y=x 上，一半散落
	n := 300
	points := make([][]int, n)
	for i := 0; i < n/2; i++ {
		points[i] = []int{i, i}
	}
	for i := n / 2; i < n; i++ {
		points[i] = []int{i, i * 2}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxPoints(points)
	}
}
