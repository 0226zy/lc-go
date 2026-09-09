package bestsightseeingpair

import "testing"

var sightseeingCases = []struct {
	name   string
	values []int
	want   int
}{
	{name: "示例1：[8,1,5,2,6]", values: []int{8, 1, 5, 2, 6}, want: 11},
	{name: "示例2：[1,2]", values: []int{1, 2}, want: 2},
	{name: "两元素最小数组", values: []int{1, 1}, want: 1},
	{name: "评分递增", values: []int{1, 2, 3, 4, 5}, want: 8},
	{name: "评分递减", values: []int{5, 4, 3, 2, 1}, want: 8},
	{name: "最优对在首尾", values: []int{10, 1, 1, 1, 10}, want: 16},
	{name: "相邻高分最优", values: []int{7, 8, 8, 1, 1}, want: 15},
}

func TestMaxScoreSightseeingPair(t *testing.T) {
	for _, tt := range sightseeingCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxScoreSightseeingPair(tt.values); got != tt.want {
				t.Errorf("MaxScoreSightseeingPair(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

func TestMaxScoreSightseeingPairOptimized(t *testing.T) {
	for _, tt := range sightseeingCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxScoreSightseeingPairOptimized(tt.values); got != tt.want {
				t.Errorf("MaxScoreSightseeingPairOptimized(%v) = %d, want %d", tt.values, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxScoreSightseeingPair(b *testing.B) {
	values := make([]int, 50000)
	for i := range values {
		values[i] = i%1000 + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxScoreSightseeingPair(values)
	}
}

func BenchmarkMaxScoreSightseeingPairOptimized(b *testing.B) {
	values := make([]int, 50000)
	for i := range values {
		values[i] = i%1000 + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxScoreSightseeingPairOptimized(values)
	}
}
