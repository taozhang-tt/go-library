package main

import "fmt"

func main() {
	inorder, postorder := []int{9,3,15,20,7}, []int{9,15,7,20,3}
	ans := buildTree(inorder, postorder)
	fmt.Println(ans)
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	hash := make(map[int]int, len(inorder))
	for idx, value := range inorder {
		hash[value] = idx
	}
	var recursion func(int, int) *TreeNode
	recursion = func (left, right int) (node *TreeNode) {
		if left > right {
			return nil
		}
		value := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		idx := hash[value]
		root := &TreeNode{Val: value,}
		root.Right = recursion(idx+1, right)
		root.Left = recursion(left, idx-1)
		return root
	}
	return recursion(0, len(inorder)-1)
}

