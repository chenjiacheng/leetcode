package queue

import (
	"fmt"
	"testing"
)

func TestLQueue(t *testing.T) {
	q := NewLQueue()
	fmt.Println("Empty:", q.Empty())

	q.Push("a")
	q.Push("b")
	q.Push("c")
	q.Print()

	a := q.Pop()
	fmt.Println("Pop:", a)
	q.Print()

	q.Clear()
	q.Print()
}
