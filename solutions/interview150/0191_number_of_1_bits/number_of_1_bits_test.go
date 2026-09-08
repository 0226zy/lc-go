package numberof1bits

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want int
	}{
		// LeetCode 官方示例
		{"示例1: 11", 0b00000000000000000000000000001011, 3},
		{"示例2: 128", 0b00000000000000000000000010000000, 1},

		// 边界：全 0 和全 1
		{"全0", 0, 0},
		{"全1", 0xFFFFFFFF, 32},

		// 边界：单比特
		{"仅最低位为1", 1, 1},
		{"仅最高位为1", 1 << 31, 1},

		// 交错位：16 个 1
		{"交错位1010", 0xAAAAAAAA, 16},
		{"交错位0101", 0x55555555, 16},

		// 连续 1 段
		{"低8位全1", 0xFF, 8},
		{"高16位全1", 0xFFFF0000, 16},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingWeight(tt.n); got != tt.want {
				t.Errorf("HammingWeight(%032b) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	// 分别用 1 较少和 1 较多的数测试，验证性能与 1 的个数相关而非固定位数
	b.Run("稀疏位1个", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingWeight(1 << 30)
		}
	})
	b.Run("稠密位32个", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HammingWeight(0xFFFFFFFF)
		}
	})
}
