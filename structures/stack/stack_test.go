package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	fmt.Println("Empty:", s.Empty())

	s.Push("a")
	s.Push("b")
	s.Push("c")
	fmt.Println("Push:", s)

	a := s.Pop()
	fmt.Println("Pop:", a)
	fmt.Println("Pop:", s)

	b := s.Top()
	fmt.Println("Top:", b)
	fmt.Println("Top:", s)

	s.Clear()
	fmt.Println("Clear:", s)
}
