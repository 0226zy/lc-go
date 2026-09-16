package wallsandgates

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

const INF = 1<<31 - 1

// cloneRooms 深拷贝二维网格，避免测试用例之间共享底层数组
func cloneRooms(rooms [][]int) [][]int {
	cp := make([][]int, len(rooms))
	for i := range rooms {
		cp[i] = make([]int, len(rooms[i]))
		copy(cp[i], rooms[i])
	}
	return cp
}

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{
			name: "示例1: 4x4 混合网格",
			rooms: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name:  "示例2: 只有一面墙",
			rooms: [][]int{{-1}},
			want:  [][]int{{-1}},
		},
		{
			name:  "单门单格",
			rooms: [][]int{{0}},
			want:  [][]int{{0}},
		},
		{
			name:  "空房间不可达保持INF",
			rooms: [][]int{{INF}},
			want:  [][]int{{INF}},
		},
		{
			name:  "被墙包围的空房间",
			rooms: [][]int{{0, -1, INF}},
			want:  [][]int{{0, -1, INF}},
		},
		{
			name: "一行通道",
			rooms: [][]int{
				{0, INF, INF, INF},
			},
			want: [][]int{
				{0, 1, 2, 3},
			},
		},
		{
			name: "两个门取最近者",
			rooms: [][]int{
				{0, INF, INF, INF, 0},
			},
			want: [][]int{
				{0, 1, 2, 1, 0},
			},
		},
		{
			name: "全是门",
			rooms: [][]int{
				{0, 0},
				{0, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rooms := cloneRooms(tt.rooms)
			WallsAndGates(rooms)
			if !utils.Equal2DIntSlice(rooms, tt.want) {
				t.Errorf("WallsAndGates 后 rooms = %v, want %v", rooms, tt.want)
			}
		})
	}
}

func BenchmarkWallsAndGates(b *testing.B) {
	// 250x250 网格：左上角一扇门，其余全为空房间
	base := make([][]int, 250)
	for r := range base {
		base[r] = make([]int, 250)
		for c := range base[r] {
			base[r][c] = INF
		}
	}
	base[0][0] = 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rooms := cloneRooms(base)
		WallsAndGates(rooms)
	}
}
