/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 15:38
 */
package leetcode_tool

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表并返回头指针
func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0], Next: nil}
	temp := head
	for i := 1; i < len(nums); i++ {
		//temp.Next = new(ListNode)
		//temp.Next.Val = nums[i]
		temp.Next = &ListNode{Val: nums[i], Next: nil}
		temp = temp.Next
	}

	return head
}

// 遍历链表
func ShowList(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
