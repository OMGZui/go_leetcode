/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 14:10
 */
package main

import "fmt"

type MinStack struct {
	stack []item
}

type item struct {
	min, x int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}
	this.stack = append(this.stack, item{
		min: min,
		x:   x,
	})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].x
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}

func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	s.Push(9)
	s.Push(-4)
	fmt.Println(s)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
	fmt.Println(s)
}
