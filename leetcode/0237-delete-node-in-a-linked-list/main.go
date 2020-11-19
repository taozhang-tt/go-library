package main

import "fmt"

/**
237. 删除链表中的节点
	https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
题目描述：
	请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。
	现有一个链表 -- head = [4,5,1,9]，它可以表示为:
示例 1：
	输入：head = [4,5,1,9], node = 5
	输出：[4,1,9]
	解释：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2：
	输入：head = [4,5,1,9], node = 1
	输出：[4,5,9]
	解释：给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
提示：
	链表至少包含两个节点。
	链表中所有节点的值都是唯一的。
	给定的节点为非末尾节点并且一定是链表中的一个有效节点。
	不要从你的函数中返回任何结果。
*/
func main() {
	node1 := &ListNode{
		Val:  4,
		Next: nil,
	}
	node2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node3 := &ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  9,
		Next: nil,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	deleteNode(node3)
	printList(node1)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
