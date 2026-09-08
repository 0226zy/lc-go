package permutations

import (
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3]", []int{1, 2, 3}, [][]int{
			{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
		}},
		{"示例2: [0,1]", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"示例3: [1]", []int{1}, [][]int{{1}}},

		// 边界：含 0 和负数
		{"含0和负数", []int{0, -1, 2}, [][]int{
			{0, -1, 2}, {0, 2, -1}, {-1, 0, 2}, {-1, 2, 0}, {2, 0, -1}, {2, -1, 0},
		}},
		// 边界：四个元素
		{"四个元素数量验证", []int{1, 2, 3, 4}, nil}, // 数量在下方单独断言
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.nums)
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Permute(%v) = %v, want %v", tt.nums, got, tt.want)
				}
				return
			}
			// n=4 时应有 4! = 24 种排列
			if len(got) != 24 {
				t.Errorf("Permute(%v) 长度 = %d, want 24", tt.nums, len(got))
			}
		})
	}
}

func TestPermuteProperties(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := Permute(nums)

	seen := make(map[string]bool, len(got))
	for _, p := range got {
		if len(p) != len(nums) {
			t.Fatalf("排列长度不对: %v", p)
		}
		sum, product := 0, 1
		for _, v := range p {
			sum += v
			product *= v
		}
		if sum != 10 || product != 24 {
			t.Fatalf("排列 %v 元素不对（元素应恰好是 1,2,3,4）", p)
		}
		key := ""
		for _, v := range p {
			key += string(rune(v + '0'))
		}
		if seen[key] {
			t.Fatalf("出现重复排列 %v", p)
		}
		seen[key] = true
	}
}

func BenchmarkPermute(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
	}{
		{"n=6", []int{1, 2, 3, 4, 5, 6}},
		{"n=7", []int{1, 2, 3, 4, 5, 6, 7}},
		{"n=8", []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Permute(bm.nums)
			}
		})
	}
}
