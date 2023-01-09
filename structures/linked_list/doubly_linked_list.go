package linked_list

import "fmt"

type DElemType interface{}

type DLinkedList struct {
	Data  DElemType
	Prior *DLinkedList
	Next  *DLinkedList
}

// NewDLinkedList 实例化双链表
func NewDLinkedList(e ...DElemType) DLinkedList {
	var l DLinkedList

	for _, v := range e {
		l.InsertElemAtEnd(v)
	}

	return l
}

// IsEmpty 判断是否为空的双链表
func (l *DLinkedList) IsEmpty() bool {
	return l.Next == nil
}

// GetLength 双链表的长度
func (l *DLinkedList) GetLength() int {
	p := l.Next
	i := 0
	for p != nil {
		i++
		p = p.Next
	}

	return i
}

// InsertElemAtHead 从头部添加元素
func (l *DLinkedList) InsertElemAtHead(e DElemType) bool {
	l.Next.Prior = &DLinkedList{
		Data:  e,
		Prior: l,
		Next:  l.Next,
	}
	l.Next = l.Next.Prior
	return true
}

// InsertElemAtEnd 从尾部添加元素
func (l *DLinkedList) InsertElemAtEnd(e DElemType) bool {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &DLinkedList{
		Data:  e,
		Prior: p,
		Next:  nil,
	}

	return true
}

// InsertElemAtIndex 在指定位置添加元素
func (l *DLinkedList) InsertElemAtIndex(e DElemType, i int) bool {
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

	p.Next.Prior = &DLinkedList{
		Data:  e,
		Prior: p,
		Next:  p.Next,
	}
	p.Next = p.Next.Prior

	return true
}

// DeleteElemAtPoint 删除指定元素
func (l *DLinkedList) DeleteElemAtPoint(e DElemType) bool {
	p := l
	for p.Next != nil && p.Next.Data != e {
		p = p.Next
	}
	p.Next.Next.Prior = p
	p.Next = p.Next.Next

	return true
}

// DeleteElemAtIndex 删除指定位置的元素
func (l *DLinkedList) DeleteElemAtIndex(i int) bool {
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
	p.Next.Next.Prior = p
	p.Next = p.Next.Next

	return true
}

// PrintElem 遍历所有元素
func (l *DLinkedList) PrintElem() {
	p := l
	for p.Next != nil {
		p = p.Next
		fmt.Println(p.Data)
	}
	fmt.Println("--------")
	for p.Prior != nil {
		fmt.Println(p.Data)
		p = p.Prior
	}
}
