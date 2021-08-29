package linked

import "github.com/qinyudong89/go-leetcode/src/structures"

type ListNode = structures.ListNode

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
