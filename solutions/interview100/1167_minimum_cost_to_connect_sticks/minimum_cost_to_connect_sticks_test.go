package minimumcosttoconnectsticks

import "testing"

func TestConnectSticks(t *testing.T) {
	tests := []struct {
		name   string
		sticks []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: [2,4,3]", []int{2, 4, 3}, 14},
		{"示例2: [1,8,3,5]", []int{1, 8, 3, 5}, 30},
		{"示例3: 单根木棍", []int{5}, 0},

		// 边界：空输入
		{"空数组", []int{}, 0},

		// 两根木棍：费用就是两数之和
		{"两根木棍", []int{3, 7}, 10},

		// 长度全部相同：4 根 1 -> 1+1=2, 1+1=2, 2+2=4，总 8
		{"全部等长", []int{1, 1, 1, 1}, 8},

		// 已排序与逆序输入结果一致
		{"升序输入", []int{1, 3, 5, 8}, 30},
		{"降序输入", []int{8, 5, 3, 1}, 30},

		// 极端值：长度均为上限
		{"大长度木棍", []int{10000, 10000, 10000}, 50000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectSticks(tt.sticks); got != tt.want {
				t.Errorf("ConnectSticks(%v) = %d, want %d", tt.sticks, got, tt.want)
			}
		})
	}
}

func BenchmarkConnectSticks(b *testing.B) {
	// 构造 10000 根木棍的极端规模输入
	sticks := make([]int, 10000)
	for i := range sticks {
		sticks[i] = i%10000 + 1
	}

	b.Run("10000根木棍", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ConnectSticks(sticks)
		}
	})
}
