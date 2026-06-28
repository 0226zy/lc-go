package courseschedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "示例1：可以完成",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "示例2：存在环",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "边界：无先修课程",
			numCourses:    3,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "多条依赖链",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "单门课程",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "自环",
			numCourses:    2,
			prerequisites: [][]int{{1, 1}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("CanFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}

func BenchmarkCanFinish(b *testing.B) {
	prerequisites := [][]int{{1, 0}, {2, 1}, {3, 2}}
	for i := 0; i < b.N; i++ {
		CanFinish(4, prerequisites)
	}
}
