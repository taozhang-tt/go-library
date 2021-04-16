package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isValidBST(root))
	fmt.Println(isValidBST2(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序遍历以后单调递增
func isValidBST(root *TreeNode) bool {
	values := inOrder(root)
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}

func inOrder(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, inOrder(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inOrder(root.Right)...)
	return ans
}

func isValidBST2(root *TreeNode) bool {
	return recursion(root, math.MinInt64, math.MaxInt64)
}

func recursion(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	return root.Val > min && root.Val < max && recursion(root.Left, min, root.Val) && recursion(root.Right, root.Val, max)
}
