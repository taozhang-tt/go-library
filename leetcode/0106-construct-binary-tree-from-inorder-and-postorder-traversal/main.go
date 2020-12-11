package main
/*
106. 从中序与后序遍历序列构造二叉树
	https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
题目描述
	根据一棵树的中序遍历与后序遍历构造二叉树。
注意:
	你可以假设树中没有重复的元素。

例如，给出
	中序遍历 inorder = [9,3,15,20,7]
	后序遍历 postorder = [9,15,7,20,3]
	返回如下的二叉树：

	  3
	 / \
	9  20
	/  \
  15   7
*/
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

/*
递归
	思想：后续遍历的最后一个元素是整棵树的根节点（想一想后续遍历的顺序），递归这一思想，那后续遍历的倒数第二个节点应该是右子树的根节点，倒数第三个是右子树的右子树的根节点（如果存在的话）
		如果不存在，拿它应该是右子树的左子树的根节点（如果存在的话），如果右子树的左孩子不存在的话，那倒数第三个节点，就是整个树的左子树的根节点，这是由后续遍历的顺序决定的
	过程：取后续遍历 postorder 最后一个节点，并查找该节点在中序遍历 inorder 的位置 idx，将 inorder 分成两部分 [0:idx] 和 [idx+1:]，那第一部分存储的即是左子树的所有节点，第二部分存储的是右子树所有的节点
		postorder 弹出最后一个节点，递归思想，先构建右子树，[idx+1:] 即是右子树的中序遍历结果
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	hash := make(map[int]int, len(inorder))	//hash 加速查找
	for idx, value := range inorder {
		hash[value] = idx
	}
	var recursion func(int, int) *TreeNode
	recursion = func (left, right int) (node *TreeNode) {	//递归构建
		if left > right {
			return nil
		}
		value := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		idx := hash[value]
		root := &TreeNode{Val: value,}
		root.Right = recursion(idx+1, right)	//一定要先构建右子树，将后序遍历的思想倒放一遍
		root.Left = recursion(left, idx-1)
		return root
	}
	return recursion(0, len(inorder)-1)
}

