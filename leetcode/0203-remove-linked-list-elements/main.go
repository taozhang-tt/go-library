package main

import "fmt"

func main() {
	// head := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 6,
	// 			Next: &ListNode{
	// 				Val: 3,
	// 				Next: &ListNode{
	// 					Val: 4,
	// 					Next: &ListNode{
	// 						Val: 5,
	// 						Next: &ListNode{
	// 							Val: 6,
	// 							Next: nil,
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// head = removeElements(head, 6)
	// printList(head)

	head := &ListNode{
		Val: 1,
		Next: nil,
	}
	head = removeElements(head, 1)
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

func removeElements(head *ListNode, val int) *ListNode {
	p:= new(ListNode)
	p.Next = head
	q := p
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return q.Next
}