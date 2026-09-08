package coursescheduleii

import "testing"

// validOrder 校验 order 是否是 numCourses 门课程的一组合法拓扑序
func validOrder(numCourses int, prerequisites [][]int, order []int) bool {
	if len(order) != numCourses {
		return false
	}
	// 每门课程恰好出现一次
	seen := make([]bool, numCourses)
	for _, c := range order {
		if c < 0 || c >= numCourses || seen[c] {
			return false
		}
		seen[c] = true
	}
	// 每门课的位置必须在其所有先修课之后
	pos := make([]int, numCourses)
	for i, c := range order {
		pos[c] = i
	}
	for _, pre := range prerequisites {
		if pos[pre[0]] <= pos[pre[1]] {
			return false
		}
	}
	return true
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		wantEmpty     bool // true 表示期望返回空数组
	}{
		// LeetCode 官方示例
		{"示例1", 2, [][]int{{1, 0}}, false},
		{"示例2", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, false},
		{"示例3: 有环返回空", 1, [][]int{{0, 0}}, true},

		// 边界：无先修要求
		{"无先修要求", 3, nil, false},

		// 边界：链式依赖
		{"链式依赖", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, false},

		// 边界：三个节点的环
		{"三角环", 3, [][]int{{1, 0}, {2, 1}, {0, 2}}, true},

		// 边界：部分课程成环
		{"部分课程成环", 4, [][]int{{1, 0}, {2, 1}, {0, 2}, {3, 0}}, true},

		// 边界：单门课程无先修
		{"单门课程", 1, nil, false},

		// 复杂无环
		{"多分支无环", 6, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}, {4, 3}, {5, 4}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindOrder(tt.numCourses, tt.prerequisites)
			if tt.wantEmpty {
				if len(got) != 0 {
					t.Errorf("FindOrder(%d, %v) = %v, want 空数组", tt.numCourses, tt.prerequisites, got)
				}
				return
			}
			if !validOrder(tt.numCourses, tt.prerequisites, got) {
				t.Errorf("FindOrder(%d, %v) = %v, 不是合法拓扑序", tt.numCourses, tt.prerequisites, got)
			}
		})
	}
}

func BenchmarkFindOrder(b *testing.B) {
	pre := make([][]int, 0, 999)
	for i := 1; i < 1000; i++ {
		pre = append(pre, []int{i, i - 1})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindOrder(1000, pre)
	}
}
