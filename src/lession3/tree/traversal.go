package tree

import "fmt"

func (node *Node) Traverse() {
	node.traverseFunc(func(node *Node) {
		node.Print()
		fmt.Print("++")
	})
	fmt.Println("----")
}

func (node *Node) traverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.traverseFunc(f)
	f(node)
	node.Right.traverseFunc(f)
}


func (node *Node) TraverseWithChannel() chan *Node{
	out := make(chan *Node)
	go func() {
		node.traverseFunc(func(node *Node){
			out <- node
		})
		close(out)
	}()
	return out

}