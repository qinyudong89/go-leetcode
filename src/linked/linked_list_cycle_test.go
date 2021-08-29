package linked

import (
	"fmt"
	"testing"
)

func TestHasCycle(t *testing.T) {
	list := &ListNode{
		Val: 10,
	}
	list.AddLast(20)
	list.AddLast(30)
	list.AddLast(20)
	//hasCycled := HasCycle(list)
	hasCycled := HasCycleByMap(list)
	fmt.Println(hasCycled)

	for list != nil {
		fmt.Printf("%+v ->", list.Val)
		list = list.Next
	}
	fmt.Println()
}
