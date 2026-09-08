package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: 11+1", "11", "1", "100"},
		{"示例2: 1010+1011", "1010", "1011", "10101"},

		// 边界：含 0
		{"a为0", "0", "0", "0"},
		{"0加非0", "0", "1", "1"},
		{"非0加0", "1", "0", "1"},

		// 边界：长度差异
		{"长度差1", "1", "111", "1000"},
		{"长度差较大", "1", "1111111", "10000000"},

		// 边界：全 1 相加产生连续进位
		{"全1相加", "111", "111", "1110"},
		{"全1加1", "111111", "1", "1000000"},

		// 边界：单个 1
		{"1加1", "1", "1", "10"},

		// 交错位相加
		{"交错位", "101", "101", "1010"},
		{"不同位", "110", "11", "1001"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddBinary(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("AddBinary(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func BenchmarkAddBinary(b *testing.B) {
	b.Run("短字符串", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AddBinary("1010", "1011")
		}
	})
	b.Run("长字符串1000位", func(b *testing.B) {
		a := makeLongBinary(1000, 1)
		c := makeLongBinary(1000, 3)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			AddBinary(a, c)
		}
	})
}

// makeLongBinary 生成长度为 n 的二进制字符串，每一位按 seed 的奇偶交替
func makeLongBinary(n, seed int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if (i+seed)%2 == 0 {
			buf[i] = '1'
		} else {
			buf[i] = '0'
		}
	}
	return string(buf)
}
