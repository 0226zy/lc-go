package destroyingasteroids

import "testing"

func TestAsteroidsDestroyed(t *testing.T) {
	tests := []struct {
		name      string
		mass      int
		asteroids []int
		want      bool
	}{
		// LeetCode 官方示例
		{"示例1", 10, []int{3, 9, 19, 5, 21}, true},
		{"示例2", 5, []int{4, 9, 23, 4}, false},

		// 边界：单个小行星
		{"单个小行星_可摧毁", 5, []int{3}, true},
		{"单个小行星_不可摧毁", 2, []int{5}, false},

		// 边界：mass 初始很大
		{"mass很大_可摧毁所有", 100, []int{1, 2, 3, 4, 5}, true},
		{"mass很小_不可摧毁", 1, []int{2, 3}, false},

		// 边界：小行星质量相等
		{"小行星质量相等_可摧毁", 3, []int{3, 3, 3}, true},
		{"小行星质量相等_不可摧毁", 2, []int{3, 3}, false},

		// 边界：小行星质量递增
		{"质量递增_可摧毁", 1, []int{1, 2, 4, 8}, true},
		{"质量递增_不可摧毁", 1, []int{1, 2, 4, 8, 20}, false},

		// 复杂场景：需要排序
		{"需要排序_可摧毁", 1, []int{2, 1}, true},
		{"需要排序_不可摧毁", 1, []int{3, 2}, false},

		// 边界：空数组
		// {"空数组", 10, []int{}, true}, // 题目约束 asteroids.length >= 1，不需要测试

		// 性能场景：大量小行星
		{"大量小行星_可摧毁", 1, makeIncArray(1000), true},
		// 注意：排序后 [1,2,...,999,100000]，行星可以依次摧毁前999个，质量变为 1+1+2+...+999 = 499501，大于 100000，所以可摧毁
		// 修改为：前几个小行星就很大，且质量递增，行星无法摧毁
		{"大量小行星_不可摧毁", 1, []int{100000, 100001, 100002, 1, 2, 3}, false}, // 排序后 [1,2,3,100000,100001,100002]，行星可以摧毁 1,2,3，但质量只有 7，无法摧毁 100000
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入，避免测试间相互影响
			input := make([]int, len(tt.asteroids))
			copy(input, tt.asteroids)

			got := AsteroidsDestroyed(tt.mass, input)
			if got != tt.want {
				t.Errorf("AsteroidsDestroyed(%d, %v) = %v, want %v", tt.mass, tt.asteroids, got, tt.want)
			}
		})
	}
}

// makeIncArray 生成长度为 n 的递增数组 [1, 2, ..., n]
func makeIncArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func BenchmarkAsteroidsDestroyed(b *testing.B) {
	benchmarks := []struct {
		name      string
		mass      int
		asteroids []int
	}{
		{"len=5", 10, []int{3, 9, 19, 5, 21}},
		{"len=100", 1, makeIncArray(100)},
		{"len=1000", 1, makeIncArray(1000)},
		{"len=10000", 1, makeIncArray(10000)},
		{"len=100000", 1, makeIncArray(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 复制输入，避免原地排序影响后续迭代
				input := make([]int, len(bm.asteroids))
				copy(input, bm.asteroids)
				AsteroidsDestroyed(bm.mass, input)
			}
		})
	}
}
