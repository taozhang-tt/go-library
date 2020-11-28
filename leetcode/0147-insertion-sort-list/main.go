package main

import "fmt"

/*
147. 对链表进行插入排序
	https://leetcode-cn.com/problems/insertion-sort-list/
题目描述：
	插入排序算法：
	插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
	每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
	重复直到所有输入数据插入完为止。
示例 1：
	输入: 4->2->1->3
	输出: 1->2->3->4
示例 2：
	输入: -1->5->3->4->0
	输出: -1->0->3->4->5
*/
func main() {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	head = insertionSortList(head)
	printList(head)

	head = &ListNode{
		Val:  4,
		Next: nil,
	}
	head = insertionSortList(head)
	printList(head)

	head = nil
	head = insertionSortList(head)
	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	virtual := &ListNode{Next: head} //虚拟头指针
	pre, curr := head, head.Next     //默认第一个元素已经有序
	for curr != nil {
		if curr.Val < pre.Val { //如果比它的上一个节点大，则已经有序
			pre.Next = curr.Next //从链表中移除待排序节点
			p := virtual
			for p.Next != nil && p.Next.Val < curr.Val {
				p = p.Next
			}
			curr.Next = p.Next
			p.Next = curr
		} else {
			pre = pre.Next
		}
		curr = pre.Next
	}
	return virtual.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
