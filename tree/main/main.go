package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	*tree.Node // Embedding
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	myNode.Print()
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("This method is shadowed.")
}

func main() {
	// var root tree.Node
	// fmt.Println(root)
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	// root.Right.Left.Print()
	// root.print()
	// fmt.Println()

	// root.setValue(100)
	// root.print()

	// pRoot := &root
	// pRoot.print()
	// pRoot.setValue(400)
	// pRoot.print()
	// root.print()
	root.Traverse()
	fmt.Println()

	// myRoot := myTreeNode{&root}
	root.postOrder()
	fmt.Println()
}
