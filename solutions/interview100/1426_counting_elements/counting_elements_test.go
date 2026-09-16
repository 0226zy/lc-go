package countingelements

import "testing"

func TestCountElements(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, 2},
		{"示例2: [1,1,3,3,5,5,7,7]", []int{1, 1, 3, 3, 5, 5, 7, 7}, 0},
		{"示例3: 重复元素单独计数", []int{1, 1, 2, 2}, 2},
		{"示例4: 乱序混合", []int{1, 3, 2, 3, 5, 0}, 3},

		// 边界：空数组与单元素
		{"空数组", []int{}, 0},
		{"单元素", []int{5}, 0},

		// 边界：极端值
		{"最大值1000无后继", []int{999, 1000}, 1},
		{"最小值0有后继", []int{0, 1}, 1},

		// 边界：全部相同元素
		{"全部相同", []int{7, 7, 7}, 0},

		// 典型场景：连续序列每个非末尾元素都满足
		{"连续递增序列", []int{1, 2, 3, 4, 5}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountElements(tt.arr); got != tt.want {
				t.Errorf("CountElements(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func BenchmarkCountElements(b *testing.B) {
	// 构造长度为 1000 的数组
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = i % 500
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountElements(arr)
	}
}
