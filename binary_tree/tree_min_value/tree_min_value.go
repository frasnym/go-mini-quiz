package main

import (
	"fmt"
	"math"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	//       5
	//     /   \
	//    11    3
	//   /  \    \
	//  4   15    12
	tree := BinaryTree{
		Root: &Node{
			Value: 5,
			Left: &Node{
				Value: 11,
				Left: &Node{
					Value: 4,
					Left:  nil,
					Right: nil,
				},
				Right: &Node{
					Value: 15,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 3,
				Left:  nil,
				Right: &Node{
					Value: 12,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	var result int

	result = treeMinValue(tree.Root)
	fmt.Println(result) // 3

	result = treeMinValueRecursive(tree.Root)
	fmt.Println(result) // 3
}

func treeMinValue(node *Node) int {
	smallest := math.MaxInt
	stack := []Node{*node}

	for len(stack) > 0 {
		current := stack[len(stack)-1] // get last value of slices
		smallest = int(math.Min(float64(smallest), float64(current.Value)))
		stack = stack[:len(stack)-1] // pop slices

		if current.Right != nil {
			stack = append(stack, *current.Right)
		}
		if current.Left != nil {
			stack = append(stack, *current.Left)
		}
	}

	return smallest
}

func treeMinValueRecursive(node *Node) int {
	if node == nil {
		return math.MaxInt
	}

	smallest := math.Min(float64(node.Value), float64(treeMinValueRecursive(node.Left)))
	smallest = math.Min(smallest, float64(treeMinValueRecursive(node.Right)))

	return int(smallest)
}
