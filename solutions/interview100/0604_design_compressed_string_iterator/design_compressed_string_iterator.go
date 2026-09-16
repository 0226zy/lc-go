package designcompressedstringiterator

// StringIterator 迭代压缩字符串
// 对形如 "L1e2t1C1o1d1e1" 的压缩字符串进行惰性解压迭代，
// 不在内存中展开完整字符串，避免超大重复次数导致内存爆炸。
// 时间复杂度: next/hasNext 均摊 O(1)  空间复杂度: O(1)
type StringIterator struct {
	s   string // 原始压缩字符串
	idx int    // 解析指针，指向下一个待解析的字母
	cur byte   // 当前正在输出的字符
	cnt int    // 当前字符剩余未输出的次数
}

// Constructor 用压缩字符串初始化迭代器
func Constructor(compressedString string) StringIterator {
	return StringIterator{s: compressedString}
}

// Next 返回下一个字符；若无更多字符则返回空格 ' '
func (it *StringIterator) Next() byte {
	if !it.HasNext() {
		return ' '
	}
	it.cnt--
	return it.cur
}

// HasNext 判断是否还有未遍历的字符
func (it *StringIterator) HasNext() bool {
	if it.cnt > 0 {
		return true
	}
	// 当前字符已耗尽，尝试解析下一组「字母 + 数字」
	if it.idx >= len(it.s) {
		return false
	}
	it.cur = it.s[it.idx]
	it.idx++
	// 读出完整的十进制数作为重复次数
	for it.idx < len(it.s) && it.s[it.idx] >= '0' && it.s[it.idx] <= '9' {
		it.cnt = it.cnt*10 + int(it.s[it.idx]-'0')
		it.idx++
	}
	return true
}
