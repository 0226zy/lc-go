package powxn

import (
	"math"
	"testing"
)

// 浮点数比较使用 epsilon，避免直接判等
const eps = 1e-5

func TestMyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		// LeetCode 官方示例
		{"示例1: 2^10", 2.00000, 10, 1024.00000},
		{"示例2: 2.1^3", 2.10000, 3, 9.26100},
		{"示例3: 2^-2", 2.00000, -2, 0.25000},

		// 边界：指数为 0 或 1
		{"任何数的0次幂为1", 5.0, 0, 1.0},
		{"0的0次幂按约定为1", 0.0, 0, 1.0},
		{"任何数的1次幂为本身", 3.14, 1, 3.14},

		// 边界：特殊底数
		{"1的任何次幂为1", 1.0, 1000, 1.0},
		{"-1的偶数次幂为1", -1.0, 1000, 1.0},
		{"-1的奇数次幂为-1", -1.0, 1001, -1.0},
		{"0的正数次幂为0", 0.0, 5, 0.0},
		{"负数的偶数次幂", -2.0, 10, 1024.0},
		{"负数的奇数次幂", -2.0, 3, -8.0},

		// 边界：负指数
		{"4^-1", 4.0, -1, 0.25},
		{"-2^-3", -2.0, -3, -0.125},
		{"0.5^2", 0.5, 2, 0.25},

		// 边界：指数取 int32 最小值，取负必须依赖 int64 才不会溢出
		{"2^-2147483648结果为0", 2.0, -2147483648, 0.0},
		{"1^-2147483648结果为1", 1.0, -2147483648, 1.0},
		{"-1^-2147483648结果为1", -1.0, -2147483648, 1.0},
		{"2^2147483647结果溢出为正无穷", 2.0, 2147483647, math.Inf(1)},

		// 常规用例
		{"3^5", 3.0, 5, 243.0},
		{"10^-3", 10.0, -3, 0.001},
		{"1.5^7", 1.5, 7, 17.0859375},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyPow(tt.x, tt.n)
			if math.Abs(got-tt.want) > eps {
				t.Errorf("MyPow(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.want)
			}
		})
	}
}

// TestMyPowStress 压力场景：大指数下迭代次数应为 O(log n)，与朴素连乘结果一致
func TestMyPowStress(t *testing.T) {
	t.Run("大指数幂结果与数学期望一致", func(t *testing.T) {
		got := MyPow(0.99999, 100000)
		want := math.Pow(0.99999, 100000) // 仅测试中用标准库做对照，实现本身不使用
		if math.Abs(got-want)/want > 1e-5 {
			t.Errorf("MyPow(0.99999, 100000) = %v, want %v", got, want)
		}
	})

	t.Run("接近1的底数大负指数", func(t *testing.T) {
		got := MyPow(1.00001, -100000)
		want := math.Pow(1.00001, -100000)
		if math.Abs(got-want)/want > 1e-5 {
			t.Errorf("MyPow(1.00001, -100000) = %v, want %v", got, want)
		}
	})
}

func BenchmarkMyPow(b *testing.B) {
	benchmarks := []struct {
		name string
		x    float64
		n    int
	}{
		{"小指数2^10", 2.0, 10},
		{"中指数1.5^1000", 1.5, 1000},
		{"大指数1.0001^2147483647", 1.0001, 2147483647},
		{"最小负指数2^-2147483648", 2.0, -2147483648},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyPow(bm.x, bm.n)
			}
		})
	}
}
