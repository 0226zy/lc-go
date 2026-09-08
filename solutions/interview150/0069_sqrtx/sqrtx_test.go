package sqrtx

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: x=4完全平方", 4, 2},
		{"示例2: x=8向下取整", 8, 2},

		// 边界：最小值
		{"x=0平方根为0", 0, 0},
		{"x=1平方根为1", 1, 1},

		// 边界：小数值
		{"x=2向下取整为1", 2, 1},
		{"x=3向下取整为1", 3, 1},
		{"x=5向下取整为2", 5, 2},

		// 完全平方数
		{"x=9完全平方", 9, 3},
		{"x=16完全平方", 16, 4},
		{"x=100完全平方", 100, 10},
		{"x=144完全平方", 144, 12},

		// 非完全平方数：验证向下取整
		{"x=10向下取整为3", 10, 3},
		{"x=15向下取整为3", 15, 3},
		{"x=17向下取整为4", 17, 4},
		{"x=99向下取整为9", 99, 9},

		// 大数场景
		{"x=10000完全平方", 10000, 100},
		{"x=1000000完全平方", 1000000, 1000},

		// 压力/边界：接近 int32 最大值
		{"x=2147395600最大完全平方", 2147395600, 46340},
		{"x=2147483647为int32最大值", 2147483647, 46340},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MySqrt(tt.x); got != tt.want {
				t.Errorf("MySqrt(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}

// 额外校验：对所有结果做数学一致性检查 r*r <= x < (r+1)*(r+1)
func TestMySqrtConsistency(t *testing.T) {
	checks := []int{0, 1, 2, 3, 4, 8, 10, 15, 16, 17, 99, 100, 101, 2147395599, 2147395600, 2147483647}
	for _, x := range checks {
		t.Run(fmt.Sprintf("一致性校验_x=%d", x), func(t *testing.T) {
			r := MySqrt(x)
			if r*r > x || (r+1)*(r+1) <= x {
				t.Errorf("MySqrt(%d) = %d 不满足 r*r <= x < (r+1)*(r+1)", x, r)
			}
		})
	}
}

func BenchmarkMySqrt(b *testing.B) {
	benchmarks := []struct {
		name string
		x    int
	}{
		{"小值x=100", 100},
		{"中值x=1000000", 1000000},
		{"int32最大值", 2147483647},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MySqrt(bm.x)
			}
		})
	}
}
