package meetingroomsii

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		// LeetCode 官方示例
		{
			"示例1: 需要两间",
			[][]int{{0, 30}, {5, 10}, {15, 20}},
			2,
		},
		{
			"示例2: 只需一间",
			[][]int{{7, 10}, {2, 4}},
			1,
		},

		// 边界：单个会议
		{"单个会议", [][]int{{1, 5}}, 1},

		// 边界：全部会议同时进行
		{
			"全部重叠",
			[][]int{{1, 10}, {2, 9}, {3, 8}},
			3,
		},

		// 边界：首尾相接可复用同一间
		{
			"首尾相接",
			[][]int{{1, 3}, {3, 5}, {5, 7}},
			1,
		},

		// 典型：部分重叠
		{
			"交错重叠",
			[][]int{{1, 4}, {2, 5}, {3, 6}},
			3,
		},
		{
			"阶梯重叠",
			[][]int{{1, 4}, {2, 3}, {4, 6}, {5, 7}},
			2,
		},

		// 典型：乱序输入
		{
			"乱序输入",
			[][]int{{15, 20}, {0, 30}, {5, 10}},
			2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMeetingRooms(tt.intervals); got != tt.want {
				t.Errorf("MinMeetingRooms(%v) = %d, want %d", tt.intervals, got, tt.want)
			}
		})
	}
}

func BenchmarkMinMeetingRooms(b *testing.B) {
	// 构造 10000 个会议：一半同时段重叠，一半串行
	makeIntervals := func() [][]int {
		intervals := make([][]int, 10000)
		for i := 0; i < 5000; i++ {
			intervals[i] = []int{0, 5000 - i} // 嵌套重叠
		}
		for i := 5000; i < 10000; i++ {
			intervals[i] = []int{(i - 5000) * 2, (i-5000)*2 + 1} // 互不重叠
		}
		return intervals
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinMeetingRooms(makeIntervals())
	}
}
