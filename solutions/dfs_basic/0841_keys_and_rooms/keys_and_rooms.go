package keysandrooms

// CanVisitAllRooms 钥匙和房间
// 有 n 个房间，起初只有 0 号房间是开着的，其余房间都上锁；
// 每个房间里放着若干把钥匙，可以打开对应的房间，判断最终能否进入所有房间。
// 时间复杂度: O(n+e) 每个房间最多访问一次，每把钥匙最多处理一次  空间复杂度: O(n) visited 数组与递归栈
func CanVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	var dfs func(room int)
	dfs = func(room int) {
		if visited[room] {
			return
		}
		visited[room] = true
		// 拿起房间里所有的钥匙，逐一尝试打开对应的房间
		for _, key := range rooms[room] {
			dfs(key)
		}
	}
	dfs(0)
	// 检查是否还有没进入过的房间
	for _, ok := range visited {
		if !ok {
			return false
		}
	}
	return true
}
