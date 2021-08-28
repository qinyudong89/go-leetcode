package structures

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	Size int
	head *Node
}

func (list *LinkedList) addFirst(data interface{}) {
	element := new(Node)
	element.data = data
	element.next = list.head
	list.head = element
	list.Size++
}

func (l *LinkedList) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}
	fmt.Println()
}
