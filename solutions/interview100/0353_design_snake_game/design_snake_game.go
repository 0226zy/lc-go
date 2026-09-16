package designsnakegame

// SnakeGame 设计贪吃蛇
// 蛇从 (0, 0) 出发，按方向移动；吃到食物得分并变长，越界或咬到自己游戏结束。
// 时间复杂度: 每次 move O(1)  空间复杂度: O(width * height)
type SnakeGame struct {
	width    int
	height   int
	food     [][]int       // 食物位置，按顺序出现
	foodIdx  int           // 下一份食物的下标
	body     [][2]int      // 蛇身双端队列：队首为蛇头，队尾为蛇尾
	occupied map[int]bool  // 蛇身占据的格子集合，键为 r*width+c
	score    int           // 当前得分（已吃掉的食物数）
	gameOver bool          // 游戏是否已结束
}

// Constructor 初始化 width x height 的贪吃蛇游戏，食物按 food 顺序出现
func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		width:    width,
		height:   height,
		food:     food,
		body:     [][2]int{{0, 0}},
		occupied: map[int]bool{0: true},
	}
}

// Move 蛇向 direction 方向移动一格
// 返回值：游戏结束返回 -1，否则返回当前得分
func (s *SnakeGame) Move(direction string) int {
	if s.gameOver {
		return -1
	}

	// 根据方向计算蛇头新坐标
	head := s.body[0]
	nr, nc := head[0], head[1]
	switch direction {
	case "U":
		nr--
	case "D":
		nr++
	case "L":
		nc--
	case "R":
		nc++
	}

	// 越界：游戏结束
	if nr < 0 || nr >= s.height || nc < 0 || nc >= s.width {
		s.gameOver = true
		return -1
	}

	// 判断新格子是否是下一份食物
	eat := s.foodIdx < len(s.food) && s.food[s.foodIdx][0] == nr && s.food[s.foodIdx][1] == nc
	if eat {
		s.score++
		s.foodIdx++
	} else {
		// 不吃食物时蛇尾立刻让出，先移除再判撞身，避免误判「跟着尾巴走」
		tail := s.body[len(s.body)-1]
		delete(s.occupied, tail[0]*s.width+tail[1])
		s.body = s.body[:len(s.body)-1]
	}

	// 咬到自己：游戏结束
	if s.occupied[nr*s.width+nc] {
		s.gameOver = true
		return -1
	}

	// 蛇头入队
	s.body = append([][2]int{{nr, nc}}, s.body...)
	s.occupied[nr*s.width+nc] = true
	return s.score
}
