package main

import "fmt"

/*
101. 对称二叉树
	https://leetcode-cn.com/problems/symmetric-tree/
题目描述
	给定一个二叉树，检查它是否是镜像对称的。
	例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
*/
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isSymmetric(root))

	fmt.Println(isSymmetric2(root))
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/*
递归
	两个指针 p, q 分别指向两棵子树的根节点，
	判断两棵子树是否对称，
		首先判断两个根节点的值是否相等，
		然后再递归判断 p 的左子树是否和 q 的右子树对称，p 的右子树是否和 q 的左子树对称
*/
func isSymmetric(root *TreeNode) bool {
    return recursion(root, root)
}

func recursion(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && recursion(p.Left, q.Right) && recursion(p.Right, q.Left)
}

/*
迭代
维护一个队列，初始时先将根节点的左节点、右节点依次入队列
每次从队列里取两个元素 p, q，两个元素对应着两棵子树的根，如果这两棵子树要对称，那根节点的值必须相等
将 p.Left、q.Right、p.Right、q.Left 依次入队
*/
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}