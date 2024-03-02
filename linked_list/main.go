package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) append(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *LinkedList) prepend(data int) {
	newNode := Node{data: data, next: l.head}
	l.head = &newNode
}

func (l *LinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)

	ll.prepend(5)

	ll.display()
}
