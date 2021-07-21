package main
import (
	"fmt"
)
type Node struct {
	data string
	left *Node
	right *Node
}
func (root *Node)PreOrderTraversal(){
	if root !=nil{
		fmt.Printf("%s ", root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
	return
}
func (root *Node)PostOrderTraversal(){
	if root !=nil{
		root.left.PostOrderTraversal()
		root.right.PostOrderTraversal()
		fmt.Printf("%s ", root.data)
	}
	return
}
func (root *Node)InOrderTraversal(){
	if root !=nil{
		root.left.InOrderTraversal()
		fmt.Printf("%s ", root.data)
		root.right.InOrderTraversal()
	}
	return
}
func main(){
	root := Node{"+", nil, nil}
	root.left = &Node{"a", nil, nil}
	root.right = &Node{"-", nil, nil}
	root.right.left = &Node{"b", nil, nil}
	root.right.right = &Node{"c", nil, nil}
	fmt.Printf("Pre Order Traversal of the given tree is: ")
	root.PreOrderTraversal()
	fmt.Printf("\n")
	fmt.Printf("Post Order Traversal of the given tree is: ")
	root.PostOrderTraversal()
	fmt.Printf("\n")
	fmt.Printf("In Order Traversal of the given tree is: ")
	root.InOrderTraversal()
	fmt.Printf("\n")

}