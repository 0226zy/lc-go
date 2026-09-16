package handshakesthatdontcross

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name      string
		numPeople int
		want      int
	}{
		// LeetCode 官方示例
		{"示例1: 2人", 2, 1},
		{"示例2: 4人", 4, 2},
		{"示例3: 6人", 6, 5},
		{"示例4: 8人", 8, 14},

		// 卡特兰数序列核对：C(0)=1, C(5)=42, C(6)=132, C(10)=16796
		{"0人", 0, 1},
		{"10人", 10, 42},
		{"12人", 12, 132},
		{"20人", 20, 16796},

		// 边界：奇数个人无法两两配对（防御性场景，题目保证输入为偶数）
		{"奇数个人", 3, 0},
		{"奇数1人", 1, 0},

		// 大规模：答案需要取模
		{"1000人取模", 1000, 581699957},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfWays(tt.numPeople); got != tt.want {
				t.Errorf("NumberOfWays(%d) = %d, want %d", tt.numPeople, got, tt.want)
			}
		})
	}
}

func BenchmarkNumberOfWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumberOfWays(1000)
	}
}
