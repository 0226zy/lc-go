package containerwithmostwater

import (
	"math/rand"
	"testing"
	"time"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "官方示例1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "官方示例2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "两端高中间低",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "三个元素",
			height: []int{1, 2, 1},
			want:   2,
		},
		{
			name:   "两个元素递增",
			height: []int{2, 3},
			want:   2,
		},
		{
			name:   "两个元素递减",
			height: []int{3, 2},
			want:   2,
		},
		{
			name:   "全部相同",
			height: []int{5, 5, 5, 5},
			want:   15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxArea(tt.height)
			if got != tt.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

// makeRandomHeight 生成指定长度的随机高度数组（0 <= height[i] <= 10^4）
func makeRandomHeight(n int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	height := make([]int, n)
	for i := range height {
		height[i] = r.Intn(10001)
	}
	return height
}

// makeWorstCaseHeight 生成最坏情况下的高度数组（单调递减）
func makeWorstCaseHeight(n int) []int {
	height := make([]int, n)
	for i := range height {
		height[i] = n - i
	}
	return height
}

// makeBestCaseHeight 生成最好情况下的高度数组（两端高中间低）
func makeBestCaseHeight(n int) []int {
	height := make([]int, n)
	for i := 1; i < n-1; i++ {
		height[i] = 1
	}
	height[0] = 10000
	height[n-1] = 10000
	return height
}

func BenchmarkMaxArea_Official(b *testing.B) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxArea_Random_1e3(b *testing.B) {
	height := makeRandomHeight(1e3, time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxArea_Random_1e4(b *testing.B) {
	height := makeRandomHeight(1e4, time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxArea_Random_1e5(b *testing.B) {
	height := makeRandomHeight(1e5, time.Now().UnixNano())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxArea_WorstCase_1e5(b *testing.B) {
	height := makeWorstCaseHeight(1e5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}

func BenchmarkMaxArea_BestCase_1e5(b *testing.B) {
	height := makeBestCaseHeight(1e5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxArea(height)
	}
}
