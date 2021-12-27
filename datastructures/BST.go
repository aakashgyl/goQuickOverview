package main

import "fmt"

type Node struct {
	key int
	lptr *Node
	rptr *Node
}

type BST struct {
	root *Node
}

func (b *BST) insert(data int){
	newNode := &Node{
		key: data,
	}

	if b.root == nil {
		b.root = newNode
		return
	}

	root := b.root

	for {
		if data < root.key {
			if root.lptr == nil {
				root.lptr = newNode
				return
			} else {
				root = root.lptr
			}
		} else {
			if root.rptr == nil {
				root.rptr = newNode
				return
			} else {
				root = root.rptr
			}
		}
	}
}

func (b *BST) inorder(){
	inorder(b.root)
}

func inorder(n *Node){
	if n != nil {
		inorder(n.lptr)
		fmt.Printf("%d -> ", n.key)
		inorder(n.rptr)
	}
}

func main() {
	data := []int{10,8,12,7,14,6,16,4,18,2,20}
	bst := BST{}
	for _, val := range data {
		bst.insert(val)
	}

	bst.inorder()	// if it prints data in sorted order, that means BST creation is working
}
