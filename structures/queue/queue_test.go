package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue(6)
	fmt.Println("Empty:", q.Empty())

	q.Push("a")
	fmt.Println(q)

	q.Push("b")
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)

	q.Pop()
	fmt.Println(q)

	q.Push("c")
	q.Push("d")
	q.Push("e")
	q.Push("f")
	q.Push("g")
	q.Push("h")
	fmt.Println(q)

	q.Push("b")
	fmt.Println("Size:", q.Size())
	fmt.Println("Full:", q.Full())
	q.Print()
	fmt.Println(q)

	q.Push("c")
	fmt.Println(q)

	fmt.Println("Size:", q.Size())
	a := q.Pop()
	fmt.Println("Pop:", a)
	fmt.Println(q)
	fmt.Println("Size:", q.Size())

	q.Push("d")
	q.Push("e")
	q.Push("f")
	fmt.Println("Size:", q.Size())

	fmt.Println(q)

	q.Push("g")
	fmt.Println(q)

	q.Print()
	fmt.Println("Full:", q.Full())
	fmt.Println("Push:", q)

	q.Clear()
	fmt.Println("Clear:", q)
}
