package movingaveragefromdatastream

import "testing"

// eps 浮点数比较的误差上限
const eps = 1e-5

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		name string
		size int
		vals []int
		want []float64
	}{
		// LeetCode 官方示例
		{"示例: 窗口3连续加入", 3, []int{1, 10, 3, 5}, []float64{1.0, 5.5, 14.0 / 3.0, 6.0}},

		// 边界：窗口大小为 1，平均值永远是新值本身
		{"窗口大小为1", 1, []int{5, 7, -3}, []float64{5.0, 7.0, -3.0}},

		// 边界：窗口从未填满，按实际个数取平均
		{"窗口未填满", 5, []int{2, 4}, []float64{2.0, 3.0}},

		// 边界：负数与零
		{"包含负数", 2, []int{-10, 10, -4}, []float64{-10.0, 0.0, 3.0}},
		{"包含零", 3, []int{0, 0, 0, 6}, []float64{0.0, 0.0, 0.0, 2.0}},

		// 典型场景：窗口滑动后旧值被丢弃
		{"窗口多次滑动", 2, []int{1, 2, 3, 4, 5}, []float64{1.0, 1.5, 2.5, 3.5, 4.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma := Constructor(tt.size)
			for i, v := range tt.vals {
				got := ma.Next(v)
				if diff := got - tt.want[i]; diff > eps || diff < -eps {
					t.Errorf("第 %d 次 next(%d) = %f, want %f", i, v, got, tt.want[i])
				}
			}
		})
	}
}

func BenchmarkMovingAverageNext(b *testing.B) {
	ma := Constructor(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ma.Next(i)
	}
}
