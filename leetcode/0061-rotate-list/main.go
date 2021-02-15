package main

/*
61. 旋转链表
	https://leetcode-cn.com/problems/rotate-list/
题目描述
	给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
示例 1:
		输入: 1->2->3->4->5->NULL, k = 2
		输出: 4->5->1->2->3->NULL
		解释:
		向右旋转 1 步: 5->1->2->3->4->NULL
		向右旋转 2 步: 4->5->1->2->3->NULL
*/
func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	// rotateRight(head, 2)

	head = &ListNode{
		Val: 1,
	}
	rotateRight(head, 1)
}

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	cnt := 0
	temp := head
	for temp != nil {
		cnt++
		temp = temp.Next
	}
	k %= cnt
	if k == 0 {
		return head
	}
	pre := &ListNode{
		Next: head,
	}
	fast, slow := pre, pre
	for (k>0) {
		fast = fast.Next
		k--
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next, head, slow.Next = head, slow.Next, nil
	return head
}