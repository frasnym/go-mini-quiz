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

	result := maxPathSumRecursive(tree.Root)
	fmt.Println(result) // 18
}

func maxPathSumRecursive(node *Node) int {
	if node == nil {
		return math.MinInt
	}
	if node.Left == nil && node.Right == nil {
		return node.Value
	}

	maxChildPathSum := math.Max(float64(maxPathSumRecursive(node.Right)), float64(maxPathSumRecursive(node.Left)))

	return int(maxChildPathSum) + node.Value
}
