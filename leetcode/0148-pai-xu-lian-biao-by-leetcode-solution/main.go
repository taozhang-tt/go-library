package main

import "fmt"

/**
148. 排序链表
    https://leetcode-cn.com/problems/sort-list/
题目描述：
    给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
进阶：
    你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/
func main() {
	head := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: nil,
					},
				},
			},
		},
	}
	head = sortList(head)
	printList(head)
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}


type ListNode struct {
    Val int
    Next *ListNode
}

//归并排序
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
    fast, slow := head, head
	for fast.Next != nil {
		if fast.Next.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
    }
    h2 := slow.Next
    slow.Next = nil
	
	return mergeList(sortList(head), sortList(h2))
}

func mergeList(h1, h2 *ListNode) *ListNode {
	ret := new(ListNode)
	head := ret
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			head.Next = h1
			h1 = h1.Next
		} else {
			head.Next = h2
			h2 = h2.Next
		}
		head = head.Next
	}
	if h1 != nil {
		head.Next = h1
	} else if h2 != nil {
		head.Next = h2
	}
	return ret.Next
}