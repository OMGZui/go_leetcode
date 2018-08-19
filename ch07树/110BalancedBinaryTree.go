/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 17:47
 */
package main

import (
	"../leetcode.tool"
	"fmt"
)

type TreeNode5 = leetcode_tool.TreeNode

func isBalanced(root *TreeNode5) bool {
	if root == nil {
		return true
	}
	// 左右结点深度相差不能大于1
	leftLen := leetcode_tool.MaxDepth(root.Left)
	rightLen := leetcode_tool.MaxDepth(root.Right)
	if leftLen-rightLen > -2 && leftLen-rightLen < 2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func main() {
	q := leetcode_tool.Ints2TreeNode([]int{1, 2, 2, 3, 3, 0, 0, 4, 4})
	p := leetcode_tool.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(isBalanced(q))
	fmt.Println(isBalanced(p))
}
