/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/29
 * Time: 16:49
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type ListNode5 = leetcode_tool.ListNode

func oddEvenList(head *ListNode5) *ListNode5 {
	if head == nil {
		return nil
	}

	odd := head
	even := head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next

		even.Next = odd.Next
		even = even.Next
	}
	// 将偶数的头结点接入到奇数结点末尾
	odd.Next = evenHead
	return head
}

func main() {
	l1 := leetcode_tool.CreateList([]int{1, 2, 3, 4, 5})
	fmt.Println(leetcode_tool.ShowList(l1))

	fmt.Println(leetcode_tool.ShowList(oddEvenList(l1)))
}
