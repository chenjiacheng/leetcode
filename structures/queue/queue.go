package queue

import "fmt"

type QElemType interface{}

type Queue struct {
	store []QElemType
	head  int
	tail  int
	size  int
}

func NewQueue(size int) *Queue {
	q := &Queue{
		store: nil,
		head:  0,
		tail:  0,
		size:  size,
	}
	q.store = make([]QElemType, size)
	return q
}

// Empty 是否为空
func (q *Queue) Empty() bool {
	return q.head == q.tail
}

// Full 是否已满
func (q *Queue) Full() bool {
	return (q.tail+1)%q.size == q.head
}

// Size 获取长度
func (q *Queue) Size() int {
	return (q.tail - q.head + q.size) % q.size
}

// Clear 清空
func (q *Queue) Clear() {
	q.store = make([]QElemType, q.size)
	q.head = 0
	q.tail = 0
}

// Push 入栈
func (q *Queue) Push(e QElemType) bool {
	/*if q.Full() {
		return false
	}*/

	q.store[q.tail] = e
	q.tail = (q.tail + 1) % q.size // 计算队尾的位置

	return true
}

// Pop 出栈
func (q *Queue) Pop() QElemType {
	if q.Empty() {
		return nil
	}

	e := q.store[q.head]
	q.head = (q.head + 1) % q.size // 计算队首的位置
	return e
}

// Print 遍历所有元素
func (q *Queue) Print() {
	size := q.Size()
	if size == 0 {
		fmt.Println("queue empty")
	}
	index := q.head
	for i := 0; i < size; i++ {
		fmt.Println(q.store[index])
		index = (index + 1) % q.size // 计算出下一队首索引
	}
}
