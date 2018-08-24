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

func CreateList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{Val: nums[0], Next: nil}
	temp := res
	for i := 1; i < len(nums); i++ {
		//temp.Next = new(ListNode)
		//temp.Next.Val = nums[i]
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return res
}

func ShowList(head *ListNode) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
