package topkfrequentelements

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{
			name: "示例1：基本用例",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
		},
		{
			name: "示例2：单元素",
			nums: []int{1},
			k:    1,
		},
		{
			name: "全部不同频率",
			nums: []int{1, 1, 2, 2, 2, 3, 3, 3, 3},
			k:    3,
		},
		{
			name: "k等于不同元素数",
			nums: []int{1, 2, 3},
			k:    3,
		},
		{
			name: "包含负数",
			nums: []int{-1, -1, 2, 2, 2, -3},
			k:    2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKFrequent(tt.nums, tt.k)
			if got == nil {
				t.Skip("未实现")
				return
			}
			if len(got) != tt.k {
				t.Errorf("TopKFrequent(%v, %d) 返回了 %d 个元素，需要 %d 个", tt.nums, tt.k, len(got), tt.k)
			}
			// 验证返回元素的频率确实是最高的
			freq := make(map[int]int)
			for _, n := range tt.nums {
				freq[n]++
			}
			type pair struct{ val, cnt int }
			pairs := make([]pair, 0, len(freq))
			for v, c := range freq {
				pairs = append(pairs, pair{v, c})
			}
			sort.Slice(pairs, func(i, j int) bool {
				return pairs[i].cnt > pairs[j].cnt
			})
			kthFreq := pairs[tt.k-1].cnt
			for _, v := range got {
				if freq[v] < kthFreq {
					t.Errorf("元素 %d 的频率 %d 低于第 %d 高的频率 %d", v, freq[v], tt.k, kthFreq)
				}
			}
		})
	}
}

func BenchmarkTopKFrequent(b *testing.B) {
	nums := []int{1, 1, 1, 2, 2, 3}
	for i := 0; i < b.N; i++ {
		TopKFrequent(nums, 2)
	}
}
