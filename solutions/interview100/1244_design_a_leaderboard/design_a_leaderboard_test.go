package designaleaderboard

import "testing"

// op 表示一次操作：名称 + 参数 + 期望返回值（仅 top 有返回值，-1 表示不检查返回值）
type op struct {
	name string
	args []int
	want int
}

// runOps 按「操作序列 + 参数 + 期望值序列」驱动测试
func runOps(t *testing.T, ops []op) {
	t.Helper()
	lb := Constructor()
	for i, o := range ops {
		switch o.name {
		case "addScore":
			lb.AddScore(o.args[0], o.args[1])
		case "top":
			if got := lb.Top(o.args[0]); got != o.want {
				t.Errorf("第 %d 步 top(%d) = %d, want %d", i, o.args[0], got, o.want)
			}
		case "reset":
			lb.Reset(o.args[0])
		default:
			t.Fatalf("未知操作: %s", o.name)
		}
	}
}

func TestLeaderboard(t *testing.T) {
	tests := []struct {
		name string
		ops  []op
	}{
		{
			// LeetCode 官方示例
			name: "官方示例",
			ops: []op{
				{name: "addScore", args: []int{1, 73}},
				{name: "addScore", args: []int{2, 56}},
				{name: "addScore", args: []int{3, 39}},
				{name: "addScore", args: []int{4, 51}},
				{name: "addScore", args: []int{5, 4}},
				{name: "top", args: []int{1}, want: 73},
				{name: "reset", args: []int{1}},
				{name: "reset", args: []int{2}},
				{name: "addScore", args: []int{2, 51}},
				{name: "top", args: []int{3}, want: 141},
			},
		},
		{
			name: "单个参赛者多次加分",
			ops: []op{
				{name: "addScore", args: []int{10, 20}},
				{name: "top", args: []int{1}, want: 20},
				{name: "addScore", args: []int{10, 30}},
				{name: "top", args: []int{1}, want: 50},
			},
		},
		{
			// 边界：K 等于参赛者总数
			name: "K等于参赛人数",
			ops: []op{
				{name: "addScore", args: []int{1, 5}},
				{name: "addScore", args: []int{2, 15}},
				{name: "addScore", args: []int{3, 10}},
				{name: "top", args: []int{3}, want: 30},
			},
		},
		{
			// 边界：重置后重新加入同一参赛者
			name: "重置后重新加入",
			ops: []op{
				{name: "addScore", args: []int{1, 100}},
				{name: "top", args: []int{1}, want: 100},
				{name: "reset", args: []int{1}},
				{name: "addScore", args: []int{1, 1}},
				{name: "top", args: []int{1}, want: 1},
			},
		},
		{
			// 边界：分数相同，top 应累加多人
			name: "并列分数",
			ops: []op{
				{name: "addScore", args: []int{1, 7}},
				{name: "addScore", args: []int{2, 7}},
				{name: "addScore", args: []int{3, 7}},
				{name: "top", args: []int{2}, want: 14},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runOps(t, tt.ops)
		})
	}
}

func BenchmarkLeaderboard(b *testing.B) {
	lb := Constructor()
	// 预先加入 1000 名参赛者
	for id := 1; id <= 1000; id++ {
		lb.AddScore(id, id%100+1)
	}

	b.Run("Top_K=10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lb.Top(10)
		}
	})
	b.Run("AddScore", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lb.AddScore(1, 1)
		}
	})
	b.Run("Reset_AddScore交替", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lb.Reset(500)
			lb.AddScore(500, 42)
		}
	})
}
