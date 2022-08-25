package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) printList() {
	currNode := l.head
	for {
		fmt.Println(currNode.data)
		if currNode.next == nil {
			break
		}
		currNode = currNode.next
	}
}

func main() {
	myList := LinkedList{}
	node1 := &node{data: 10}
	node2 := &node{data: 50}
	node3 := &node{data: 20}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.printList()
}
