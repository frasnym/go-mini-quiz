package main

import (
	"fmt"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func main() {
	//       a
	//     /   \
	//    b     c
	//   / \     \
	//  d   e     f
	tree := BinaryTree{
		Root: &Node{
			Value: "a",
			Left: &Node{
				Value: "b",
				Left: &Node{
					Value: "d",
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: "e",
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: "c",
				Left:  nil,
				Right: &Node{
					Value: "f",
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	result := breadthFirstValues(tree.Root)
	fmt.Println(result)
}

func breadthFirstValues(node *Node) []string {
	result := []string{}
	stack := []Node{*node}

	for len(stack) > 0 {
		current := stack[len(stack)-1] // get last value of slices
		result = append(result, current.Value)
		stack = stack[:len(stack)-1] // pop slices

		if current.Left != nil {
			stack = append([]Node{*current.Left}, stack...)
		}
		if current.Right != nil {
			stack = append([]Node{*current.Right}, stack...)
		}
	}

	return result
}
