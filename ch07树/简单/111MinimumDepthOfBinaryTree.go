/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/20
 * Time: 14:37
 */
package 简单

import (
	"../../leetcode.tool"
	"fmt"
)


type TreeNode6 = leetcode_tool.TreeNode

func minDepth(root *TreeNode6) int {
	return leetcode_tool.MinDepth(root)
}

func main() {
	var null = 1000
	q := leetcode_tool.CreateTreeNode([]int{3, 9, 20, null, null, 15, 7})
	q2 := leetcode_tool.CreateTreeNode([]int{1, 2})
	fmt.Println(minDepth(q))
	fmt.Println(minDepth(q2))
}
