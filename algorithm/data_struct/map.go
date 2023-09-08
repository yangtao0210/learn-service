package data_struct

import (
	"sync"
)

// Set 集合结构体 goframe提供了gSet类型
type Set struct {
	m            map[int]struct{}
	len          int
	sync.RWMutex //锁：实现并发安全
}

// InitSet 初始化set
func InitSet(cap int) *Set {
	temp := make(map[int]struct{}, cap)
	return &Set{m: temp}
}

// Add 添加元素
func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = struct{}{}
	s.len = len(s.m)
}

// Remove 删除元素
func (s *Set) Remove(item int) {
	s.Lock()
	defer s.Unlock()
	// set为空
	if s.len == 0 {
		return
	}
	// 从map中删除指定键，若不存在就什么也不做
	// The delete built-in function deletes the element with the specified key
	// (m[key]) from the map. If m is nil or there is no such element, delete
	// is a no-op.
	delete(s.m, item)
	s.len = len(s.m)
}

// IsHave 集合中是否包含
func (s *Set) IsHave(item int) bool {
	s.Lock()
	defer s.Unlock()
	if s.len == 0 {
		return false
	}
	_, ok := s.m[item]
	return ok
}

// Clear 清除集合内元素
func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

// ToSlice 将Set的key转为列表
func (s *Set) ToSlice() []int {
	s.Lock()
	defer s.Unlock()
	list := make([]int, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
