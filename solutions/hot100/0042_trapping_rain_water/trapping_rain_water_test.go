package trappingrainwater

import "testing"

func TestTrap(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "示例1：官方用例",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "示例2：凹槽用例",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "边界：只有两个柱子",
			height: []int{1, 2},
			want:   0,
		},
		{
			name:   "边界：完全递增",
			height: []int{1, 2, 3, 4, 5},
			want:   0,
		},
		{
			name:   "边界：完全递减",
			height: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "边界：全等高",
			height: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "单柱子",
			height: []int{5},
			want:   0,
		},
		{
			name:   "V型",
			height: []int{3, 0, 3},
			want:   3,
		},
		{
			name:   "多个凹槽",
			height: []int{2, 0, 1, 0, 2},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Trap(tt.height)
			if got != tt.want {
				t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}

func BenchmarkTrap(b *testing.B) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	for i := 0; i < b.N; i++ {
		Trap(height)
	}
}
