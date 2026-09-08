package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: n=2 有两种方法", 2, 2},
		{"示例2: n=3 有三种方法", 3, 3},
		{"补充示例: n=4 有五种方法", 4, 5},

		// 边界：最小输入
		{"边界: n=1 只有一种方法", 1, 1},

		// 斐波那契数列递推验证
		{"递推: n=5 等于 f(4)+f(3)=8", 5, 8},
		{"递推: n=6 等于 f(5)+f(4)=13", 6, 13},
		{"递推: n=7 等于 f(6)+f(5)=21", 7, 21},
		{"递推: n=10 等于 89", 10, 89},

		// 压力场景：约束最大值 n=45，结果仍在 int 范围内
		{"压力: n=44 等于 1134903170", 44, 1134903170},
		{"压力: n=45 等于 1836311903", 45, 1836311903},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"小规模 n=10", 10},
		{"中规模 n=25", 25},
		{"约束最大值 n=45", 45},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ClimbStairs(bm.n)
			}
		})
	}
}
