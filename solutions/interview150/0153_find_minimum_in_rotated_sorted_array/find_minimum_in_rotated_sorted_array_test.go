package findminimuminrotatedsortedarray

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 旋转3次", []int{3, 4, 5, 1, 2}, 1},
		{"示例2: 最小值为0", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"示例3: 未旋转", []int{11, 13, 15, 17}, 11},

		// 边界：单元素
		{"单元素", []int{1}, 1},
		{"单元素负数", []int{-5}, -5},

		// 边界：两个元素
		{"两元素旋转", []int{2, 1}, 1},
		{"两元素未旋转", []int{1, 2}, 1},

		// 边界：最小值在头部（未旋转）
		{"最小值在头部", []int{1, 2, 3, 4, 5}, 1},

		// 边界：最小值在尾部（旋转 1 次）
		{"最小值在尾部", []int{2, 3, 4, 5, 1}, 1},

		// 边界：最小值为负数
		{"最小值为负数", []int{3, 4, 5, -1, 1, 2}, -1},

		// 边界：含极值 -5000 / 5000
		{"含最小约束值", []int{-5000, 0, 1, 2}, -5000},
		{"含最大约束值", []int{5000, -5000, 0}, -5000},

		// 边界：旋转点紧邻 mid 的各种形态
		{"旋转点在中间", []int{6, 7, 0, 1, 2, 4, 5}, 0},
		{"旋转点靠左", []int{4, 1, 2, 3}, 1},
		{"旋转点靠右", []int{2, 3, 4, 1}, 1},

		// 较大数组：旋转 2500 次
		{"长数组大旋转", makeRotated(5000, 2500), 0},
		{"长数组未旋转", makeRotated(5000, 0), 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMin(tt.nums); got != tt.want {
				t.Errorf("FindMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// makeRotated 构造长度为 n 的数组 [0,1,...,n-1] 旋转 k 次的结果
func makeRotated(n, k int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = (i + k) % n
	}
	return nums
}

func BenchmarkFindMin(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"n=10_旋转5", makeRotated(10, 5)},
		{"n=100_旋转33", makeRotated(100, 33)},
		{"n=1000_旋转500", makeRotated(1000, 500)},
		{"n=5000_旋转2500", makeRotated(5000, 2500)},
		{"n=5000_未旋转", makeRotated(5000, 0)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMin(bm.nums)
			}
		})
	}
}
