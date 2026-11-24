package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Height returns the number of nodes on the longest root-to-leaf path
// (0 for an empty tree): one more than the taller subtree.
func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := Height(root.Left), Height(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	fmt.Println(Height(root))
}
