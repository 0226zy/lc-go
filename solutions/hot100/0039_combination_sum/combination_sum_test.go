package combinationsum

import (
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		wantLen    int // 期望的组合数量
	}{
		{
			name:       "示例1：基本用例",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			wantLen:    2,
		},
		{
			name:       "示例2：多个组合",
			candidates: []int{2, 3, 5},
			target:     8,
			wantLen:    3,
		},
		{
			name:       "示例3：无解",
			candidates: []int{2},
			target:     1,
			wantLen:    0,
		},
		{
			name:       "target等于候选值",
			candidates: []int{1},
			target:     1,
			wantLen:    1,
		},
		{
			name:       "重复选取同一候选",
			candidates: []int{1},
			target:     3,
			wantLen:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)
			if got == nil {
				t.Skip("未实现")
				return
			}
			if len(got) != tt.wantLen {
				t.Errorf("CombinationSum(%v, %d) 返回 %d 个组合, 期望 %d 个", tt.candidates, tt.target, len(got), tt.wantLen)
			}
			// 验证每个组合的和等于 target
			for _, combo := range got {
				sum := 0
				for _, v := range combo {
					sum += v
				}
				if sum != tt.target {
					t.Errorf("组合 %v 的和为 %d, 不等于 target %d", combo, sum, tt.target)
				}
				// 验证组合已排序（去重保证）
				if !sort.IntsAreSorted(combo) {
					t.Errorf("组合 %v 未排序", combo)
				}
			}
		})
	}
}

func BenchmarkCombinationSum(b *testing.B) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	for i := 0; i < b.N; i++ {
		CombinationSum(candidates, target)
	}
}
