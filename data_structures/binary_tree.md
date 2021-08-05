# Binary Tree

A tree data structure represents a hierarchical dataset, like a graph.

Each node in the tree has a value, as well as links to other nodes in the tree. In a binary tree, each node has at most 2 children. These children are referred to as `left` and `right`.

## Traversing a tree.

There are many ways to iterate or traverse over the nodes in a tree.

## Pre-Order Traversal

Visit the root first, then recursively visit the left, then the right nodes in a tree.

```go
// preorderTraversal will visit the root, left, then right nodes.
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    ints := []int{root.Val}
    
    if root.Left != nil {
        ints = append(ints, preorderTraversal(root.Left)...)
    }
    
    if root.Right != nil {
        ints = append(ints, preorderTraversal(root.Right)...)
    }
    
    return ints
}
```

## In-Order Traversal

Recursively visit the left, then the root, then the right nodes in a tree.

For a binary search tree, we can retrieve data in sorted order with an in-order traversal.

```go
// inorderTraversal will visit the left, root, then right nodes.
func inorderTraversal(root *TreeNode) []int {
    var ints []int

    if root == nil {
        return ints
    }
    
    if root.Left != nil {
        ints = append(ints, inorderTraversal(root.Left)...)
    }
    
    ints = append(ints, root.Val)
    
    if root.Right != nil {
        ints = append(ints, inorderTraversal(root.Right)...)
    }
    
    return ints
}
```

## Post-Order Traversal

Visit the left nodes first, visit the right nodes second, visit the root node last.

```go
// postorderTraversal will visit the left nodes first, the right nodes second, and the root node last.
func postorderTraversal(root *TreeNode) []int {
    var ints []int

    if root == nil {
        return ints
    }
    
    if root.Left != nil {
        ints = append(ints, postorderTraversal(root.Left)...)
    }
    
    if root.Right != nil {
        ints = append(ints, postorderTraversal(root.Right)...)
    }
    
    return append(ints, root.Val)
}
```

Deleting from a tree is always in post-order.

## Level-Order Traversal

Level-order traversal searches the tree level by level. Completely search one level before going to the next.

Typically, level-order traversal is paired with a queue (LIFO).

```go
// levelOrder will retrieve values from each level before proceeding to the next.
func levelOrder(root *TreeNode) [][]int {
    var ints [][]int
    var queue []*TreeNode
    
    // Start the queue with the root, if it's not empty.
    if root != nil {
        queue = append(queue, root)
    }
    
    for len(queue) > 0 {
        var level []int
        
        // cache the size of the queue before starting the loop, since we'll be adding to the queue inside the loop.
        size := len(queue)
        for i := 0; i < size; i++ {
        
            // pop the first element off the queue.
            node := queue[0]
            queue = queue[1:]
            
            // Add the current node's value to the level
            level = append(level, node.Val)
            
            // add left and right to the queue, to be ran on the next iteration of the outer loop.
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        ints = append(ints, level)
    }
    
    return ints
}
```

## Solving Tree Problems Recursively

## Top-Down Recursion

In top-down recursion, we start at the top of the tree, compute some values, then send them to the next level of the tree through recursion. This is also known as preorder traversal.

```go
func topDownDepth(n *Node, depth int) int {
    if n == nil {
        return depth
    }
    
    left := topDownDepth(n.Left, depth + 1)
    ight := topDownDepth(n.Right, depth + 1)
    
    if left > right {
        return left
    }
    return right
}
```

## Bottom-Up Recursion

```go
func bottomUpDepth(n *Node, depth int) int {
    if node == nil {
        return depth
    }
    
    left := bottomUpDepth(n.Left, depth)
    right := bottomUpDepth(n.Right, depth)
    
    if left > right {
        return left + 1
    }
    return right + 1
}
```

