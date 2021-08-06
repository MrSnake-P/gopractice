package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var array []int

	for ; head != nil; head = head.Next {
		array = append(array, head.Val)
	}

	n := len(array)
	for i, v := range array[:n/2] {
		if v != array[n-i-1] {
			return false
		}
	}

	return true
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	a := isPalindrome(&l1)
	fmt.Println(a)
}
