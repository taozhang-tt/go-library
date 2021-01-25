package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归写法
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ansL := inorderTraversal(root.Left)
	ans = ansL
	ans = append(ans, root.Val)
	ansR := inorderTraversal(root.Right)
	ans = append(ans, ansR...)
	return ans
}