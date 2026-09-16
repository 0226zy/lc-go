package encodeanddecodestrings

import (
	"fmt"
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestEncodeDecode(t *testing.T) {
	tests := []struct {
		name string
		strs []string
	}{
		// LeetCode 官方示例
		{"示例1: 普通单词", []string{"Hello", "World"}},
		{"示例2: 单个空字符串", []string{""}},

		// 边界：空列表
		{"空列表", []string{}},

		// 边界：空字符串与分隔符混淆
		{"多个空字符串", []string{"", "", ""}},
		{"空与非空混合", []string{"", "a", ""}},

		// 边界：内容包含分隔符与特殊字符
		{"内容含分隔符#", []string{"ab#cd", "12#"}},
		{"内容含数字与长度前缀混淆", []string{"5#Hello", "3#abc"}},
		{"内容含空格与换行", []string{"a b", "c\nd", "\t\n "}},
		{"内容含中文等宽字符", []string{"你好", "世界，加油！"}},
		{"内容含全部可见ASCII", []string{"!@#$%^&*()_+-=[]{}|;:'\",.<>/?"}},

		// 典型场景
		{"多单词句子", []string{"the", "quick", "brown", "fox"}},
		{"单字符列表", []string{"a"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()
			encoded := codec.Encode(tt.strs)
			got := codec.Decode(encoded)
			// 空列表与空切片在解码后统一视为长度 0
			if len(tt.strs) == 0 {
				if len(got) != 0 {
					t.Errorf("Decode(Encode(%v)) = %v, want 空列表", tt.strs, got)
				}
				return
			}
			if !utils.EqualStringSlice(got, tt.strs) {
				t.Errorf("Decode(Encode(%v)) = %v, want %v", tt.strs, got, tt.strs)
			}
		})
	}
}

func TestEncodeFormat(t *testing.T) {
	// 验证编码格式：长度前缀 + '#' + 内容
	codec := Constructor()
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"普通编码", []string{"Hello", "World"}, "5#Hello5#World"},
		{"空字符串编码", []string{""}, "0#"},
		{"空列表编码为空串", []string{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := codec.Encode(tt.strs); got != tt.want {
				t.Errorf("Encode(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}

func BenchmarkEncodeDecode(b *testing.B) {
	// 构造 200 个长度为 200 的字符串，达到题目数据上限
	strs := make([]string, 200)
	content := make([]byte, 200)
	for i := range content {
		content[i] = byte('a' + i%26)
	}
	for i := range strs {
		strs[i] = fmt.Sprintf("%d#%s", i, string(content))
	}

	codec := Constructor()
	b.Run("编码200x200", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			codec.Encode(strs)
		}
	})
	encoded := codec.Encode(strs)
	b.Run("解码200x200", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			codec.Decode(encoded)
		}
	})
}
