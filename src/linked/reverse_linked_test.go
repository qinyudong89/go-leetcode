package linked

import (
	"testing"
)

func TestReverse(t *testing.T) {
	list2 := &ListNode{
		Val: 20,
		Next: &ListNode{
			Val:  30,
			Next: nil,
		},
	}
	list1 := &ListNode{
		Val:  10,
		Next: list2,
	}
	ReverseList(list1)
}
