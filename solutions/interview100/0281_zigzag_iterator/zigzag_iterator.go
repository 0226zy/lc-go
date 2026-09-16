package zigzagiterator

// ZigzagIterator 锯齿迭代器
// 交替返回两个向量中的元素；当一个向量耗尽后，继续返回另一个向量的剩余元素。
// 时间复杂度: Constructor/next/hasnext 均为 O(1)  空间复杂度: O(1)
type ZigzagIterator struct {
	v1, v2 []int // 两个输入向量
	i, j   int   // 当前读取位置
	turn   int   // 轮到谁出元素：0 表示 v1，1 表示 v2
}

// Constructor 用两个向量初始化锯齿迭代器
func Constructor(v1 []int, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{v1: v1, v2: v2}
}

// next 按锯齿顺序返回下一个元素（调用前保证 hasnext 为 true）
func (it *ZigzagIterator) next() int {
	// 轮到 v1 且 v1 未耗尽，或轮到 v2 但 v2 已耗尽 → 从 v1 取
	if (it.turn == 0 && it.i < len(it.v1)) || it.j >= len(it.v2) {
		it.turn = 1
		val := it.v1[it.i]
		it.i++
		return val
	}
	it.turn = 0
	val := it.v2[it.j]
	it.j++
	return val
}

// hasnext 判断是否还有下一个元素
func (it *ZigzagIterator) hasnext() bool {
	return it.i < len(it.v1) || it.j < len(it.v2)
}
