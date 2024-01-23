package main

import (
	"fmt"
	"io"
	"os"
)

//************************* BINARY SEARCH TREE ********************************

//every node has two nodes attached to it, left and right, and data attached to that node
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

//every tree has a root node
type BinaryTree struct {
	root *BinaryNode
}

//insert root node or in other nodes
func (bt *BinaryTree) insert(data int64) *BinaryTree {
	if bt.root == nil {
		bt.root = &BinaryNode{nil, nil, data}
	} else {
		bt.root.insert(data)
	}
	return bt
}

//insert data in other nodes
func (bn *BinaryNode) insert(data int64) {
	if bn == nil {
		return
	}
	if data <= bn.data {
		if bn.left == nil {
			bn.left = &BinaryNode{nil, nil, data}
		} else {
			bn.left.insert(data)
		}
	} else {
		if bn.right == nil {
			bn.right = &BinaryNode{nil, nil, data}
		} else {
			bn.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func bst() {
	tree := &BinaryTree{}
	tree.insert(100).
		insert(-20).
		insert(-50)

	print(os.Stdout, tree.root, 0, 'M')
}
