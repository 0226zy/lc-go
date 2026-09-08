package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: target=5", []int{1, 3, 5, 6}, 5, 2},
		{"示例2: target=2", []int{1, 3, 5, 6}, 2, 1},
		{"示例3: target=7", []int{1, 3, 5, 6}, 7, 4},

		// 边界：target 比所有元素都小，插到开头
		{"插入开头", []int{1, 3, 5, 6}, 0, 0},
		{"单元素且更小", []int{5}, 1, 0},

		// 边界：target 比所有元素都大，插到末尾
		{"插入末尾", []int{1, 3, 5, 6}, 10, 4},
		{"单元素且更大", []int{5}, 10, 1},

		// 边界：恰好命中
		{"命中首元素", []int{1, 3, 5, 6}, 1, 0},
		{"命中尾元素", []int{1, 3, 5, 6}, 6, 3},

		// 边界：单元素恰好命中
		{"单元素命中", []int{5}, 5, 0},

		// 边界：插入到中间相邻两数之间
		{"插入3和5之间", []int{1, 3, 5, 6}, 4, 2},
		{"负数目标", []int{-5, -3, -1}, -4, 1},
		{"两个元素插中间", []int{1, 3}, 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("SearchInsert(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	benchmarks := []struct {
		name   string
		nums   []int
		target int
	}{
		{"len=10", makeSortedArray(10), 7},
		{"len=100", makeSortedArray(100), 77},
		{"len=1000", makeSortedArray(1000), 777},
		{"len=10000", makeSortedArray(10000), 7777},
		{"len=10000-不存在", makeSortedArray(10000), 9999},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SearchInsert(bm.nums, bm.target)
			}
		})
	}
}

// makeSortedArray 生成长度为 n 的升序数组 [0, 2, 4, ...]
func makeSortedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 2 * i
	}
	return arr
}
