package main

import "fmt"

type node struct {
	data int
}

type stack struct {
	nodes []node
	count int
}

func (s *stack) Push(n node) {
	s.nodes = append(s.nodes, n)
	s.count++
}

func (s *stack) Pop() {
	//s.nodes = s.nodes[:s.count-1]
	if s.count > 0 {
		s.count--
		fmt.Println("***pop : ", s.nodes[s.count].data)
		s.nodes = s.nodes[:s.count]
	}
}

func Stack() {
	fmt.Println("**********STACK************")
	n1 := node{3}
	n2 := node{4}
	n3 := node{5}
	n4 := node{6}

	s := stack{}
	s.Push(n1)
	s.Push(n2)
	s.Push(n3)
	s.Push(n4)
	fmt.Println("***pushed nodes : ", s.nodes)

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()

	fmt.Println("***nodes remaining : ", s.nodes)
	fmt.Println("***************************")
}

type queue struct {
	nodes []node
	count int
}

func (q *queue) Push(n node) {
	q.nodes = append(q.nodes, n)
	q.count++
}

func (q *queue) Pop() {
	if q.count > 0 {
		fmt.Println("***pop : ", q.nodes[0].data)
		q.nodes = q.nodes[q.count-(q.count-1):]
		q.count--
	}
}

func Queue() {
	fmt.Println("***********QUEUE***********")
	n1 := node{3}
	n2 := node{4}
	n3 := node{5}
	n4 := node{6}

	q := queue{}
	q.Push(n1)
	q.Push(n2)
	q.Push(n3)
	q.Push(n4)
	fmt.Println("***pushed nodes : ", q.nodes)

	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println("***nodes remaining : ", q.nodes)
	fmt.Println("***************************")
}
