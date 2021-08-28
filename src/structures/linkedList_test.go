package structures

import (
	"testing"
)

func TestLinkedList_Init(t *testing.T) {
	list := new(LinkedList)
	list.AddFirst(20)
	list.AddFirst(12)
	list.AddFirst(78)
	list.Display()
}
