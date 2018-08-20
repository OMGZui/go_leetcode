/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 16:58
 */
package main

import (
	"../../leetcode.tool"
	"fmt"
)

type TreeNode4 = leetcode_tool.TreeNode

func sortedArrayToBST(nums []int) *TreeNode4 {
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := n / 2
	return &TreeNode4{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func main() {
	q := []int{0, 0, 3, 7, 9, 15, 20}
	sortTree := sortedArrayToBST(q)
	fmt.Println(sortTree)
	fmt.Println(leetcode_tool.Tree2PreOrder(sortTree))
	fmt.Println(leetcode_tool.Tree2InOrder(sortTree))
}
