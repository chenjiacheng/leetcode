package linked_list

import (
	"testing"
)

func TestCDLinkedList(t *testing.T) {
	l := NewCDLinkedList()
	if l.IsEmpty() == false {
		t.Log("IsEmpty error")
	}

	l = NewCDLinkedList("a", "B", "c")
	l.InsertElemAtHead(111)
	l.InsertElemAtIndex(222, 2)
	l.DeleteElemAtPoint("B")
	l.DeleteElemAtIndex(2)
	l.PrintElem()
}
