package structures

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Size int
	Head *Node
}

func (list *LinkedList) AddFirst(data interface{}) {
	element := new(Node)
	element.Data = data
	element.Next = list.Head
	list.Head = element
	list.Size++
}

func (l *LinkedList) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v ->", list.Data)
		list = list.Next
	}
	fmt.Println()
}
