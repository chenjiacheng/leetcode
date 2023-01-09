package linked_list

import (
	"testing"
)

func TestSLinkedList_IsEmpty(t *testing.T) {
	l := NewSLinkedList()
	if l.IsEmpty() == false {
		t.Log("IsEmpty error")
	}

	l = NewSLinkedList("a")
	if l.IsEmpty() == true {
		t.Log("IsEmpty error")
	}
}

func TestSLinkedList_GetLength(t *testing.T) {
	l := NewSLinkedList()
	if l.GetLength() != 0 {
		t.Log("GetLength error")
	}

	l = NewSLinkedList("a")
	if l.GetLength() != 1 {
		t.Log("GetLength error")
	}

	l = NewSLinkedList("a", "b", "c")
	if l.GetLength() != 3 {
		t.Log("GetLength error")
	}
}
