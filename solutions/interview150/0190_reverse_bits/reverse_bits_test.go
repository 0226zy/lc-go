package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		// LeetCode 官方示例
		{"示例1: 0x00000095", 0b00000010100101000001111010011100, 0b00111001011110000010100101000000},

		// 边界：全 0 和全 1，反转后不变
		{"全0", 0, 0},
		{"全1", 0xFFFFFFFF, 0xFFFFFFFF},

		// 边界：只有最低位为 1，反转后只有最高位为 1
		{"仅最低位为1", 1, 1 << 31},
		{"仅最高位为1", 1 << 31, 1},

		// 边界：高低半位互换
		{"0xFFFF0000", 0xFFFF0000, 0x0000FFFF},
		{"0x0000FFFF", 0x0000FFFF, 0xFFFF0000},

		// 交错位
		{"0xAAAAAAAA", 0xAAAAAAAA, 0x55555555},
		{"0x55555555", 0x55555555, 0xAAAAAAAA},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBits(tt.n); got != tt.want {
				t.Errorf("ReverseBits(%032b) = %032b, want %032b", tt.n, got, tt.want)
			}
		})
	}
}

// TestReverseBitsInverse 验证“反转两次等于自身”这一性质，随机位模式批量检验
func TestReverseBitsInverse(t *testing.T) {
	for _, n := range []uint32{0x12345678, 0x89ABCDEF, 0x0F0F0F0F, 0xDEADBEEF, 7, 1 << 30} {
		if got := ReverseBits(ReverseBits(n)); got != n {
			t.Errorf("ReverseBits(ReverseBits(%032b)) = %032b", n, got)
		}
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBits(0b00000010100101000001111010011100)
	}
}
