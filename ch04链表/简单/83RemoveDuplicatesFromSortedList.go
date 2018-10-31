/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 10:56
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type ListNode2 = leetcode_tool.ListNode

func deleteDuplicates(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}

	tail := head
	for tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}
	return head
}

func main() {
	l := leetcode_tool.CreateList([]int{1, 1, 2, 3, 3})
	fmt.Println(l)
	fmt.Println(leetcode_tool.ShowList(l))

	head := deleteDuplicates(l)
	fmt.Println(head)
	fmt.Println(leetcode_tool.ShowList(head))
}
