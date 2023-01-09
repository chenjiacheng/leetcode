package linked_list

import "fmt"

type SElemType interface{}

type SLinkedList struct {
	Data SElemType
	Next *SLinkedList
}

// NewSLinkedList 实例化单链表
func NewSLinkedList(e ...SElemType) SLinkedList {
	var l SLinkedList

	for _, v := range e {
		l.InsertElemAtEnd(v)
	}

	return l
}

// IsEmpty 判断是否为空的单链表
func (l *SLinkedList) IsEmpty() bool {
	return l.Next == nil
}

// GetLength 单链表的长度
func (l *SLinkedList) GetLength() int {
	p := l.Next
	i := 0
	for p != nil {
		i++
		p = p.Next
	}

	return i
}

// InsertElemAtHead 从头部添加元素
func (l *SLinkedList) InsertElemAtHead(e SElemType) bool {
	l.Next = &SLinkedList{
		Data: e,
		Next: l.Next,
	}

	return true
}

// InsertElemAtEnd 从尾部添加元素
func (l *SLinkedList) InsertElemAtEnd(e SElemType) bool {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &SLinkedList{
		Data: e,
		Next: nil,
	}

	return true
}

// InsertElemAtIndex 在指定位置添加元素
func (l *SLinkedList) InsertElemAtIndex(e SElemType, i int) bool {
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

	p.Next = &SLinkedList{
		Data: e,
		Next: p.Next,
	}

	return true
}

// DeleteElemAtPoint 删除指定元素
func (l *SLinkedList) DeleteElemAtPoint(e SElemType) bool {
	p := l
	for p.Next != nil && p.Next.Data != e {
		p = p.Next
	}
	p.Next = p.Next.Next

	return true
}

// DeleteElemAtIndex 删除指定位置的元素
func (l *SLinkedList) DeleteElemAtIndex(i int) bool {
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
func (l *SLinkedList) PrintElem() {
	p := l
	for p.Next != nil {
		p = p.Next
		fmt.Println(p.Data)
	}
}
