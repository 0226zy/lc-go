package permutations

import (
	"fmt"
	"sort"
	"testing"
)

// normalize 将二维切片中每个一维切片排序后，再按字典序整体排序，
// 用于对"顺序无关"的排列结果做比较。
func normalize(perms [][]int) []string {
	keys := make([]string, len(perms))
	for i, p := range perms {
		cp := make([]int, len(p))
		copy(cp, p)
		sort.Ints(cp)
		keys[i] = fmt.Sprint(cp)
	}
	sort.Strings(keys)
	return keys
}

// equalPermutations 无序比较两个排列集合是否相同。
func equalPermutations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	ka, kb := normalize(a), normalize(b)
	for i := range ka {
		if ka[i] != kb[i] {
			return false
		}
	}
	return true
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "官方示例1: 三元素数组",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name: "官方示例2: 两元素数组",
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "官方示例3: 单元素数组",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "边界: 包含负数",
			nums: []int{-1, 2},
			want: [][]int{{-1, 2}, {2, -1}},
		},
		{
			name: "边界: 包含0和负数的三个元素",
			nums: []int{0, -1, 2},
			want: [][]int{
				{0, -1, 2}, {0, 2, -1}, {-1, 0, 2}, {-1, 2, 0}, {2, 0, -1}, {2, -1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Permute(tt.nums)
			if !equalPermutations(got, tt.want) {
				t.Errorf("Permute(%v) = %v, 期望（无序）%v", tt.nums, got, tt.want)
			}
		})
	}
}

// TestPermuteCountAndUnique 验证 n 个元素恰好产生 n! 个互不重复的排列。
func TestPermuteCountAndUnique(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := Permute(nums)

	if len(got) != 24 {
		t.Fatalf("Permute(%v) 返回 %d 个排列, 期望 24", nums, len(got))
	}

	seen := make(map[string]bool, len(got))
	for _, p := range got {
		if len(p) != len(nums) {
			t.Fatalf("排列长度错误: %v", p)
		}
		key := fmt.Sprint(p)
		if seen[key] {
			t.Fatalf("出现重复排列: %v", p)
		}
		seen[key] = true
	}
}

func BenchmarkPermute(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Permute(nums)
	}
}
