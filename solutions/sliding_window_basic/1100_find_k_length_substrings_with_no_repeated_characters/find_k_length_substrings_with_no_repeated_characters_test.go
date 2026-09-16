package findklengthsubstringswithnorepeatedcharacters

import "testing"

func TestNumKLenSubstrNoRepeats(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: havefunonleetcode", "havefunonleetcode", 5, 6},
		{"示例2: k 大于字符串长度", "home", 5, 0},

		// 边界：k = 1，每个字符自成一个合法子串
		{"k等于1", "abca", 1, 4},
		{"单字符k等于1", "a", 1, 1},

		// 边界：k 等于字符串长度
		{"k等于长度且无重复", "abcd", 4, 1},
		{"k等于长度但有重复", "abca", 4, 0},

		// 边界：全部字符相同
		{"全部重复字符", "aaaa", 2, 0},

		// 典型场景
		{"周期串abcabc", "abcabc", 3, 4},
		{"无重复字符串", "abcdef", 3, 4},
		{"重复字符在窗口两端", "abba", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumKLenSubstrNoRepeats(tt.s, tt.k); got != tt.want {
				t.Errorf("NumKLenSubstrNoRepeats(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkNumKLenSubstrNoRepeats(b *testing.B) {
	// 构造长字符串：abc...z 循环，共 10000 个字符
	long := make([]byte, 0, 10000)
	for len(long) < 10000 {
		for c := byte('a'); c <= byte('z'); c++ {
			long = append(long, c)
		}
	}
	long = long[:10000]

	benchmarks := []struct {
		name string
		s    string
		k    int
	}{
		{"len=100,k=5", string(long[:100]), 5},
		{"len=1000,k=26", string(long[:1000]), 26},
		{"len=10000,k=26", string(long), 26},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NumKLenSubstrNoRepeats(bm.s, bm.k)
			}
		})
	}
}
