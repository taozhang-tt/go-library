package main

import "fmt"

/*
103. 二叉树的锯齿形层序遍历
	https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
题目描述：
	给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
例如：
	给定二叉树 [3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回锯齿形层序遍历如下：
	[
	[3],
	[20,9],
	[15,7]
	]
*/
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
模拟层序遍历
	偶数层时，插入元素顺序从左向右
	奇数层时，插入元素从右向左
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue, level := []*TreeNode{root}, []int{0}
	for len(queue) > 0 {
		currNode, currLevel := queue[0], level[0]
		queue, level = queue[1:], level[1:]
		if currNode == nil {
			continue
		}
		queue = append(queue, currNode.Right, currNode.Left)
		if len(ans) <= currLevel {
			ans = append(ans, []int{})
		}
		if currLevel%2 == 0 {
			ans[currLevel] = append([]int{currNode.Val}, ans[currLevel]...)
		} else {
			ans[currLevel] = append(ans[currLevel], currNode.Val)
		}
		level = append(level, currLevel+1, currLevel+1)
	}
	return ans
}
