package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func main() {
	a := &ListNode{Val: 1, Next: &ListNode{Val:2, Next: &ListNode{Val:3, Next: &ListNode{Val:4, Next: &ListNode{Val:5, Next: nil}}}}}
	b := removeNthFromEnd(a, 2)
	fmt.Println(b.Next.Next)
}
