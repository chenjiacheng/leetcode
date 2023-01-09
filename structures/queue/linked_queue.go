package queue

import "fmt"

type LQElemType interface{}

type node struct {
	data LQElemType
	next *node
}

type LQueue struct {
	head *node // 头指针
	tail *node // 尾指针
}

func NewLQueue() *LQueue {
	n := &node{
		data: nil,
		next: nil,
	}
	return &LQueue{
		head: n,
		tail: n,
	}
}

// Empty 是否为空
func (q *LQueue) Empty() bool {
	return q.head == q.tail
}

// Size 获取长度
func (q *LQueue) Size() int {
	i := 0
	p := q.head
	for p != nil {
		i++
		p = p.next
	}

	return i
}

// Clear 清空
func (q *LQueue) Clear() {
	n := &node{
		data: nil,
		next: nil,
	}
	q.head = n
	q.tail = n
}

// Push 入栈
func (q *LQueue) Push(e LQElemType) bool {
	p := &node{e, nil}
	q.tail.next = p
	q.tail = p
	return true
}

// Pop 出栈
func (q *LQueue) Pop() LQElemType {
	if q.Empty() {
		return nil
	}

	p := q.head.next
	e := p.data
	q.head = p.next
	return e
}

// Print 遍历所有元素
func (q *LQueue) Print() {
	p := q.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
