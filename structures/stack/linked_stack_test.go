package stack

import (
	"fmt"
	"testing"
)

func TestLStack(t *testing.T) {
	s := NewLStack()
	fmt.Println("Empty:", s.Empty())

	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Print()

	a := s.Pop()
	fmt.Println("Pop:", a)
	s.Print()

	b := s.Top()
	fmt.Println("Top:", b)
	s.Print()

	s.Clear()
	s.Print()
}
