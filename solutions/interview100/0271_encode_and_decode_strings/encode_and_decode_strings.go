package encodeanddecodestrings

import (
	"strconv"
	"strings"
)

// Codec 字符串编解码器
// 设计一个算法，将字符串列表编码为单个字符串，并能解码回原来的列表。
// 采用「长度前缀」编码：每个字符串编码为 "长度#内容"，可处理包含任意字符的字符串。
type Codec struct{}

// Constructor 创建一个编解码器实例
func Constructor() Codec {
	return Codec{}
}

// Encode 将字符串列表编码为单个字符串
// 时间复杂度: O(L) L 为所有字符串总长度  空间复杂度: O(L)
func (c *Codec) Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		// 写入长度前缀 + 分隔符 + 内容
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

// Decode 将编码后的字符串解码回原来的字符串列表
// 时间复杂度: O(L)  空间复杂度: O(L)
func (c *Codec) Decode(s string) []string {
	var result []string
	i := 0
	for i < len(s) {
		// 找到分隔符位置，解析出字符串长度
		j := strings.IndexByte(s[i:], '#')
		j += i
		l, _ := strconv.Atoi(s[i:j])
		// 按长度精确截取内容，内容可包含任意字符
		result = append(result, s[j+1:j+1+l])
		i = j + 1 + l
	}
	return result
}
