package meetingrooms

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      bool
	}{
		// LeetCode 官方示例
		{
			"示例1: 存在冲突",
			[][]int{{0, 30}, {5, 10}, {15, 20}},
			false,
		},
		{
			"示例2: 互不冲突",
			[][]int{{7, 10}, {2, 4}},
			true,
		},

		// 边界：空输入与单会议
		{"空输入", [][]int{}, true},
		{"单个会议", [][]int{{1, 5}}, true},

		// 边界：首尾相接不算冲突
		{
			"首尾相接",
			[][]int{{1, 3}, {3, 5}, {5, 7}},
			true,
		},

		// 典型：一个长会议覆盖多个短会议
		{
			"长会议覆盖",
			[][]int{{1, 100}, {10, 20}, {30, 40}},
			false,
		},

		// 典型：排序后冲突才显现（输入无序）
		{
			"乱序冲突",
			[][]int{{10, 20}, {1, 5}, {4, 8}},
			false,
		},

		// 典型：多个会议全部首尾相接
		{
			"密集排期无冲突",
			[][]int{{0, 5}, {5, 10}, {10, 15}, {15, 20}},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanAttendMeetings(tt.intervals); got != tt.want {
				t.Errorf("CanAttendMeetings(%v) = %v, want %v", tt.intervals, got, tt.want)
			}
		})
	}
}

func BenchmarkCanAttendMeetings(b *testing.B) {
	// 构造 10000 个互不冲突的会议
	makeIntervals := func() [][]int {
		intervals := make([][]int, 10000)
		for i := 0; i < 10000; i++ {
			intervals[i] = []int{i * 2, i*2 + 1}
		}
		return intervals
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次拷贝一份，避免排序结果影响后续轮次
		CanAttendMeetings(makeIntervals())
	}
}
