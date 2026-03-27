package tests

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
	"github.com/0226zy/lc-go/solutions/easy"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"示例1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"示例2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"示例3", []int{3, 3}, 6, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := easy.TwoSum(tt.nums, tt.target)
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
