package main

import "fmt"

/**
328. 奇偶链表
	https://leetcode-cn.com/problems/odd-even-linked-list/
题目描述
	给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
	请尝试使用原地算法完成。你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。
示例 1:
	输入: 1->2->3->4->5->NULL
	输出: 1->3->5->2->4->NULL
示例 2:
	输入: 2->1->3->5->6->4->7->NULL 
	输出: 2->3->6->7->1->5->4->NULL
*/
func main() {
	var head *ListNode
	head = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 7,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	head = oddEvenList(head)
	printList(head)
}

type ListNode struct {
    Val int
    Next *ListNode
}
/**
双指针
	先分离奇偶节点，组成两条链表 h1(奇), h2（偶）
	再合并 h1 和 h2
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p1, h2, p2 := head, head.Next, head.Next	//h1 就是 head，省略了～
	for p2 != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = h2
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}