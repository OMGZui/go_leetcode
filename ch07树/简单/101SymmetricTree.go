/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 15:19
 */
package main

import (
	"fmt"
	"../../leetcode.tool"
)

type TreeNode2 = leetcode_tool.TreeNode

func isSymmetric(root *TreeNode2) bool {
	if root == nil {
		return true
	}

	return recur(root.Left, root.Right)

}

func recur(left *TreeNode2, right *TreeNode2) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && recur(left.Left, right.Right) && recur(left.Right, right.Left)
}

func main() {
	q := leetcode_tool.CreateTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(isSymmetric(q))
}
