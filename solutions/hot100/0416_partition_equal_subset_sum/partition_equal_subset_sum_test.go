package partitionequalsubsetsum

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"官方示例1", []int{1, 5, 11, 5}, true},
		{"官方示例2", []int{1, 2, 3, 5}, false},
		{"单元素", []int{1}, false},
		{"两元素相等", []int{1, 1}, true},
		{"两元素不等", []int{1, 2}, false},
		{"奇数和", []int{1, 2, 5}, false},
		{"全相同偶数个", []int{2, 2, 2, 2}, true},
		{"全相同奇数个", []int{2, 2, 2}, false},
		{"最大元素超一半", []int{100, 1, 1, 1, 1}, false},
		{"最大元素等于一半", []int{5, 1, 1, 1, 1, 1}, true},
		{"单元素为1", []int{1}, false},
		{"连续整数", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, true},  // sum=55? 奇数 false
		{"能平分", []int{1, 2, 3, 4, 5, 6, 7}, true},              // sum=28, target=14, [7,6,1]或[7,4,3]等
		{"大数测试", []int{99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false}, // 99+49=148, target=74, 99>74 剪枝 false
	}
	// 修正连续整数用例：1+2+...+10=55 奇数，应为 false
	tests[11] = struct {
		name string
		nums []int
		want bool
	}{"连续1到10奇数和", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPartition(tt.nums); got != tt.want {
				t.Errorf("CanPartition(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCanPartition2D(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"官方示例1", []int{1, 5, 11, 5}, true},
		{"官方示例2", []int{1, 2, 3, 5}, false},
		{"单元素", []int{1}, false},
		{"两元素相等", []int{1, 1}, true},
		{"奇数和", []int{1, 2, 5}, false},
		{"全相同偶数个", []int{2, 2, 2, 2}, true},
		{"最大元素超一半", []int{100, 1, 1, 1, 1}, false},
		{"最大元素等于一半", []int{5, 1, 1, 1, 1, 1}, true},
		{"能平分", []int{1, 2, 3, 4, 5, 6, 7}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPartition2D(tt.nums); got != tt.want {
				t.Errorf("CanPartition2D(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestCanPartitionBitset(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"官方示例1", []int{1, 5, 11, 5}, true},
		{"官方示例2", []int{1, 2, 3, 5}, false},
		{"单元素", []int{1}, false},
		{"两元素相等", []int{1, 1}, true},
		{"奇数和", []int{1, 2, 5}, false},
		{"全相同偶数个", []int{2, 2, 2, 2}, true},
		{"最大元素超一半", []int{100, 1, 1, 1, 1}, false},
		{"最大元素等于一半", []int{5, 1, 1, 1, 1, 1}, true},
		{"能平分", []int{1, 2, 3, 4, 5, 6, 7}, true},
		{"大数测试", []int{99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false}, // 99>target 剪枝
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanPartitionBitset(tt.nums); got != tt.want {
				t.Errorf("CanPartitionBitset(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// TestConsistency 验证三种实现结果一致
func TestConsistency(t *testing.T) {
	cases := [][]int{
		{1, 5, 11, 5},
		{1, 2, 3, 5},
		{1, 1},
		{1, 2, 5},
		{2, 2, 2, 2},
		{100, 1, 1, 1, 1},
		{5, 1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5, 6, 7},
		{3, 3, 3, 4, 5},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
	}
	for _, nums := range cases {
		a := CanPartition(nums)
		b := CanPartition2D(nums)
		c := CanPartitionBitset(nums)
		if a != b || b != c {
			t.Errorf("三种实现结果不一致: nums=%v, 1D=%v, 2D=%v, Bitset=%v", nums, a, b, c)
		}
	}
}

// TestRandomConsistency 随机测试三种实现一致性
func TestRandomConsistency(t *testing.T) {
	state := int64(12345)
	rand := func(max int) int {
		state = (state*1103515245 + 12345) & 0x7fffffff
		return int(state)%max + 1
	}
	for iter := 0; iter < 500; iter++ {
		n := rand(20)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rand(50)
		}
		a := CanPartition(nums)
		b := CanPartition2D(nums)
		c := CanPartitionBitset(nums)
		if a != b || b != c {
			t.Errorf("随机测试不一致: nums=%v, 1D=%v, 2D=%v, Bitset=%v", nums, a, b, c)
		}
	}
}

func BenchmarkCanPartition(b *testing.B) {
	nums := make([]int, 200)
	for i := range nums {
		nums[i] = (i*3 + 1) % 100
		if nums[i] == 0 {
			nums[i] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanPartition(nums)
	}
}

func BenchmarkCanPartitionBitset(b *testing.B) {
	nums := make([]int, 200)
	for i := range nums {
		nums[i] = (i*3 + 1) % 100
		if nums[i] == 0 {
			nums[i] = 1
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanPartitionBitset(nums)
	}
}
