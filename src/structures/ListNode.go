package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) AddLast(val int) {
	element := new(ListNode)
	element.Val = val
	if list.Val == 0 {
		list = element
	} else {
		for list.Next != nil {
			list = list.Next
		}
		list.Next = element
	}
}
