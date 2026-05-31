package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		want  []int
	}{
		// LeetCode 官方示例
		{"示例1: [0,1,0,3,12]", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"示例2: [0]", []int{0}, []int{0}},

		// 边界：全为0
		{"全为0", []int{0, 0, 0}, []int{0, 0, 0}},
		{"单个0", []int{0}, []int{0}},

		// 边界：全非0
		{"全非0", []int{1, 2, 3}, []int{1, 2, 3}},
		{"单个非0", []int{5}, []int{5}},

		// 边界：0在开头
		{"0在开头", []int{0, 1, 2, 3}, []int{1, 2, 3, 0}},
		{"多个0在开头", []int{0, 0, 1, 2}, []int{1, 2, 0, 0}},

		// 边界：0在结尾
		{"0在结尾", []int{1, 2, 3, 0}, []int{1, 2, 3, 0}},
		{"多个0在结尾", []int{1, 2, 0, 0}, []int{1, 2, 0, 0}},

		// 边界：0在中间
		{"0在中间", []int{1, 0, 3}, []int{1, 3, 0}},
		{"多个0在中间", []int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},

		// 边界：正负交替含0
		{"正负交替含0", []int{-1, 0, 2, 0, -3}, []int{-1, 2, -3, 0, 0}},

		// 较大数组
		{"混合数组", []int{0, 1, 0, 3, 12, 0, 5}, []int{1, 3, 12, 5, 0, 0, 0}},
		{"长数组无0", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			MoveZeroes(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("MoveZeroes(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums  []int
	}{
		{"len=10", []int{0, 1, 0, 3, 12, 0, 5, 0, 7, 0}},
		{"len=100", generateArrayWithZeros(100)},
		{"len=1000", generateArrayWithZeros(1000)},
		{"len=10000", generateArrayWithZeros(10000)},
		{"len=100000", generateArrayWithZeros(100000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nums := make([]int, len(bm.nums))
				copy(nums, bm.nums)
				MoveZeroes(nums)
			}
		})
	}
}

// generateArrayWithZeros 生成长度为 n 的数组，随机插入 0
func generateArrayWithZeros(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i%3 == 0 {
			arr[i] = 0
		} else {
			arr[i] = i
		}
	}
	return arr
}
