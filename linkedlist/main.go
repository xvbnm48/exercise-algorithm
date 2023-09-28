package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	// make a new head with var name second
	// and current head move to second
	second := l.head
	// n is a value from this func
	// this n move to head as new head
	l.head = n
	// and current head is from n, and connection to next node. Next node is a previous head.
	l.head.next = second
	l.length++
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 49}
	node3 := &node{data: 38}
	node4 := &node{data: 18}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	fmt.Println(myList)
}
