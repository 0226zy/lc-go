package movingaveragefromdatastream

// MovingAverage 数据流中的移动平均值
// 维护一个容量为 size 的滑动窗口，next 加入新值并返回窗口内所有整数的平均值。
// 时间复杂度: 每次 next O(1)  空间复杂度: O(size)
type MovingAverage struct {
	size  int   // 窗口容量
	queue []int // 窗口内的元素，队首为最老元素
	sum   int   // 窗口内元素之和，避免每次遍历求和
}

// Constructor 用窗口大小 size 初始化 MovingAverage
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

// Next 向数据流中加入 val，返回最近最多 size 个数的平均值
func (ma *MovingAverage) Next(val int) float64 {
	// 窗口已满：弹出最老的元素
	if len(ma.queue) == ma.size {
		ma.sum -= ma.queue[0]
		ma.queue = ma.queue[1:]
	}
	ma.queue = append(ma.queue, val)
	ma.sum += val
	return float64(ma.sum) / float64(len(ma.queue))
}
