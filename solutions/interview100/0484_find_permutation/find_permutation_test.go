package findpermutation

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestFindPermutation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		// LeetCode 官方示例
		{"示例1: 单个I", "I", []int{1, 2}},
		{"示例2: DI", "DI", []int{2, 1, 3}},

		// 边界：单字符
		{"单个D", "D", []int{2, 1}},

		// 边界：全为同一种字符
		{"全为I", "III", []int{1, 2, 3, 4}},
		{"全为D", "DDD", []int{4, 3, 2, 1}},

		// 典型场景
		{"D开头交错", "DDI", []int{3, 2, 1, 4}},
		{"I与D交替", "IDID", []int{1, 3, 2, 5, 4}},
		{"两段D被I隔开", "DDIIDD", []int{3, 2, 1, 4, 7, 6, 5}},
		{"D结尾", "ID", []int{1, 3, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPermutation(tt.s); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("FindPermutation(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkFindPermutation(b *testing.B) {
	// 构造长字符串：一半交替、一半全 D
	long := make([]byte, 0, 100000)
	for len(long) < 50000 {
		long = append(long, 'I', 'D')
	}
	for len(long) < 100000 {
		long = append(long, 'D')
	}

	benchmarks := []struct {
		name string
		s    string
	}{
		{"len=2", "DI"},
		{"len=1000 交替", string(long[:1000])},
		{"len=100000 混合", string(long)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindPermutation(bm.s)
			}
		})
	}
}
