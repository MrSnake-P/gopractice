package main

import "fmt"

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) inOut() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1]) // 压入输入栈最后一个元素到输出栈
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {		// 如何输出栈里面没有元素，那么将输入栈的压入过来
		this.inOut()
	}
	x := this.outStack[len(this.outStack) - 1]
	this.outStack = this.outStack[:len(this.outStack) - 1]
	return x
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		this.inOut()
	}
	return this.outStack[len(this.outStack) - 1]
}

func (this *MyQueue) Empty() bool {
	return len(this.outStack) == 0 && len(this.inStack) == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	queue.Pop()
	fmt.Println(queue.Peek())
}