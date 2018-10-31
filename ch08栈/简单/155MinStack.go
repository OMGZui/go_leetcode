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

func (minStack *MinStack) Push(x int) {
	min := x
	if len(minStack.stack) > 0 && minStack.GetMin() < x {
		min = minStack.GetMin()
	}
	minStack.stack = append(minStack.stack, item{
		min: min,
		x:   x,
	})
}

func (minStack *MinStack) Pop() {
	l := len(minStack.stack)
	minStack.stack = minStack.stack[:l-1]
}

func (minStack *MinStack) Top() int {
	l := len(minStack.stack)
	return minStack.stack[l-1].x
}

func (minStack *MinStack) GetMin() int {
	l := len(minStack.stack)
	return minStack.stack[l-1].min
}

func main() {
	s := MinStack{}
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
}
