package main

import "fmt"

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	//       3
	//     /   \
	//    11    4
	//   / \     \
	//  4  -2     1
	tree := BinaryTree{
		Root: &Node{
			Value: 3,
			Left: &Node{
				Value: 11,
				Left: &Node{
					Value: 4,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: -2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 4,
				Left:  nil,
				Right: &Node{
					Value: 1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	var result int

	result = treeSumRecursive(tree.Root)
	fmt.Println(result) // 21

	result = treeSum(tree.Root)
	fmt.Println(result) // 21
}

func treeSumRecursive(node *Node) int {
	if node == nil {
		return 0
	}

	return node.Value + treeSumRecursive(node.Left) + treeSumRecursive(node.Right)
}

func treeSum(node *Node) int {
	result := 0
	stack := []Node{*node}

	for len(stack) > 0 {
		current := stack[len(stack)-1] // get last value of slices
		result += current.Value
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
