package findtheduplicatenumber

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"官方示例1", []int{1, 3, 4, 2, 2}, 2},
		{"官方示例2", []int{3, 1, 3, 4, 2}, 3},
		{"官方示例3", []int{1, 1}, 1},
		{"重复在末尾", []int{1, 2, 3, 4, 4}, 4},
		{"重复在开头", []int{2, 2, 3, 4, 1}, 2},
		{"大数组", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicate(tt.nums); got != tt.want {
				t.Errorf("FindDuplicate(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestFindDuplicateBinary(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"官方示例1", []int{1, 3, 4, 2, 2}, 2},
		{"官方示例2", []int{3, 1, 3, 4, 2}, 3},
		{"官方示例3", []int{1, 1}, 1},
		{"重复在末尾", []int{1, 2, 3, 4, 4}, 4},
		{"重复在开头", []int{2, 2, 3, 4, 1}, 2},
		{"大数组", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateBinary(tt.nums); got != tt.want {
				t.Errorf("FindDuplicateBinary(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkFindDuplicate(b *testing.B) {
	nums := make([]int, 10001)
	for i := 0; i < 10000; i++ {
		nums[i] = i + 1
	}
	nums[10000] = 5000
	for i := 0; i < b.N; i++ {
		FindDuplicate(nums)
	}
}

func BenchmarkFindDuplicateBinary(b *testing.B) {
	nums := make([]int, 10001)
	for i := 0; i < 10000; i++ {
		nums[i] = i + 1
	}
	nums[10000] = 5000
	for i := 0; i < b.N; i++ {
		FindDuplicateBinary(nums)
	}
}
