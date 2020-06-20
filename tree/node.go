package tree

import "fmt"

// Node is tree node definition
type Node struct {
	Value       int
	Left, Right *Node
}

// Print prints out the node value
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

// SetValue set node value
func (node *Node) SetValue(value int) {
	node.Value = value
}

// Traverse is in-order traversal
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// CreateNode initialize node with given value
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// func main() {
// 	var root treeNode
// 	// fmt.Println(root)
// 	root = treeNode{value: 3}
// 	root.left = &treeNode{}
// 	root.right = &treeNode{5, nil, nil}
// 	root.right.left = new(treeNode)
// 	root.left.right = createNode(2)

// 	root.right.left.setValue(4)
// 	root.right.left.print()
// 	// root.print()
// 	// fmt.Println()

// 	// root.setValue(100)
// 	// root.print()

// 	// pRoot := &root
// 	// pRoot.print()
// 	// pRoot.setValue(400)
// 	// pRoot.print()
// 	// root.print()
// 	root.traverse()
// }
