package keysandrooms

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		// LeetCode 官方示例
		{
			"示例1: 链式拿钥匙可进所有房间",
			[][]int{{1}, {2}, {3}, {}},
			true,
		},
		{
			"示例2: 3号房间钥匙拿不到",
			[][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			false,
		},

		// 边界情况
		{"只有一个房间", [][]int{{}}, true},
		{"两个房间可互达", [][]int{{1}, {0}}, true},
		{"0号房间自循环", [][]int{{0, 1}, {}}, true},
		{"0号房间没有任何钥匙", [][]int{{}, {0}}, false},
		{"最后一间房的钥匙在自己房里", [][]int{{1}, {}, {2}}, false},
		{"所有钥匙都指向已访问的房间", [][]int{{1, 2}, {1, 2}, {1}}, true},

		// 链式递归：钥匙顺序与房间编号无关
		{
			"逆序链式可达",
			[][]int{{2}, {0}, {1}},
			true,
		},
		{
			"孤立房间拿不到钥匙",
			[][]int{{1}, {3}, {2}, {}},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("CanVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

// makeChainRooms 生成 n 个房间的链式布局：房间 i 里放着 i+1 的钥匙（最后一个房间为空）
func makeChainRooms(n int) [][]int {
	rooms := make([][]int, n)
	for i := 0; i < n-1; i++ {
		rooms[i] = []int{i + 1}
	}
	return rooms
}

func BenchmarkCanVisitAllRooms(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"100个房间链式", 100},
		{"1000个房间链式", 1000},
	}

	for _, bm := range benchmarks {
		rooms := makeChainRooms(bm.n)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanVisitAllRooms(rooms)
			}
		})
	}
}
