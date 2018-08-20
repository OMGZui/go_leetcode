/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 16:25
 */
package 简单

import (
	"../../leetcode.tool"
	"fmt"
)

type TreeNode3 = leetcode_tool.TreeNode

func maxDepth(root *TreeNode3) int {
	return leetcode_tool.MaxDepth(root)
}

func main() {
	var null = 1000
	q := leetcode_tool.CreateTreeNode([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(maxDepth(q))
}
