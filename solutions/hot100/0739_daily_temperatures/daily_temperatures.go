package dailytemperatures

// DailyTemperatures 每日温度
// 给定一个整数数组 temperatures，表示每天的温度，返回一个数组 answer，
// 其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。
// 如果气温在这之后都不会升高，在该位置用 0 来代替。
// 时间复杂度:  空间复杂度:
func DailyTemperatures(temperatures []int) []int {
	stack := []int{}
	ret := make([]int, len(temperatures), len(temperatures))

	for i, t := range temperatures {

		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]

			stack = stack[0 : len(stack)-1]
			ret[last] = i - last

		}
		stack = append(stack, i)

	}
	return ret
}
