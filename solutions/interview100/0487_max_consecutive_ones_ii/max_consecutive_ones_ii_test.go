package maxconsecutiveonesii

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 翻转中间0", []int{1, 0, 1, 1, 0}, 4},
		{"示例2: 两个0不可兼得", []int{1, 0, 1, 1, 0, 1}, 4},

		// 边界：单元素
		{"单个1", []int{1}, 1},
		{"单个0翻转为1", []int{0}, 1},

		// 边界：全 0 / 全 1
		{"全0只能翻转一个", []int{0, 0, 0}, 1},
		{"全1无需翻转", []int{1, 1, 1, 1}, 4},

		// 典型场景
		{"0在开头", []int{0, 1, 1, 1}, 4},
		{"0在结尾", []int{1, 1, 1, 0}, 4},
		{"间隔的两个0", []int{1, 0, 1, 0, 1}, 3},
		{"长串1被0分隔", []int{1, 1, 0, 1, 1, 1, 0, 1, 1}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.nums); got != tt.want {
				t.Errorf("FindMaxConsecutiveOnes(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkFindMaxConsecutiveOnes(b *testing.B) {
	// 构造 10 万元素的混合数组：连续若干个 1 后跟一个 0，循环
	long := make([]int, 0, 100000)
	for len(long) < 100000 {
		long = append(long, 1, 1, 1, 1, 1, 1, 1, 0)
	}
	long = long[:100000]

	benchmarks := []struct {
		name string
		nums []int
	}{
		{"len=5", []int{1, 0, 1, 1, 0}},
		{"len=1000", long[:1000]},
		{"len=100000", long},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMaxConsecutiveOnes(bm.nums)
			}
		})
	}
}
