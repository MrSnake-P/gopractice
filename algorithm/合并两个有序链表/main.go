package main

import "fmt"

type ListNode struct {
	Val		int
	Next 	*ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode{
	preHead := &ListNode{}	// 创建一个虚拟头节点

	prev := preHead		// 以指针的形式操作
	for l1 !=nil && l2 != nil {		// 链表1，2都不为空
		if l1.Val <= l2.Val {
			prev.Next = l1		// 当其中一个链表的值较小时，指向此链表中的节点
			l1 = l1.Next		// 向后移动一个节点
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next		// 指针要随着链表的节点移动而移动
	}
	// 结束后，之多还有一个非空链表，将它拼接在后面
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	return preHead.Next
}

func mergeTwoListsDiGUi(l1, l2 *ListNode) *ListNode {
	if l1 == nil {			// 终止条件
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		// 判断哪个节点的值小
		l1.Next = mergeTwoListsDiGUi(l1.Next, l2)	// 递归的去寻找l1下一个节点，与l2的节点那个小
		return l1
	} else {
		l2.Next = mergeTwoListsDiGUi(l1, l2.Next)	// 递归的去寻找l2下一个节点，与l1的节点那个小
		return l2
	}
}

func main() {
	l1 := ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	//a := mergeTwoLists(&l1, &l2)

	b := mergeTwoListsDiGUi(&l1, &l2)

	//fmt.Println(a)
	//fmt.Println(a.Next)
	//fmt.Println(a.Next.Next)

	for {
		//fmt.Println(a)
		fmt.Println(b)
		//a = a.Next
		b = b.Next
		//if a == nil {
		//	break
		//}
		if b == nil {
			break
		}
	}


}
