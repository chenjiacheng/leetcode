package linked_list

import "fmt"

type CSElemType interface{}

type CSLinkedList struct {
	Data CSElemType
	Next *CSLinkedList
}

// NewCSLinkedList 实例化循环单链表
func NewCSLinkedList(e ...CSElemType) *CSLinkedList {
	l := &CSLinkedList{}
	l.Data = nil
	l.Next = l

	for _, v := range e {
		l.InsertElemAtEnd(v)
	}

	return l
}

// IsEmpty 判断是否为空的循环单链表
func (l *CSLinkedList) IsEmpty() bool {
	return l.Next == l
}

// GetLength 循环单链表的长度
func (l *CSLinkedList) GetLength() int {
	p := l.Next
	i := 0
	for p != l {
		i++
		p = p.Next
	}

	return i
}

// InsertElemAtHead 从头部添加元素
func (l *CSLinkedList) InsertElemAtHead(e CSElemType) bool {
	l.Next = &CSLinkedList{
		Data: e,
		Next: l.Next,
	}

	return true
}

// InsertElemAtEnd 从尾部添加元素
func (l *CSLinkedList) InsertElemAtEnd(e CSElemType) bool {
	p := l
	for p.Next != l {
		p = p.Next
	}
	p.Next = &CSLinkedList{
		Data: e,
		Next: l,
	}

	return true
}

// InsertElemAtIndex 在指定位置添加元素
func (l *CSLinkedList) InsertElemAtIndex(e CSElemType, i int) bool {
	if i < 1 {
		return false
	}

	p := l
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p == nil {
		return false
	}

	p.Next = &CSLinkedList{
		Data: e,
		Next: p.Next,
	}

	return true
}

// DeleteElemAtPoint 删除指定元素
func (l *CSLinkedList) DeleteElemAtPoint(e CSElemType) bool {
	p := l
	for p.Next != nil && p.Next.Data != e {
		p = p.Next
	}
	p.Next = p.Next.Next

	return true
}

// DeleteElemAtIndex 删除指定位置的元素
func (l *CSLinkedList) DeleteElemAtIndex(i int) bool {
	if i < 1 {
		return false
	}

	p := l
	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	if p == nil {
		return false
	}
	p.Next = p.Next.Next

	return true
}

// PrintElem 遍历所有元素
func (l *CSLinkedList) PrintElem() {
	p := l
	for p.Next != l {
		p = p.Next
		fmt.Println(p.Data)
	}
}
