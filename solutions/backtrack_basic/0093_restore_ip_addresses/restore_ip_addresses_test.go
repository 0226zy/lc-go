package restoreipaddresses

import (
	"reflect"
	"sort"
	"testing"
)

// normalize 对字符串切片排序，用于忽略 IP 地址结果的顺序差异
func normalize(result []string) []string {
	out := make([]string, len(result))
	copy(out, result)
	sort.Strings(out)
	return out
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		// LeetCode 官方示例
		{"示例1: s=\"25525511135\"", "25525511135",
			[]string{"255.255.11.135", "255.255.111.35"}},
		{"示例2: s=\"0000\"", "0000",
			[]string{"0.0.0.0"}},
		{"示例3: s=\"101023\"", "101023",
			[]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},

		// 边界：长度不足 4，无法切出 4 段
		{"长度不足4", "123", nil},
		// 边界：长度超过 12，必然无解
		{"长度超过12", "1111111111111", nil},
		// 边界：含前导零的组合，0 只能单独成段
		{"含前导零", "010010",
			[]string{"0.10.0.10", "0.100.1.0"}},
		// 边界：全为 0 且长度大于 4，0 不能成多位段，无解
		{"全零超长", "00000", nil},
		// 边界：段值恰好等于 255
		{"段值恰好255", "255255255255",
			[]string{"255.255.255.255"}},
		// 边界：段值超过 255 的分支应被剪枝（如 "525"）
		{"段值超255剪枝", "25525525525",
			[]string{"255.255.255.25"}},
		// 边界：任何切法都有段超过 255，无解
		{"段值全超255无解", "25625625625", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RestoreIpAddresses(tt.s)
			got, want := normalize(got), normalize(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("RestoreIpAddresses(%q) = %v, want %v", tt.s, got, want)
			}
		})
	}
}

func BenchmarkRestoreIpAddresses(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{"全1长度12", "111111111111"},
		{"官方示例1", "25525511135"},
		{"混合长度20", "12345678901234567890"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				RestoreIpAddresses(bm.s)
			}
		})
	}
}
