package firstuniquenumber

import "testing"

// runOps 按「操作序列 + 参数 + 期望值」驱动测试
// ops 取值为 "show"（ShowFirstUnique）或 "add"（Add）
func runOps(t *testing.T, nums []int, ops []string, args []int, want []int) {
	t.Helper()
	fu := Constructor(nums)
	addIdx, wantIdx := 0, 0
	for _, op := range ops {
		switch op {
		case "show":
			if got := fu.ShowFirstUnique(); got != want[wantIdx] {
				t.Errorf("第 %d 次 ShowFirstUnique() = %d, want %d", wantIdx, got, want[wantIdx])
			}
			wantIdx++
		case "add":
			fu.Add(args[addIdx])
			addIdx++
		}
	}
}

func TestFirstUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		ops  []string
		args []int // add 操作的参数，按顺序
		want []int // show 操作的期望返回值，按顺序
	}{
		{
			// LeetCode 官方示例 1
			"示例1: [2,3,5] 逐步加入重复元素",
			[]int{2, 3, 5},
			[]string{"show", "add", "show", "add", "show", "add", "show"},
			[]int{5, 2, 3},
			[]int{2, 2, 3, -1},
		},
		{
			// LeetCode 官方示例 2
			"示例2: [7,7,7,7,7,7] 初始无唯一数字",
			[]int{7, 7, 7, 7, 7, 7},
			[]string{"show", "add", "add", "add", "add", "add", "show"},
			[]int{7, 3, 3, 7, 17},
			[]int{-1, 17},
		},

		// 边界：单个元素先唯一后失效
		{
			"单元素加入重复后无唯一",
			[]int{809},
			[]string{"show", "add", "show"},
			[]int{809},
			[]int{809, -1},
		},

		// 边界：多次查询不改变结果
		{
			"连续多次查询幂等",
			[]int{1, 2, 2},
			[]string{"show", "show", "show"},
			[]int{},
			[]int{1, 1, 1},
		},

		// 边界：先全部重复，再加入新唯一数字
		{
			"全部重复后加入新数字",
			[]int{4, 4},
			[]string{"show", "add", "show", "add", "show"},
			[]int{9, 9},
			[]int{-1, 9, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runOps(t, tt.nums, tt.ops, tt.args, tt.want)
		})
	}
}

func BenchmarkFirstUnique(b *testing.B) {
	// 构造 50000 个数字（一半重复），模拟 add/show 混合调用
	nums := make([]int, 50000)
	for i := range nums {
		nums[i] = i % 25000
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fu := Constructor(nums)
		fu.ShowFirstUnique()
		fu.Add(100000000)
		fu.ShowFirstUnique()
	}
}
