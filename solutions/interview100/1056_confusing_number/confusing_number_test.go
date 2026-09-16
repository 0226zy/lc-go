package confusingnumber

import "testing"

func TestConfusingNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: 6", 6, true},
		{"示例2: 89", 89, true},
		{"示例3: 11", 11, false},
		{"示例4: 25", 25, false},

		// 边界：0 与单位数
		{"0 旋转后仍是 0", 0, false},
		{"单位数 1", 1, false},
		{"单位数 8", 8, false},
		{"单位数 9", 9, true},

		// 边界：包含无效数字
		{"包含 2", 12, false},
		{"包含 5", 56, false},
		{"包含 7", 917, false},

		// 边界：旋转后与原数相同
		{"对称数 88", 88, false},
		{"对称数 101", 101, false},
		{"对称数 181", 181, false},
		{"对称数 69 旋转后仍是 69", 69, false},

		// 边界：前导零
		{"8000 旋转后是 0008 即 8", 8000, true},
		{"100 旋转后是 001 即 1", 100, true},
		{"10 旋转后是 01 即 1", 10, true},

		// 典型场景
		{"609 旋转后是 609", 609, false},
		{"916 旋转后是 916", 916, false},
		{"大数 1000000000", 1000000000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConfusingNumber(tt.n); got != tt.want {
				t.Errorf("ConfusingNumber(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkConfusingNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConfusingNumber(689069816)
	}
}
