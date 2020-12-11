package main

import (
	"fmt"
	"math/rand"
)

/*
108. 将有序数组转换为二叉搜索树(简单)
	https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
题目描述
	将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
	本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
示例:
	给定有序数组: [-10,-3,0,5,9],
	一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
	   0
	  / \
	-3   9
	/   /
   -10  5
*/
func main() {
	nums := []int{-10,-3,0,5,9}
	ans := sortedArrayToBST(nums)
	fmt.Println(ans)
	ans = sortedArrayToBST2(nums)
	fmt.Println(ans)
	ans = sortedArrayToBST3(nums)
	fmt.Println(ans)
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/*
递归，中序遍历（总是选择中间位置的元素作为根节点）
*/
func sortedArrayToBST(nums []int) *TreeNode {
	var recursion func(int, int) *TreeNode
	recursion = func (left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right)/2
		value := nums[mid]
		root := &TreeNode{Val: value}
		root.Left = recursion(left, mid-1)
		root.Right = recursion(mid+1, right)
		return root
	}
	return recursion(0, len(nums)-1)
}

/*
递归，中序遍历（总是选择中间位置右侧的元素作为根节点）
*/
func sortedArrayToBST2(nums []int) *TreeNode {
	var recursion func(int, int) *TreeNode
	recursion = func (left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right+1)/2
		value := nums[mid]
		root := &TreeNode{Val: value}
		root.Left = recursion(left, mid-1)
		root.Right = recursion(mid+1, right)
		return root
	}
	return recursion(0, len(nums)-1)
}

/*
递归，中序遍历（总是选择任意一个中间位置的元素作为根节点）
*/
func sortedArrayToBST3(nums []int) *TreeNode {
	var recursion func(int, int) *TreeNode
	recursion = func (left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right+ rand.Intn(2))/2
		value := nums[mid]
		root := &TreeNode{Val: value}
		root.Left = recursion(left, mid-1)
		root.Right = recursion(mid+1, right)
		return root
	}
	return recursion(0, len(nums)-1)
}
