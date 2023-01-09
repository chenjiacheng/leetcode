package stack

import "fmt"

type LSElemType interface{}

type node struct {
	data LSElemType
	next *node
}

type LStack struct {
	top   *node // 头指针
	count int
}

func NewLStack() *LStack {
	return &LStack{}
}

// Empty 是否为空
func (s *LStack) Empty() bool {
	return s.count == 0
}

// Size 获取长度
func (s *LStack) Size() int {
	return s.count
}

// Clear 清空
func (s *LStack) Clear() {
	s.count = 0
	s.top = nil
}

// Push 入栈
func (s *LStack) Push(e LSElemType) bool {
	s.top = &node{e, s.top}
	s.count++
	return true
}

// Pop 出栈
func (s *LStack) Pop() LSElemType {
	if s.count == 0 {
		return nil
	}

	e := s.top.data
	s.top = s.top.next
	return e
}

// Top 获取栈顶元素
func (s *LStack) Top() LSElemType {
	if s.count == 0 {
		return nil
	}

	e := s.top.data
	return e
}

// Print 遍历所有元素
func (s *LStack) Print() {
	p := s.top
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
