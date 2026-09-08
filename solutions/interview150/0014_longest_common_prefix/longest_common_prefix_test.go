package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		// LeetCode 官方示例
		{"示例1: flower/flow/flight", []string{"flower", "flow", "flight"}, "fl"},
		{"示例2: dog/racecar/car", []string{"dog", "racecar", "car"}, ""},

		// 边界：空数组（题目约束最少 1 个，此处验证健壮性）
		{"空数组", []string{}, ""},

		// 边界：单个字符串
		{"单个字符串", []string{"hello"}, "hello"},
		{"单个空字符串", []string{""}, ""},

		// 边界：含空字符串
		{"首字符串为空", []string{"", "abc"}, ""},
		{"中间字符串为空", []string{"abc", "", "abd"}, ""},
		{"尾字符串为空", []string{"abc", "abd", ""}, ""},

		// 边界：完全相同的字符串
		{"完全相同", []string{"abc", "abc", "abc"}, "abc"},

		// 边界：一个字符串是另一个的前缀
		{"前缀关系", []string{"ab", "abc", "abd"}, "ab"},
		{"长串在前", []string{"abcd", "ab", "abc"}, "ab"},

		// 边界：只有首字符相同
		{"仅首字符相同", []string{"apple", "apply", "apex"}, "ap"},
		{"首字符不同", []string{"cat", "dog"}, ""},

		// 边界：单字符
		{"单字符相同", []string{"a", "a", "a"}, "a"},
		{"单字符不同", []string{"a", "b"}, ""},

		// 边界：大小写敏感
		{"大小写敏感", []string{"Flower", "flow", "flight"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("LongestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	benchmarks := []struct {
		name string
		strs []string
	}{
		{"3个字符串", []string{"flower", "flow", "flight"}},
		{"5个字符串", []string{"interview", "internet", "internal", "interval", "integer"}},
		{"100个相同前缀", generateStrings(100, "commonprefix", "suffix")},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestCommonPrefix(bm.strs)
			}
		})
	}
}

// generateStrings 生成 n 个具有相同前缀的字符串
func generateStrings(n int, prefix, suffix string) []string {
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		strs[i] = prefix + suffix
	}
	return strs
}
