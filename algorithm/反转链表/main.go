package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		prev = curr

		curr.Next = prev
		curr = next
	}
	return prev
}

func main() {
	head := ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	a := reverseList(&head)
	fmt.Println(a)
	fmt.Println(a.Next)
	fmt.Println(a.Next.Next)
}
