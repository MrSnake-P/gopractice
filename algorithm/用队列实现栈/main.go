package main

type MyStack struct {
	queue []int
}

func Constructor() (s MyStack) {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	n := len(this.queue)
	this.queue = append(this.queue, x)
	// 交换位置，将后入队列的放到前端，作为栈顶
	for ; n > 0; n-- {
		this.queue = append(this.queue, this.queue[0])
		this.queue = this.queue[1:]
	}
}

func (this *MyStack) Pop() int{
	x := this.queue[0]
	this.queue = this.queue[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}


func main() {

}
