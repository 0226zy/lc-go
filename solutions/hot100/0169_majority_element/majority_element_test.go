package majorityelement

import "testing"

func TestMajorityElement_OfficialExamples(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "官方示例2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMajorityElement_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "单元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "全部相同",
			nums: []int{5, 5, 5, 5, 5},
			want: 5,
		},
		{
			name: "多数在开头",
			nums: []int{1, 1, 1, 2, 3},
			want: 1,
		},
		{
			name: "多数在末尾",
			nums: []int{2, 3, 1, 1, 1},
			want: 1,
		},
		{
			name: "含负数",
			nums: []int{-1, -1, -1, 2, 3},
			want: -1,
		},
		{
			name: "交替后连续",
			nums: []int{1, 2, 1, 2, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMajorityElement_AllSolutions(t *testing.T) {
	solutions := []func([]int) int{
		MajorityElement,
		MajorityElementHashMap,
	}

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "官方示例2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "单元素",
			nums: []int{7},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fn := range solutions {
				got := fn(tt.nums)
				if got != tt.want {
					t.Errorf("solution[%d](%v) = %d, want %d", i, tt.nums, got, tt.want)
				}
			}
		})
	}
}

func BenchmarkMajorityElement_Voting(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		MajorityElement(nums)
	}
}

func BenchmarkMajorityElement_HashMap(b *testing.B) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	for i := 0; i < b.N; i++ {
		MajorityElementHashMap(nums)
	}
}
