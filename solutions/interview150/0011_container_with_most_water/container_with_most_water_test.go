package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: 经典例子[1,8,6,2,5,4,8,3,7]", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"示例2: 两条线[1,1]", []int{1, 1}, 1},

		// 边界：最小规模 n = 2
		{"两条等高的线[5,5]", []int{5, 5}, 5},
		{"两条线一高一矮[3,9]", []int{3, 9}, 3},

		// 边界：含 0 高度
		{"一侧高度为0", []int{0, 9, 0}, 0},
		{"中间全为0", []int{4, 0, 0, 4}, 12},

		// 边界：单调递增与单调递减
		{"单调递增[1,2,3,4,5]", []int{1, 2, 3, 4, 5}, 6},
		{"单调递减[5,4,3,2,1]", []int{5, 4, 3, 2, 1}, 6},

		// 边界：全部等高
		{"全部等高[7,7,7,7]", []int{7, 7, 7, 7}, 21},

		// 边界：最优解在两端
		{"最高两条线在两端[10,1,1,1,10]", []int{10, 1, 1, 1, 10}, 40},
		// 边界：最优解在中间
		{"最高两条线在中间[1,10,1,10,1]", []int{1, 10, 1, 10, 1}, 20},

		// 边界：一条线极高但相邻线矮，宽度不足以弥补
		{"一侧极高另一侧矮[1,8,100,2]", []int{1, 8, 100, 2}, 8},

		// 边界：最大值约束 height[i] = 10^4
		{"高度为最大值10000", []int{10000, 0, 0, 10000}, 30000},

		// 易错：短板决定面积，次优高度组合可能更大
		{"宽度弥补高度不足[2,3,4,5,18,17,6]", []int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

// TestMaxAreaLarge 压力测试：n = 10^5 的极限规模
func TestMaxAreaLarge(t *testing.T) {
	n := 100000
	height := make([]int, n)
	// 两端最高（10000），中间为 0，最优解必为两端，面积为 10000 * (n-1)
	height[0], height[n-1] = 10000, 10000
	want := 10000 * (n - 1)
	if got := MaxArea(height); got != want {
		t.Errorf("MaxArea(大数组两端最高) = %d, want %d", got, want)
	}

	// 全部高度为 10000，最优解仍为两端
	for i := range height {
		height[i] = 10000
	}
	if got := MaxArea(height); got != want {
		t.Errorf("MaxArea(大数组全部等高) = %d, want %d", got, want)
	}
}

func BenchmarkMaxArea(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"小规模n=10", 10},
		{"中规模n=1000", 1000},
		{"大规模n=100000", 100000},
	}

	for _, bm := range benchmarks {
		// 构造锯齿形数组，避免全部命中同一分支
		height := make([]int, bm.n)
		for i := range height {
			height[i] = (i*7919 + 13) % 10001
		}
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				MaxArea(height)
			}
		})
	}
}
