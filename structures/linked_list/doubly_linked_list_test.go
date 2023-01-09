package linked_list

import (
	"testing"
)

func TestDLinkedList(t *testing.T) {
	l := NewDLinkedList()
	if l.IsEmpty() == false {
		t.Log("IsEmpty error")
	}

	l = NewDLinkedList("a", "B", "c")
	l.InsertElemAtHead(111)
	l.InsertElemAtIndex(222, 2)
	l.DeleteElemAtPoint("B")
	l.DeleteElemAtIndex(2)
	l.PrintElem()
}
