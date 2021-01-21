package main

type ListNode struct {
    Val int
    Next *ListNode
}

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
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	head = reverseList(head)
}

//递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	finalHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return finalHead
}

//迭代
func reverseList2(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}