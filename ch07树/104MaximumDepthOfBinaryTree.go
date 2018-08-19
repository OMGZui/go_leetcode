/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 16:25
 */
package main

import (
	"../leetcode.tool"
	"fmt"
)

type TreeNode3 = leetcode_tool.TreeNode

func maxDepth(root *TreeNode3) int {
	if root == nil {
		return 0
	}

	return 1 + leetcode_tool.Max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	q := leetcode_tool.Ints2TreeNode([]int{3, 9, 20, 0, 0, 15, 7})
	fmt.Println(maxDepth(q))
}
