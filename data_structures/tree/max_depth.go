// Write an O(n) algorithm to find the maximum depth of a tree
package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func MaxDepth(n *Node) int {
	if n == nil {
		return 0
	}

	l := MaxDepth(n.left)
	r := MaxDepth(n.right)

	if l > r {
		return l + 1
	}

	return r + 1
}

func main() {
	n := Node{
		val: 2,
		left: &Node{
			val: 1,
		},
		right: &Node{
			val: 3,
			left: &Node{
				val: 4,
			},
		},
	}
	fmt.Println(MaxDepth(&n))
}
