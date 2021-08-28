package linked

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var newHead *ListNode
	for head != nil {
		head, head.Next, newHead = head.Next, newHead, head
	}
	return newHead
}
