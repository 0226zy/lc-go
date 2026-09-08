package threesum

import (
	"sort"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

// sortTriplets 将三元组切片按字典序排序，使整体顺序无关的比较成为可能
func sortTriplets(triplets [][]int) {
	sort.Slice(triplets, func(i, j int) bool {
		for k := 0; k < 3; k++ {
			if triplets[i][k] != triplets[j][k] {
				return triplets[i][k] < triplets[j][k]
			}
		}
		return false
	})
}

// equalTriplets 整体顺序无关地比较两个三元组切片（假定每个三元组内部已按升序排列）
func equalTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([][]int, len(a))
	copy(aCopy, a)
	bCopy := make([][]int, len(b))
	copy(bCopy, b)
	sortTriplets(aCopy)
	sortTriplets(bCopy)
	return utils.Equal2DIntSlice(aCopy, bCopy)
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// LeetCode 官方示例
		{"示例1: 常规混合数组", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"示例2: 无解三元素", []int{0, 1, 1}, [][]int{}},
		{"示例3: 全零三元素", []int{0, 0, 0}, [][]int{{0, 0, 0}}},

		// 边界：最小长度
		{"恰好三元素且和为0", []int{-1, 0, 1}, [][]int{{-1, 0, 1}}},
		{"恰好三元素但和不为0", []int{1, 2, 3}, [][]int{}},

		// 边界：大量重复元素需要去重
		{"全零多元素只输出一组", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"重复正负对需要去重", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"相同值出现三次只取一次", []int{-1, -1, -1, 2}, [][]int{{-1, -1, 2}}},

		// 边界：全正或全负无解
		{"全正数无解", []int{1, 2, 3, 4, 5}, [][]int{}},
		{"全负数无解", []int{-5, -4, -3, -2, -1}, [][]int{}},

		// 边界：含极值元素
		{"含极值-10^5与10^5", []int{-100000, 0, 100000}, [][]int{{-100000, 0, 100000}}},
		{"极值无法配对", []int{-100000, 1, 100000}, [][]int{}},

		// 多组解
		{"多组解混合数组", []int{-4, -2, -2, -1, 0, 1, 2, 2, 3, 4}, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制一份输入，避免排序污染表驱动的原始数据
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			if got := ThreeSum(nums); !equalTriplets(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	// 压力场景：3000 个元素，在 [-10^5, 10^5] 范围内伪随机分布
	const n = 3000
	nums := make([]int, n)
	seed := int64(42)
	for i := range nums {
		// 简单的线性同余伪随机数，保证每次 benchmark 数据一致
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		nums[i] = int(seed%200001) - 100000
	}

	b.Run("3000个随机元素", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numsCopy := make([]int, n)
			copy(numsCopy, nums)
			ThreeSum(numsCopy)
		}
	})

	// 极端场景：3000 个全零，双指针大量跳过重复值
	zeros := make([]int, n)
	b.Run("3000个全零元素", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			zerosCopy := make([]int, n)
			copy(zerosCopy, zeros)
			ThreeSum(zerosCopy)
		}
	})
}
