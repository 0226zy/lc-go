package lettercombinations

import (
	"sort"
	"testing"
)

var phoneMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

var implementations = []struct {
	name string
	fn   func(string) []string
}{
	{"backtrack", LetterCombinations},
	{"iterative", LetterCombinationsIterative},
}

// equalStringSlice 比较两个字符串切片（忽略顺序，nil 与空切片视为相等）
func equalStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	sort.Strings(a)
	bCopy := make([]string, len(b))
	copy(bCopy, b)
	sort.Strings(bCopy)
	for i := range a {
		if a[i] != bCopy[i] {
			return false
		}
	}
	return true
}

// generateExpected 根据 digits 生成所有期望的字母组合
func generateExpected(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	var dfs func(idx int, path []byte)
	dfs = func(idx int, path []byte) {
		if idx == len(digits) {
			res = append(res, string(path))
			return
		}
		for _, ch := range phoneMap[digits[idx]] {
			path = append(path, ch[0])
			dfs(idx+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, nil)
	return res
}

// expectedCount 计算期望的组合数量
func expectedCount(digits string) int {
	if len(digits) == 0 {
		return 0
	}
	count := 1
	for i := 0; i < len(digits); i++ {
		count *= len(phoneMap[digits[i]])
	}
	return count
}

// TestLetterCombinations 表驱动测试，覆盖官方示例及关键边界
func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{
			name:   "示例1-两个数字",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "示例2-空字符串",
			digits: "",
			want:   []string{},
		},
		{
			name:   "单个数字-2",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},
		{
			name:   "单个数字-3",
			digits: "3",
			want:   []string{"d", "e", "f"},
		},
		{
			name:   "单个数字-4",
			digits: "4",
			want:   []string{"g", "h", "i"},
		},
		{
			name:   "单个数字-5",
			digits: "5",
			want:   []string{"j", "k", "l"},
		},
		{
			name:   "单个数字-6",
			digits: "6",
			want:   []string{"m", "n", "o"},
		},
		{
			name:   "单个数字-7",
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			name:   "单个数字-8",
			digits: "8",
			want:   []string{"t", "u", "v"},
		},
		{
			name:   "单个数字-9",
			digits: "9",
			want:   []string{"w", "x", "y", "z"},
		},
		{
			name:   "最大长度4个数字",
			digits: "2345",
			want: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil",
				"aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil",
				"afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil",
				"bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil",
				"begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil",
				"bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil",
				"cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil",
				"cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil",
				"cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil",
			},
		},
	}

	for _, impl := range implementations {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				got := impl.fn(tt.digits)
				if !equalStringSlice(got, tt.want) {
					t.Errorf("%s(%q) = %v, want %v", impl.name, tt.digits, got, tt.want)
				}
			})
		}
	}
}

// TestLetterCombinationsAllDigits 枚举所有可能的 digits 组合进行验证
func TestLetterCombinationsAllDigits(t *testing.T) {
	digitsChars := []byte{'2', '3', '4', '5', '6', '7', '8', '9'}

	// 生成所有长度 0~4 的 digits 字符串
	var allDigits []string
	var dfs func(path []byte, maxLen int)
	dfs = func(path []byte, maxLen int) {
		if len(path) > 0 {
			allDigits = append(allDigits, string(path))
		}
		if len(path) == maxLen {
			return
		}
		for _, ch := range digitsChars {
			path = append(path, ch)
			dfs(path, maxLen)
			path = path[:len(path)-1]
		}
	}
	// 单独加入空字符串
	allDigits = append(allDigits, "")
	dfs(nil, 4)

	for _, impl := range implementations {
		for _, digits := range allDigits {
			t.Run(impl.name+"/digits_"+digits, func(t *testing.T) {
				got := impl.fn(digits)

				// 验证返回数量
				wantCount := expectedCount(digits)
				if len(got) != wantCount {
					t.Errorf("%s(%q) len=%d, want %d", impl.name, digits, len(got), wantCount)
					return
				}

				// 验证每个组合的正确性
				wantSet := make(map[string]bool)
				for _, s := range generateExpected(digits) {
					wantSet[s] = true
				}
				gotSet := make(map[string]bool)
				for _, s := range got {
					// 长度检查
					if len(s) != len(digits) {
						t.Errorf("%s(%q) got invalid length %q", impl.name, digits, s)
						return
					}
					// 字符有效性检查
					for i := 0; i < len(s); i++ {
						valid := false
						for _, ch := range phoneMap[digits[i]] {
							if ch == string(s[i]) {
								valid = true
								break
							}
						}
						if !valid {
							t.Errorf("%s(%q) got invalid char %q at pos %d in %q", impl.name, digits, s[i], i, s)
							return
						}
					}
					if gotSet[s] {
						t.Errorf("%s(%q) duplicate result %q", impl.name, digits, s)
						return
					}
					gotSet[s] = true
				}

				// 验证集合相等
				if len(gotSet) != len(wantSet) {
					t.Errorf("%s(%q) result set size mismatch", impl.name, digits)
				}
				for s := range wantSet {
					if !gotSet[s] {
						t.Errorf("%s(%q) missing expected result %q", impl.name, digits, s)
					}
				}
			})
		}
	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterCombinations("2345")
	}
}

func BenchmarkLetterCombinationsIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterCombinationsIterative("2345")
	}
}

// BenchmarkLetterCombinationsAllDigits 对所有 digits 组合进行基准测试
func BenchmarkLetterCombinationsAllDigits(b *testing.B) {
	digitsChars := []byte{'2', '3', '4', '5', '6', '7', '8', '9'}
	var allDigits []string
	var dfs func(path []byte, maxLen int)
	dfs = func(path []byte, maxLen int) {
		if len(path) > 0 {
			allDigits = append(allDigits, string(path))
		}
		if len(path) == maxLen {
			return
		}
		for _, ch := range digitsChars {
			path = append(path, ch)
			dfs(path, maxLen)
			path = path[:len(path)-1]
		}
	}
	dfs(nil, 4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, digits := range allDigits {
			LetterCombinations(digits)
		}
	}
}

func BenchmarkLetterCombinationsIterativeAllDigits(b *testing.B) {
	digitsChars := []byte{'2', '3', '4', '5', '6', '7', '8', '9'}
	var allDigits []string
	var dfs func(path []byte, maxLen int)
	dfs = func(path []byte, maxLen int) {
		if len(path) > 0 {
			allDigits = append(allDigits, string(path))
		}
		if len(path) == maxLen {
			return
		}
		for _, ch := range digitsChars {
			path = append(path, ch)
			dfs(path, maxLen)
			path = path[:len(path)-1]
		}
	}
	dfs(nil, 4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, digits := range allDigits {
			LetterCombinationsIterative(digits)
		}
	}
}
