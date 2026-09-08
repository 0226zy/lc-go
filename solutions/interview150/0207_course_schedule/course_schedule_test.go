package courseschedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		// LeetCode 官方示例
		{"示例1: 可以完成", 2, [][]int{{1, 0}}, true},
		{"示例2: 存在环", 2, [][]int{{1, 0}, {0, 1}}, false},

		// 边界：无先修要求
		{"无先修要求", 3, nil, true},
		{"单门课程", 1, nil, true},

		// 边界：链式依赖
		{"链式依赖", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},

		// 边界：复杂无环
		{
			"多分支无环",
			6,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3}, {5, 4}},
			true,
		},

		// 边界：自环
		{"课程依赖自身", 1, [][]int{{0, 0}}, false},

		// 边界：三个节点的环
		{
			"三角环",
			3,
			[][]int{{1, 0}, {2, 1}, {0, 2}},
			false,
		},

		// 边界：部分有环
		{
			"部分课程成环",
			4,
			[][]int{{1, 0}, {2, 1}, {0, 2}, {3, 0}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("CanFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}

func TestCanFinishDFS(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{"示例1: 可以完成", 2, [][]int{{1, 0}}, true},
		{"示例2: 存在环", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"无先修要求", 3, nil, true},
		{"链式依赖", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{"自环", 1, [][]int{{0, 0}}, false},
		{
			"三角环",
			3,
			[][]int{{1, 0}, {2, 1}, {0, 2}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinishDFS(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("CanFinishDFS(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}

// makeChain 生成 n 门课程的链式先修关系：0<-1<-2<-...<-n-1
func makeChain(n int) [][]int {
	pre := make([][]int, 0, n-1)
	for i := 1; i < n; i++ {
		pre = append(pre, []int{i, i - 1})
	}
	return pre
}

// makeCycle 生成 n 门课程的环形先修关系（无解）
func makeCycle(n int) [][]int {
	pre := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		pre = append(pre, []int{i, (i + 1) % n})
	}
	return pre
}

func BenchmarkCanFinish(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		pre  [][]int
	}{
		{"100门链式", 100, makeChain(100)},
		{"1000门链式", 1000, makeChain(1000)},
		{"1000门环", 1000, makeCycle(1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanFinish(bm.n, bm.pre)
			}
		})
	}
}
