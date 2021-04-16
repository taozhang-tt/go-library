package main

import "fmt"

/*
236. 二叉树的最近公共祖先
	给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/
func main() {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 8,
		},
	}
	fmt.Println(lowestCommonAncestor(root, &TreeNode{Val: 2}, &TreeNode{Val: 8}))
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}