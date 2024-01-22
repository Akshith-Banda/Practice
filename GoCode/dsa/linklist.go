package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(i int) {
	//if head node is nil then add present node to it
	node := &Node{data: i}
	if l.head == nil {
		l.head = node
		return
	}

	//move to next node and check if it's nil
	curr := l.head
	for curr.next != nil { //if not move to next node
		curr = curr.next
	}
	//if node is nil add present node to it
	curr.next = node
}

func (l *List) remove(i int) {
	//if head is nil return
	if l.head == nil {
		return
	}

	//if not comapre value of head to i
	if l.head.data == i { //if matches removes that node and replace it with next node
		l.head = l.head.next
	}

	//else move to next node
	curr := l.head
	for curr.next != nil && curr.next.data != i {
		curr = curr.next
	}

	//if data matches replace that node with next node
	if curr.next != nil {
		curr.next = curr.next.next
	}

}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d\n", curr.data)
		curr = curr.next
	}
}

func linklist() {
	list := new(List)
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	printList(list)
	fmt.Println("removing 2")
	list.remove(2)
	printList(list)
}
