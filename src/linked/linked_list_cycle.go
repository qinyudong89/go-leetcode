package linked

/**
used fast、slow pointer
*/
func HasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

/**
used map
*/
func HasCycleByMap(head *ListNode) bool {
	myMap := make(map[*ListNode]int)
	for head != nil {
		if myMap[head] != 0 {
			return true
		}
		myMap[head] = 1
		head = head.Next
	}
	return false
}
