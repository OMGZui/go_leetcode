/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/23
 * Time: 15:08
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type ListNode3 = leetcode_tool.ListNode

func addTwoNumbers(l1 *ListNode3, l2 *ListNode3) *ListNode3 {
	// 头结点
	head := &ListNode3{}
	//尾结点
	current := head
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		current.Next = &ListNode3{Val: sum % 10}
		current = current.Next
	}
	return head.Next
}

//主函数测试
func main() {
	l1 := leetcode_tool.CreateList([]int{1, 4, 3})
	fmt.Println(leetcode_tool.ShowList(l1))

	l2 := leetcode_tool.CreateList([]int{5, 6, 4})
	fmt.Println(leetcode_tool.ShowList(l2))

	merge := addTwoNumbers(l1, l2)
	fmt.Println(merge)
	fmt.Println(leetcode_tool.ShowList(merge))
}
