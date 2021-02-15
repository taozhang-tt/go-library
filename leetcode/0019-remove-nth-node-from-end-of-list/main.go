package main
/*
19. 删除链表的倒数第 N 个结点
	https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
题目描述
	给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
	进阶：你能尝试使用一趟扫描实现吗？
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
	removeNthFromEnd(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	fast, slow := pre, pre
	for n >= 0 {
		fast = fast.Next
		n--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
