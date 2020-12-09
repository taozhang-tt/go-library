package main

import "fmt"

/*
100. 相同的树
	https://leetcode-cn.com/problems/same-tree/
题目描述
	给定两个二叉树，编写一个函数来检验它们是否相同。
	如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
示例 1:
  输入:       1         1
			/ \       / \
			2   3     2   3

			[1,2,3],   [1,2,3]
	输出: true
示例 2:
  输入:      1          1
			/           \
			2             2

			[1,2],     [1,null,2]

	输出: false
示例 3:
  输入:       1         1
			/ \       / \
			2   1     1   2

			[1,2,1],   [1,1,2]
	输出: false
*/
func main() {
	p := &TreeNode{
		Val: 1,
	}
	q := &TreeNode{
		Val: 1,
	}
	fmt.Println(isSameTree(p, q))
	fmt.Println(isSameTree2(p, q))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
递归
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/*
迭代
*/
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
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
		queue = append(queue, p.Left, q.Left, p.Right, q.Right)
	}
	return true
}
