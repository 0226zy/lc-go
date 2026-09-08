package gasstation

// CanCompleteCircuit 加油站
// 在一条环路上有 n 个加油站，第 i 个加油站有 gas[i] 升汽油。
// 从第 i 个加油站开往第 i+1 个加油站需要消耗 cost[i] 升汽油。
// 汽车油箱容量无限，从某个加油站出发，若初始油量为 0，问是否存在一个出发站
// 能绕环路行驶一周；存在则返回该站下标，否则返回 -1。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	total, tank, start := 0, 0, 0
	for i := 0; i < n; i++ {
		diff := gas[i] - cost[i]
		total += diff // 全程累计净收益，若为负则无论从哪出发都走不完
		tank += diff  // 从当前 start 出发的剩余油量
		if tank < 0 {
			// 从 start..i 中任何一站出发都无法到达 i+1，
			// 只能改从下一站 i+1 重新尝试
			start = i + 1
			tank = 0
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
