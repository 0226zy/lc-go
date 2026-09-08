package plusone

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		// LeetCode 官方示例
		{"示例1: 123加一得124", []int{1, 2, 3}, []int{1, 2, 4}},
		{"示例2: 4321加一得4322", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"示例3: 9加一得10", []int{9}, []int{1, 0}},

		// 边界：单个数字
		{"单个数字0加一得1", []int{0}, []int{1}},
		{"单个数字1加一得2", []int{1}, []int{2}},
		{"单个数字8加一得9", []int{8}, []int{9}},

		// 进位：末位为 9 产生进位但不扩容
		{"末位9进位: 129加一得130", []int{1, 2, 9}, []int{1, 3, 0}},
		{"末位9进位: 19加一得20", []int{1, 9}, []int{2, 0}},

		// 进位：连续多位 9 级联进位
		{"连续9进位: 299加一得300", []int{2, 9, 9}, []int{3, 0, 0}},
		{"连续9进位: 1999加一得2000", []int{1, 9, 9, 9}, []int{2, 0, 0, 0}},

		// 全 9：需要扩容一位
		{"全9扩容: 99加一得100", []int{9, 9}, []int{1, 0, 0}},
		{"全9扩容: 999加一得1000", []int{9, 9, 9}, []int{1, 0, 0, 0}},

		// 边界：数组长度达到约束上限 100 且全为 9
		{"100个9扩容为1后跟100个0", repeat(9, 100), oneFollowedByZeros(100)},
		// 边界：数组长度为 100，正常无进位
		{"长度为100无进位", append(repeat(9, 99), 8), append(repeat(9, 99), 9)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusOne(tt.digits); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("PlusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	benchmarks := []struct {
		name   string
		digits []int
	}{
		{"无进位", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 8}},
		{"全9需要扩容", repeat(9, 100)},
		{"长度为100仅末位加一", append(repeat(1, 99), 8)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 每次拷贝一份，避免基准之间互相污染
				PlusOne(append([]int{}, bm.digits...))
			}
		})
	}
}

// repeat 生成包含 n 个 value 的切片
func repeat(value, n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = value
	}
	return out
}

// oneFollowedByZeros 生成 1 后跟 n 个 0 的切片（即 10 的 n 次方的位表示）
func oneFollowedByZeros(n int) []int {
	out := make([]int, n+1)
	out[0] = 1
	return out
}
