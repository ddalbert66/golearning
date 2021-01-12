package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

//round2
func (treeNode *Node) Print() {
	fmt.Print(treeNode.Value)
}

//hello setval
func (treeNode *Node) SetVal(i int) {
	if treeNode == nil {
		fmt.Println("Setting value to nil")
		return
	}
	treeNode.Value = i
}

func (treeNode *Node) traverse() {
	if treeNode == nil {
		return
	}
	treeNode.Left.traverse()
	treeNode.Print()
	treeNode.Right.traverse()
}

func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
