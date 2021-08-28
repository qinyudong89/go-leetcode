package structures

import (
	"testing"
)

func TestLinkedList_Init(t *testing.T) {
	list := new(LinkedList)
	list.addFirst(20)
	list.addFirst(12)
	list.addFirst(78)
	list.Display()
}
