package parallelcourses

import "testing"

func TestMinimumSemesters(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		relations [][]int
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: 两门先修课指向同一门", 3, [][]int{{1, 3}, {2, 3}}, 2},
		{"示例2: 三课程循环依赖", 3, [][]int{{1, 2}, {2, 3}, {3, 1}}, -1},

		// 边界：单门课程、无先修关系
		{"单门课程无先修", 1, [][]int{}, 1},
		{"多门课程互不依赖", 4, [][]int{}, 1},

		// 链式依赖：学期数等于链长
		{"单链条依赖", 3, [][]int{{1, 2}, {2, 3}}, 3},

		// 菱形依赖：1->2, 1->3, 2->4, 3->4，需 3 学期
		{"菱形依赖", 4, [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}}, 3},

		// 含环：部分课程可修但无法全部修完
		{"部分成环", 4, [][]int{{1, 2}, {2, 3}, {3, 2}, {3, 4}}, -1},

		// 存在孤立课程：可与其余课程第一学期同修
		{"含孤立课程", 4, [][]int{{1, 2}, {2, 3}}, 3},

		// 自环被约束排除，但双节点环等价
		{"双节点互相依赖", 2, [][]int{{1, 2}, {2, 1}}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSemesters(tt.n, tt.relations); got != tt.want {
				t.Errorf("MinimumSemesters(%d, %v) = %d, want %d", tt.n, tt.relations, got, tt.want)
			}
		})
	}
}

func BenchmarkMinimumSemesters(b *testing.B) {
	// 构造 5000 门课程的长链：1->2->...->5000，是分层最深的极端情况
	n := 5000
	relations := make([][]int, 0, n-1)
	for i := 1; i < n; i++ {
		relations = append(relations, []int{i, i + 1})
	}

	b.Run("长链5000门课", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MinimumSemesters(n, relations)
		}
	})

	// 构造宽层依赖：全部指向最后一门课，2 学期即可完成
	wide := make([][]int, 0, n-1)
	for i := 1; i < n; i++ {
		wide = append(wide, []int{i, n})
	}
	b.Run("宽层5000门课", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MinimumSemesters(n, wide)
		}
	})
}
