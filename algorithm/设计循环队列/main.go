package main

import "fmt"
type MyCircularQueue struct {
	queue     []int // 固定大小的数组，用于保存元素
	headIndex int   // 队首的索引
	count     int   // 循环队列当前的长度，headIndex-count队尾索引
	capacity  int   // 容量，最多可容纳元素
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k, k),
		headIndex: 0,
		count: 0,
		capacity: k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool{
	if this.count == this.capacity {	// 容量满了，不可插入
		return false
	}
	this.queue[(this.headIndex + this.count) % this.capacity] = value
	this.count += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.count == 0 {
		return false
	}
	this.headIndex = (this.headIndex + 1) % this.capacity
	this.count -= 1
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.count == 0 {
		return -1
	}
	return this.queue[this.headIndex]
}


func (this *MyCircularQueue) Rear() int {
	if this.count == 0{
		return -1
	}
	return this.queue[(this.headIndex + this.count - 1) % this.capacity]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.capacity
}

func main() {
	queue := Constructor(5)	// 建立一个固定的队列
	queue.EnQueue(1)
	fmt.Println(queue.count)
	queue.EnQueue(2)
	fmt.Println(queue)
	queue.DeQueue()
	fmt.Println(queue)

}