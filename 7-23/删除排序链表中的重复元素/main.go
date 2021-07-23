package main

import "fmt"

type ListNode struct {
	Val		int
	Next 	*ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	prev := head

	for prev.Next != nil {
		if prev.Next.Val == prev.Val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	a := deleteDuplicates(&l1)
	//fmt.Println(a)
	for {
		fmt.Println(a)
		a = a.Next
		if a == nil {
			break
		}
	}
}