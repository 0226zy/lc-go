package searchinrotatedsortedarray

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 目标在旋转点右侧", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"示例2: 目标不存在", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"示例3: 单元素未命中", []int{1}, 0, -1},

		// 边界：单元素命中/未命中
		{"单元素命中", []int{1}, 1, 0},
		{"单元素未命中", []int{2}, 3, -1},

		// 边界：未旋转的有序数组
		{"未旋转数组命中", []int{1, 2, 3, 4, 5}, 3, 2},
		{"未旋转数组未命中", []int{1, 2, 3, 4, 5}, 6, -1},
		{"未旋转数组找最小", []int{1, 2, 3, 4, 5}, 1, 0},
		{"未旋转数组找最大", []int{1, 2, 3, 4, 5}, 5, 4},

		// 边界：旋转点为 0（等于未旋转）
		{"旋转后等于原数组", []int{0, 1, 2}, 2, 2},

		// 边界：命中最大元素（旋转点前的最后一个元素）
		{"命中旋转点前最大值", []int{4, 5, 6, 7, 0, 1, 2}, 7, 3},
		{"命中旋转点", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},

		// 边界：目标在左半有序区
		{"目标在左半有序区", []int{6, 7, 0, 1, 2, 4, 5}, 7, 1},
		{"目标在左半首元素", []int{6, 7, 0, 1, 2, 4, 5}, 6, 0},

		// 边界：目标在右半有序区
		{"目标在右半有序区", []int{6, 7, 0, 1, 2, 4, 5}, 2, 4},
		{"目标在右半末元素", []int{6, 7, 0, 1, 2, 4, 5}, 5, 6},

		// 边界：负数与极值
		{"含负数", []int{3, -1, 1}, -1, 1},
		{"含负数未命中", []int{3, -1, 1}, 2, -1},

		// 两个元素的数组
		{"两元素命中左", []int{3, 1}, 3, 0},
		{"两元素命中右", []int{3, 1}, 1, 1},
		{"两元素未命中", []int{3, 1}, 2, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkSearch(b *testing.B) {
	// 构造一个长度为 n 的旋转排序数组：前半为大数，后半为小数
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=10", 10},
		{"n=100", 100},
		{"n=1000", 1000},
		{"n=10000", 10000},
		{"n=100000", 100000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			nums := make([]int, bm.n)
			k := bm.n / 2 // 旋转点
			for i := 0; i < bm.n; i++ {
				nums[i] = (i + k) % bm.n
			}
			target := bm.n / 3
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Search(nums, target)
			}
		})
	}
}
