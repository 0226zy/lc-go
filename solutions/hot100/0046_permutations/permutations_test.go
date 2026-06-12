package permutations

import (
	"fmt"
	"sort"
	"testing"
)

// intSliceKey 将 []int 转换为可用于 map 的字符串键
func intSliceKey(a []int) string {
	return fmt.Sprintf("%v", a)
}

// equalPermutations 比较两个排列集合是否相等（顺序无关）
func equalPermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[string]int)
	for _, p := range a {
		count[intSliceKey(p)]++
	}
	for _, p := range b {
		key := intSliceKey(p)
		count[key]--
		if count[key] < 0 {
			return false
		}
	}
	return true
}

// isValidPermutation 检查 perm 是否是 nums 的一个有效排列
func isValidPermutation(perm, nums []int) bool {
	if len(perm) != len(nums) {
		return false
	}
	a := append([]int(nil), perm...)
	b := append([]int(nil), nums...)
	sort.Ints(a)
	sort.Ints(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// countPermutations 计算 n 的阶乘
func countPermutations(n int) int {
	if n <= 1 {
		return 1
	}
	return n * countPermutations(n-1)
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "示例1",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
				{2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "示例2",
			nums: []int{0, 1},
			want: [][]int{
				{0, 1}, {1, 0},
			},
		},
		{
			name: "示例3-单元素",
			nums: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			name: "包含负数",
			nums: []int{-1, 0, 1},
			want: [][]int{
				{-1, 0, 1}, {-1, 1, 0}, {0, -1, 1},
				{0, 1, -1}, {1, -1, 0}, {1, 0, -1},
			},
		},
		{
			name: "两个负数",
			nums: []int{-2, -1},
			want: [][]int{
				{-2, -1}, {-1, -2},
			},
		},
		{
			name: "最大长度",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: nil, // 不直接比较，只检查数量和合法性
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.nums)

			expectedCount := countPermutations(len(tt.nums))
			if len(got) != expectedCount {
				t.Errorf("Permute() returned %d permutations, want %d", len(got), expectedCount)
			}

			for _, p := range got {
				if !isValidPermutation(p, tt.nums) {
					t.Errorf("Permute() returned invalid permutation %v for nums %v", p, tt.nums)
				}
			}

			if tt.want != nil && !equalPermutations(got, tt.want) {
				t.Errorf("Permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermute_NoDuplicates(t *testing.T) {
	nums := []int{1, 2, 3}
	got := Permute(nums)
	seen := make(map[string]bool)
	for _, p := range got {
		key := intSliceKey(p)
		if seen[key] {
			t.Errorf("Permute() returned duplicate permutation %v", p)
		}
		seen[key] = true
	}
}

func BenchmarkPermute(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		_ = Permute(nums)
	}
}
