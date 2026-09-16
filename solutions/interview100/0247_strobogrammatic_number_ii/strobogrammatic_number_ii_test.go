package strobogrammaticnumberii

import (
	"sort"
	"testing"
)

// equalStringSliceUnordered 比较两个字符串切片是否包含相同元素（顺序无关）
func equalStringSliceUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]string, len(a))
	bCopy := make([]string, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Strings(aCopy)
	sort.Strings(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

func TestFindStrobogrammatic(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: n=2", 2, []string{"11", "69", "88", "96"}},
		{"示例2: n=1", 1, []string{"0", "1", "8"}},

		// 典型场景：n=3，中间位可取 0/1/8，两侧为 11/69/88/96
		{"n=3", 3, []string{
			"101", "609", "808", "906",
			"111", "619", "818", "916",
			"181", "689", "888", "986",
		}},

		// 边界：n=4，两侧两层数对，不能含 0 开头
		{"n=4", 4, []string{
			"1001", "1111", "1691", "1881", "1961",
			"6009", "6119", "6699", "6889", "6969",
			"8008", "8118", "8698", "8888", "8968",
			"9006", "9116", "9696", "9886", "9966",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindStrobogrammatic(tt.n)
			if !equalStringSliceUnordered(got, tt.want) {
				t.Errorf("FindStrobogrammatic(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

// TestFindStrobogrammaticCount 校验较大 n 时解的个数，并抽查每个解确实中心对称
func TestFindStrobogrammaticCount(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int // 解的个数：偶数 n = 4 * 5^(n/2-1)，奇数 n = 4 * 3 * 5^((n-3)/2)（n>=2），n=1 为 3
	}{
		{"n=1 共3个", 1, 3},
		{"n=2 共4个", 2, 4},
		{"n=3 共12个", 3, 12},
		{"n=4 共20个", 4, 20},
		{"n=5 共60个", 5, 60},
		{"n=6 共100个", 6, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindStrobogrammatic(tt.n)
			if len(got) != tt.want {
				t.Errorf("FindStrobogrammatic(%d) 返回 %d 个结果, want %d", tt.n, len(got), tt.want)
			}
			for _, num := range got {
				if !isStrobogrammatic(num) {
					t.Errorf("FindStrobogrammatic(%d) 产生了非中心对称数 %q", tt.n, num)
				}
				if tt.n > 1 && num[0] == '0' {
					t.Errorf("FindStrobogrammatic(%d) 产生了带前导零的 %q", tt.n, num)
				}
			}
		})
	}
}

// isStrobogrammatic 验证一个数字字符串是否为中心对称数
func isStrobogrammatic(num string) bool {
	rotate := map[byte]byte{'0': '0', '1': '1', '6': '9', '8': '8', '9': '6'}
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		r, ok := rotate[num[i]]
		if !ok || r != num[j] {
			return false
		}
	}
	return true
}

func BenchmarkFindStrobogrammatic(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=4", 4},
		{"n=8", 8},
		{"n=12", 12},
		{"n=14", 14},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindStrobogrammatic(bm.n)
			}
		})
	}
}
