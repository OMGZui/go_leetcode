/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/29
 * Time: 19:36
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type ListNode = leetcode_tool.ListNode

type TreeNode = leetcode_tool.TreeNode

func sortedChild(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	if head.Next == tail {
		//如果只剩这一个元素了
		return &TreeNode{Val: head.Val}
	}

	mid, tmp := head, head
	//一个前进一个位置,一个前进2个位置,最后mid会到中间
	for tmp != tail && tmp.Next != tail {
		mid = mid.Next
		tmp = tmp.Next.Next
	}

	return &TreeNode{
		Val:   mid.Val,
		Left:  sortedChild(head, mid),
		Right: sortedChild(mid.Next, tail),
	}
}

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedChild(head, nil)
}

func main() {
	q := []int{-10, -3, 0, 5, 9}
	head := leetcode_tool.CreateList(q)
	sortTree := sortedListToBST(head)
	fmt.Println(sortTree)
	fmt.Println(leetcode_tool.Tree2PreOrder(sortTree))
	fmt.Println(leetcode_tool.Tree2InOrder(sortTree))
}
