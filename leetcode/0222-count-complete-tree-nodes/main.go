package main

import (
	"fmt"
	"math"
	"sort"
)

/*
222. 完全二叉树的节点个数
	https://leetcode-cn.com/problems/count-complete-tree-nodes/
题目描述
	给出一个完全二叉树，求出该树的节点个数。

说明：
	完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。
示例:
	输入:
		1
	/ \
	2   3
	/ \  /
	4  5 6

	输出: 6
*/
func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(countNodes(root))
	fmt.Println(countNodes2(root))
	fmt.Println(countNodes3(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
层序遍历，暴力解法
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count, queue := 0, []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr != nil {
			count++
			queue = append(queue, curr.Left, curr.Right)
		}
	}
	return count
}

/*
深度优先遍历先求得最大深度
后序遍历求得最后一层有多少节点
*/
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	head, maxDep := root, 0
	for head != nil {
		maxDep++
		head = head.Left
	}
	count := depth(root, 0, maxDep)
	return int(math.Pow(float64(2), float64(maxDep-1))) - 1 + count
}

func depth(root *TreeNode, currDep, maxDep int) (count int) {
	currDep++
	if root.Left == nil && currDep == maxDep {
		return 1
	}
	if root.Left != nil {
		count += depth(root.Left, currDep, maxDep)
	}
	if root.Right != nil {
		count += depth(root.Right, currDep, maxDep)
	}
	return count
}

/*
官方题解: 二分查找 + 位运算
*/
func countNodes3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}
