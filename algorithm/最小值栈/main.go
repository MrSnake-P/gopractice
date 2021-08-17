package main

import "math"

type MinStack struct {
	Stack    []int
	MinStack []int // 辅助栈，存储最小值
}

func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		MinStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	// 首先压入栈
	this.Stack = append(this.Stack, x)
	// 再把较小的值压入栈
	top := this.MinStack[len(this.Stack)-1] // 取辅助栈顶元素
	this.MinStack = append(this.MinStack, Min(x, top))
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
