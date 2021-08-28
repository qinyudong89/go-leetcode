package structures

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	prev *Node
	next *Node
	Val  int
}

//初始化
func (l *LinkedList) Init() {
	newNode := new(Node)
	newNode.prev = nil
	newNode.next = nil
	l.Head = newNode
	l.Tail = newNode
	l.Size = 0
}

//头插法
func (l *LinkedList) addFirst(data int) {
	newNode := new(Node)
	newNode.Val = data

	if l.Size == 0 {
		l.Init()
	} else {
		l.Tail.next = newNode
		l.Tail = newNode
	}
}

func (l *LinkedList) Display() {
	list := l.Head
	for list != nil {
		fmt.Printf("%+v ->", list.Val)
		list = list.next
	}
	fmt.Println()
}
