package linked_list

import "fmt"

type CDElemType interface{}

type CDLinkedList struct {
	Data  CDElemType
	Prior *CDLinkedList
	Next  *CDLinkedList
}

// NewCDLinkedList 实例化循环双链表
func NewCDLinkedList(e ...CDElemType) *CDLinkedList {
	l := &CDLinkedList{}
	l.Data = nil
	l.Prior = l
	l.Next = l

	for _, v := range e {
		l.InsertElemAtEnd(v)
	}

	return l
}

// IsEmpty 判断是否为空的循环双链表
func (l *CDLinkedList) IsEmpty() bool {
	return l.Prior == l.Next
}

// GetLength 循环双链表的长度
func (l *CDLinkedList) GetLength() int {
	p := l.Next
	i := 0
	for p != l {
		i++
		p = p.Next
	}

	return i
}

// InsertElemAtHead 从头部添加元素
func (l *CDLinkedList) InsertElemAtHead(e CDElemType) bool {
	l.Next.Prior = &CDLinkedList{
		Data:  e,
		Prior: l,
		Next:  l.Next,
	}
	l.Next = l.Next.Prior
	return true
}

// InsertElemAtEnd 从尾部添加元素
func (l *CDLinkedList) InsertElemAtEnd(e CDElemType) bool {
	p := l
	for p.Next != l {
		p = p.Next
	}
	p.Next = &CDLinkedList{
		Data:  e,
		Prior: p,
		Next:  l,
	}

	return true
}

// InsertElemAtIndex 在指定位置添加元素
func (l *CDLinkedList) InsertElemAtIndex(e CDElemType, i int) bool {
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

	p.Next.Prior = &CDLinkedList{
		Data:  e,
		Prior: p,
		Next:  p.Next,
	}
	p.Next = p.Next.Prior

	return true
}

// DeleteElemAtPoint 删除指定元素
func (l *CDLinkedList) DeleteElemAtPoint(e CDElemType) bool {
	p := l
	for p.Next != nil && p.Next.Data != e {
		p = p.Next
	}
	p.Next.Next.Prior = p
	p.Next = p.Next.Next

	return true
}

// DeleteElemAtIndex 删除指定位置的元素
func (l *CDLinkedList) DeleteElemAtIndex(i int) bool {
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
func (l *CDLinkedList) PrintElem() {
	p := l
	for p.Next != l.Prior {
		p = p.Next
		fmt.Println(p.Data)
	}
	fmt.Println("--------")
	for p != l {
		fmt.Println(p.Data)
		p = p.Prior
	}
}
