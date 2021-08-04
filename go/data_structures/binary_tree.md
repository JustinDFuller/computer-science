# Binary Tree

A tree data structure represents a hierarchical dataset, like a graph.

Each node in the tree has a value, as well as links to other nodes in the tree. In a binary tree, each node has at most 2 children. These children are referred to as `left` and `right`.

## Traversing a tree.

There are many ways to iterate or traverse over the nodes in a tree.

## Pre-Order Traversal

Visit the root first, then recursively visit the left, then the right nodes in a tree.

```go
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

## Post-Order Traversal

Visit the left nodes first, visit the right nodes second, visit the root node last.

Deleting from a tree is always in post-order.

