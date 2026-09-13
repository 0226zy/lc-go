package twosumii

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

var twoSumCases = []struct {
	name     string
	numbers  []int
	target   int
	want     []int
}{
	// LeetCode 官方示例
	{name: "示例1：[2,7,11,15] target=9", numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
	{name: "示例2：[2,3,4] target=6", numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
	{name: "示例3：[-1,0] target=-1", numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},

	// 边界：仅两个元素
	{name: "两元素刚好", numbers: []int{1, 2}, target: 3, want: []int{1, 2}},
	{name: "两元素负数", numbers: []int{-5, -3}, target: -8, want: []int{1, 2}},

	// 边界：解在两端 / 相邻
	{name: "解在两端", numbers: []int{1, 2, 3, 4, 8}, target: 9, want: []int{1, 5}},
	{name: "解偏右侧", numbers: []int{1, 2, 3, 4, 5}, target: 9, want: []int{4, 5}},
	{name: "相邻两数", numbers: []int{1, 3, 4, 5, 7, 10, 11}, target: 9, want: []int{3, 4}},

	// 边界：含零、重复值
	{name: "含零", numbers: []int{0, 0, 3, 4}, target: 0, want: []int{1, 2}},
	{name: "重复值", numbers: []int{1, 1, 1, 1, 2}, target: 3, want: []int{1, 5}},

	// 边界：较大区间（唯一解）
	{name: "较大数", numbers: []int{-1000, -500, 0, 500, 999}, target: -1, want: []int{1, 5}},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range twoSumCases {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.numbers, tt.target)
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i
	}
	target := 19997 // 倒数第二与倒数第一之和附近
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(numbers, target)
	}
}
