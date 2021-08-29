package linked

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	list2 := new(ListNode)
	list3 := new(ListNode)
	list2.Val = 20
	list2.Next = list3
	list3.Val = 30

	list1 := &ListNode{
		Val:  10,
		Next: list2,
	}

	list1 = ReverseList(list1)
	for list1 != nil {
		fmt.Printf("%+v ->", list1.Val)
		list1 = list1.Next
	}
	fmt.Println()
}
