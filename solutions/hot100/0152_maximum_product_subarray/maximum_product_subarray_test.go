package maximumproductsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"官方示例1", []int{2, 3, -2, 4}, 6},
		{"官方示例2", []int{-2, 0, -1}, 0},
		{"单元素正数", []int{5}, 5},
		{"单元素负数", []int{-3}, -3},
		{"全正数", []int{1, 2, 3, 4}, 24},
		{"全负数偶数个", []int{-2, -3, -4, -5}, 120},
		{"全负数奇数个", []int{-2, -3, -4}, 12},
		{"含零分段", []int{0, 2, 3, 0, 4, 5}, 20},
		{"负负得正", []int{-2, 3, -4}, 24},
		{"中间有零", []int{-2, 0, -3, -4}, 12},
		{"单个零", []int{0}, 0},
		{"多个零", []int{0, 0, 0}, 0},
		{"负数开头", []int{-1, 2, 3, -4}, 24},
		{"边界值 10", []int{10, 10, 10}, 1000},
		{"边界值 -10", []int{-10, -10, -10}, 100},
		{"交替正负", []int{1, -1, 2, -2, 3, -3}, 36},
		{"长数组全1", func() []int {
			nums := make([]int, 100)
			for i := range nums {
				nums[i] = 1
			}
			return nums
		}(), 1},
		{"上界规模交替", func() []int {
			nums := make([]int, 20000)
			for i := range nums {
				if i%2 == 0 {
					nums[i] = 2
				} else {
					nums[i] = -2
				}
			}
			return nums
		}(), func() int {
			// 20000 个交替 2,-2，乘积规律：2,-4,8,-16,... 绝对值 2^k
			// 20000 为偶数，整体乘积 = 2^20000，但会溢出 32 位
			// 实际由于会溢出，需看实际值。这里用小规模验证即可
			// 此用例改为更合理的期望
			return 0 // 占位，下方单独处理
		}()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "上界规模交替" {
				// 大规模交替用例单独验证：两种实现结果应一致
				a := MaxProduct(tt.nums)
				b := MaxProductPrefixSuffix(tt.nums)
				if a != b {
					t.Errorf("两种实现结果不一致: DP=%d, PrefixSuffix=%d", a, b)
				}
				return
			}
			if got := MaxProduct(tt.nums); got != tt.want {
				t.Errorf("MaxProduct(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxProductPrefixSuffix(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"官方示例1", []int{2, 3, -2, 4}, 6},
		{"官方示例2", []int{-2, 0, -1}, 0},
		{"单元素正数", []int{5}, 5},
		{"单元素负数", []int{-3}, -3},
		{"全正数", []int{1, 2, 3, 4}, 24},
		{"全负数偶数个", []int{-2, -3, -4, -5}, 120},
		{"全负数奇数个", []int{-2, -3, -4}, 12},
		{"含零分段", []int{0, 2, 3, 0, 4, 5}, 20},
		{"负负得正", []int{-2, 3, -4}, 24},
		{"中间有零", []int{-2, 0, -3, -4}, 12},
		{"单个零", []int{0}, 0},
		{"多个零", []int{0, 0, 0}, 0},
		{"负数开头", []int{-1, 2, 3, -4}, 24},
		{"边界值 10", []int{10, 10, 10}, 1000},
		{"边界值 -10", []int{-10, -10, -10}, 100},
		{"交替正负", []int{1, -1, 2, -2, 3, -3}, 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProductPrefixSuffix(tt.nums); got != tt.want {
				t.Errorf("MaxProductPrefixSuffix(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// TestConsistency 验证两种实现结果一致
func TestConsistency(t *testing.T) {
	cases := [][]int{
		{2, 3, -2, 4},
		{-2, 0, -1},
		{-2, 3, -4},
		{0, 2, 3, 0, 4, 5},
		{-1, 2, 3, -4},
		{1, -1, 2, -2, 3, -3},
		{-2, -3, -4, -5},
		{0},
		{-3},
		{2, -3, 4, -5, 6, -1, 2},
	}
	for _, nums := range cases {
		a := MaxProduct(nums)
		b := MaxProductPrefixSuffix(nums)
		if a != b {
			t.Errorf("两种实现结果不一致: nums=%v, DP=%d, PrefixSuffix=%d", nums, a, b)
		}
	}
}

// 随机测试：两种实现必须一致
func TestRandomConsistency(t *testing.T) {
	// 固定种子的伪随机，保证可复现
	// 限制长度较短（≤8）且元素绝对值小，避免乘积溢出 32 位导致截断差异
	seed := int64(42)
	cases := generateRandomCases(seed, 500, 8)
	for _, nums := range cases {
		a := MaxProduct(nums)
		b := MaxProductPrefixSuffix(nums)
		if a != b {
			t.Errorf("随机测试不一致: nums=%v, DP=%d, PrefixSuffix=%d", nums, a, b)
		}
	}
}

func generateRandomCases(seed int64, count, maxLen int) [][]int {
	// 简单的 LCG 伪随机
	state := seed
	rand := func() int {
		state = (state*1103515245 + 12345) & 0x7fffffff
		// 限制元素绝对值 ≤2，避免长数组乘积溢出 32 位
		vals := []int{-2, -1, 0, 1, 2}
		return vals[int(state)%len(vals)]
	}
	cases := make([][]int, count)
	for c := 0; c < count; c++ {
		length := int(state)%maxLen + 1
		cases[c] = make([]int, length)
		for i := 0; i < length; i++ {
			cases[c][i] = rand()
		}
	}
	return cases
}

func BenchmarkMaxProduct(b *testing.B) {
	nums := make([]int, 20000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = 2
		} else {
			nums[i] = -2
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProduct(nums)
	}
}

func BenchmarkMaxProductPrefixSuffix(b *testing.B) {
	nums := make([]int, 20000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = 2
		} else {
			nums[i] = -2
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProductPrefixSuffix(nums)
	}
}
