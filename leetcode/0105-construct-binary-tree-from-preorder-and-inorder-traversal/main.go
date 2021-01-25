package main

import "fmt"

/**
105. 从前序与中序遍历序列构造二叉树
	https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
题目描述
	根据一棵树的前序遍历与中序遍历构造二叉树。

	注意:
	你可以假设树中没有重复的元素。
	例如，给出
	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]

	返回如下的二叉树：
	 3
	/ \
	9  20
	  /  \
	 15   7
*/
func main() {
	var preorder, inorder []int
	var root *TreeNode
	preorder = []int{3,9,20,15,7}
	inorder = []int{9,3,15,20,7}
	root = buildTree(preorder, inorder)
	fmt.Println(root)
	preorder = []int{3,9,20,15,7}
	inorder = []int{9,3,15,20,7}
	root = buildTree2(preorder, inorder)
	fmt.Println(root)
	preorder = []int{3,9,20,15,7}
	inorder = []int{9,3,15,20,7}
	root = buildTree3(preorder, inorder)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
遍历先序遍历数组，获取到一个元素 e, 在中序遍历数组[i, j] 中查找该元素，索引值为 k，那么 k 对应的元素及为一个根结点，[i, k-1] 为左子树上的节点，[k+1, j] 为右子树上所有的节点
递归上述过程
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	root, _ := build(preorder, inorder)
	return root
}

func build(preorder []int, inorder []int) (*TreeNode, []int) {
	if len(preorder) == 0 {
		return nil, preorder
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	preorder = preorder[1:]
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			if i > 0 {
				root.Left, preorder = build(preorder, inorder[0:i])
			}
			if i < len(inorder)-1 {
				root.Right, preorder = build(preorder, inorder[i+1:])
			}
			break
		}
	}
	return root, preorder
}

/**
官方题解
	思路一样，但是优化了
*/
func buildTree2(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root := &TreeNode{preorder[0], nil, nil}
    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            break
        }
    }
    root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])	//这里很巧妙呀
    root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
    return root
}

func buildTree3(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}
	idx := 0
	for ; idx < len(inorder); idx++ {	//在中序遍历里找到对应的元素，以该元素分界，左侧的元素inorder[:idx]属于左子树，右侧的元素inorder[idx+1:]属于右子树
		if inorder[idx] == preorder[0] {
			break
		}
	}
	root.Left = buildTree3(preorder[1:len(inorder[:idx])+1], inorder[:idx])	//在先序遍历中，左子树的元素优先于右子树的元素被访问，所以确定了左子树元素的个数inorder[:idx]，那么先序遍历的前 len(inorder[:idx]) 个元素都属于左子树
	root.Right = buildTree3(preorder[len(inorder[:idx])+1:], inorder[idx+1:])
	return root
}