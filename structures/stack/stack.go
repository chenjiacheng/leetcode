package stack

type SElemType interface{}

type Stack struct {
	store []SElemType
}

func NewStack() *Stack {
	return &Stack{}
}

// Empty 是否为空
func (s *Stack) Empty() bool {
	return len(s.store) == 0
}

// Size 获取长度
func (s *Stack) Size() int {
	return len(s.store)
}

// Clear 清空
func (s *Stack) Clear() {
	s.store = make([]SElemType, 0)
}

// Push 入栈
func (s *Stack) Push(e SElemType) bool {
	s.store = append(s.store, e)
	return true
}

// Pop 出栈
func (s *Stack) Pop() SElemType {
	if len(s.store) == 0 {
		return nil
	}

	e := s.store[len(s.store)-1]
	s.store = s.store[:len(s.store)-1]
	return e
}

// Top 获取栈顶元素
func (s *Stack) Top() SElemType {
	if len(s.store) == 0 {
		return nil
	}

	e := s.store[0]
	return e
}
