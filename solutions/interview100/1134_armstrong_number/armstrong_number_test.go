package armstrongnumber

import "testing"

func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 153是阿姆斯特朗数", 153, true},
		{"示例2: 123不是", 123, false},

		// 边界：一位数恒为阿姆斯特朗数（d^1 == d）
		{"最小一位数1", 1, true},
		{"一位数9", 9, true},

		// 经典阿姆斯特朗数
		{"三位数370", 370, true},
		{"三位数371", 371, true},
		{"四位数9474", 9474, true},
		{"四位数1634", 1634, true},

		// 非阿姆斯特朗数
		{"差1的9475", 9475, false},
		{"幂次混淆100", 100, false},
		{"较大反例12345678", 12345678, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArmstrong(tt.n); got != tt.want {
				t.Errorf("IsArmstrong(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkIsArmstrong(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"3位数", 153},
		{"8位数", 12345678},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsArmstrong(bm.n)
			}
		})
	}
}
