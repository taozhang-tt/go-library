package main

import "fmt"

/*
938. 二叉搜索树的范围和
	https://leetcode-cn.com/problems/range-sum-of-bst/
题目描述
	给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。
*/
func main() {
	var root = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(root, 7, 15))

	root = &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(root, 6, 10))
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 深搜，递归
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	} else if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}

// 广度优先搜索
func rangeSumBST2(root *TreeNode, low, high int) (sum int) {
    q := []*TreeNode{root}
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node == nil {
            continue
        }
        if node.Val > high {
            q = append(q, node.Left)
        } else if node.Val < low {
            q = append(q, node.Right)
        } else {
            sum += node.Val
            q = append(q, node.Left, node.Right)
        }
    }
    return
}