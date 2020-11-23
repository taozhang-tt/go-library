package main

import "fmt"

/**
876. 链表的中间结点
	https://leetcode-cn.com/problems/middle-of-the-linked-list/
题目描述：
	给定一个头结点为 head 的非空单链表，返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。
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
					Next: nil,
				},
			},
		},
	}
	head = middleNode(head)
	fmt.Println(head.Val)

	head = &ListNode{
		Val: 1,
		Next: nil,
	}
	head = middleNode(head)
	fmt.Println(head.Val)

	head = &ListNode{
		Val: 1,
		Next: nil,
	}
	head = middleNode2(head)
	fmt.Println(head.Val)
}

type ListNode struct {
    Val int
    Next *ListNode
}

//如果有两个中间节点，就返回第二个
func middleNode(head *ListNode) *ListNode {
	var slow, fast = head, head
	for fast.Next != nil {
		slow = slow.Next
		if fast.Next.Next == nil {
			return slow
		}
		fast = fast.Next.Next
	}
	return slow
}

//如果有两个中间节点，就返回第一个
func middleNode2(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil {
		if fast.Next.Next == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}