/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/29
 * Time: 15:47
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type ListNode4 = leetcode_tool.ListNode

func swapPairsGood(head *ListNode4) *ListNode4 {

	// 新建空头结点指向第一个结点
	p := &ListNode4{}
	p.Next = head
	head = p

	for p.Next != nil && p.Next.Next != nil {
		p1 := p.Next
		p2 := p.Next.Next
		p3 := p.Next.Next.Next

		// 交换
		p1.Next = p3
		p2.Next = p1
		p.Next = p2
		p = p.Next.Next
	}

	return head.Next
}

func main() {
	l1 := leetcode_tool.CreateList([]int{1, 2, 3, 4})
	fmt.Println(leetcode_tool.ShowList(l1))

	fmt.Println(leetcode_tool.ShowList(swapPairsGood(l1)))
}
