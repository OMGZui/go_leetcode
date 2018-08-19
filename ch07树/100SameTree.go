/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 20:40
 */
package main

import (
	"fmt"
	"leetcode.tool"
)

type TreeNode = leetcode_tool.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := leetcode_tool.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6})
	q := leetcode_tool.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(isSameTree(p, q))
	fmt.Println(leetcode_tool.Tree2PreOrder(p))
	fmt.Println(leetcode_tool.Tree2InOrder(p))
	fmt.Println(leetcode_tool.Tree2PostOrder(p))
}
