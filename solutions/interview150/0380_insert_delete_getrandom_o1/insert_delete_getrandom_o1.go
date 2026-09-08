package insertdeletegetrandomo1

import "math/rand"

// RandomizedSet O(1) 时间插入、删除和获取随机元素
// 实现一个满足以下要求的数据结构：
//   - insert(val)： val 不存在时插入，返回是否插入成功；
//   - remove(val)： val 存在时删除，返回是否删除成功；
//   - getRandom： 等概率随机返回一个现有元素。
//
// 三个操作的平均时间复杂度均为 O(1)。
//
// 实现思路：动态切片存值 + 哈希表记录"值 -> 切片下标"。
//   - 插入：直接追加到切片尾部，哈希表记录其下标；
//   - 删除：用切片最后一个元素覆盖待删元素，然后截断末尾（交换删除），
//     使切片保持紧凑，从而在 getRandom 中可以 O(1) 随机定位；
//   - 获取随机：rand.Intn(len(nums)) 直接按下标取值。
type RandomizedSet struct {
	vals  []int       // 存储所有元素，保持紧凑
	index map[int]int // 值 -> 在 vals 中的下标
}

// Constructor 创建空的 RandomizedSet
func Constructor() RandomizedSet {
	return RandomizedSet{
		vals:  make([]int, 0),
		index: make(map[int]int),
	}
}

// Insert 插入元素，已存在则返回 false
// 时间复杂度: O(1)  空间复杂度: O(1) 均摊
func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.index[val]; ok {
		return false
	}
	s.index[val] = len(s.vals)
	s.vals = append(s.vals, val)
	return true
}

// Remove 删除元素，不存在则返回 false
// 时间复杂度: O(1)
func (s *RandomizedSet) Remove(val int) bool {
	i, ok := s.index[val]
	if !ok {
		return false
	}
	last := len(s.vals) - 1
	lastVal := s.vals[last]
	// 用末尾元素覆盖待删位置，保持切片紧凑
	s.vals[i] = lastVal
	s.index[lastVal] = i
	// 截断末尾，完成删除
	s.vals = s.vals[:last]
	delete(s.index, val)
	return true
}

// GetRandom 等概率随机返回一个元素（保证集合非空时调用）
// 时间复杂度: O(1)
func (s *RandomizedSet) GetRandom() int {
	return s.vals[rand.Intn(len(s.vals))]
}
