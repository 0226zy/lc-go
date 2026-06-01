package mincostcandies

import "testing"

func TestMinimumCost(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "示例1-三颗糖果",
			cost: []int{1, 2, 3},
			want: 5,
		},
		{
			name: "示例2-六颗糖果",
			cost: []int{6, 5, 7, 9, 2, 2},
			want: 23,
		},
		{
			name: "示例3-两颗糖果",
			cost: []int{5, 5},
			want: 10,
		},
		{
			name: "一颗糖果",
			cost: []int{10},
			want: 10,
		},
		{
			name: "四颗糖果",
			cost: []int{1, 2, 3, 4},
			want: 8, // 降序: 4,3,2,1 -> 买4,3, 免费2, 买1 = 8
		},
		{
			name: "全部相同价格",
			cost: []int{5, 5, 5, 5, 5},
			want: 20, // 降序: 5,5,5,5,5 -> 买5,5 免费5, 买5,5 = 20
		},
		{
			name: " prices递增",
			cost: []int{1, 2, 3, 4, 5, 6},
			want: 16, // 降序: 6,5,4,3,2,1 -> 买6,5 免费4, 买3,2 免费1 = 16
		},
		{
			name: " prices递减",
			cost: []int{6, 5, 4, 3, 2, 1},
			want: 16, // 同上
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制切片，避免排序修改原数据影响后续调试
			cost := make([]int, len(tt.cost))
			copy(cost, tt.cost)
			got := MinimumCost(cost)
			if got != tt.want {
				t.Errorf("MinimumCost(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}

func BenchmarkMinimumCost(b *testing.B) {
	cost := []int{6, 5, 7, 9, 2, 2, 1, 3, 8, 4, 10, 15, 12, 11, 20, 18, 25, 30, 22, 28}
	for i := 0; i < b.N; i++ {
		c := make([]int, len(cost))
		copy(c, cost)
		MinimumCost(c)
	}
}
