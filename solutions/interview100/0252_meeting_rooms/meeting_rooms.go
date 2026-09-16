package meetingrooms

import "sort"

// CanAttendMeetings 会议室
// 给定会议时间区间数组 intervals，判断一个人能否参加所有会议（区间互不重叠）。
// 时间复杂度: O(n log n) 瓶颈在排序  空间复杂度: O(log n) 排序递归栈
func CanAttendMeetings(intervals [][]int) bool {
	// 按开始时间升序排序，之后只需检查相邻区间是否重叠
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		// 当前会议在上一个会议结束前开始，时间冲突
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
