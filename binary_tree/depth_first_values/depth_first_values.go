package main

import "fmt"

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

	result := depthFirstValues(tree.Root)
	fmt.Println(result)

	result = depthFirstValuesRecursive(tree.Root)
	fmt.Println(result)
}

func depthFirstValues(root *Node) []string {
	result := []string{}
	stack := []Node{*root}

	for len(stack) > 0 {
		current := stack[len(stack)-1] // get last value of slices
		result = append(result, current.Value)
		stack = stack[:len(stack)-1] // pop slices

		if current.Right != nil {
			stack = append(stack, *current.Right)
		}
		if current.Left != nil {
			stack = append(stack, *current.Left)
		}
	}

	return result
}

func depthFirstValuesRecursive(node *Node) []string {
	if node == nil {
		return []string{}
	}

	leftNode := depthFirstValuesRecursive(node.Left)
	rightNode := depthFirstValuesRecursive(node.Right)

	result := []string{node.Value}
	result = append(result, leftNode...)
	result = append(result, rightNode...)

	return result

}
