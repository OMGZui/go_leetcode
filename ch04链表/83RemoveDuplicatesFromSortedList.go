/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 10:56
 */
package main

import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteDuplicates(head *ListNode2) *ListNode2 {
	if head == nil {
		return nil
	}

	tail := head
	for tail.Next != nil {
		if tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		}else {
			tail = tail.Next
		}
	}
	return head
}

func main() {
	l := createList2([]int{1, 1, 2, 3, 3})
	fmt.Println(l)
	fmt.Println(showList2(l))

	head := deleteDuplicates(l)
	fmt.Println(head)
	fmt.Println(showList2(head))
}


func createList2(nums []int) *ListNode2 {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode2{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode2{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}

func showList2(head *ListNode2) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
