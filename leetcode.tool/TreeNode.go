/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/19
 * Time: 15:26
 */
package leetcode_tool

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 利用数组生成二叉树
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		// 每次都从新的根结点开始
		queue = queue[1:]

		if i < n {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// 先序遍历 根结点->左结点->右结点
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2PreOrder(root.Left)...)
	res = append(res, Tree2PreOrder(root.Right)...)

	return res
}

// 中序遍历 左结点->根结点->右结点
func Tree2InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2InOrder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2InOrder(root.Right)...)

	return res
}

// 后序遍历 左结点->右结点->根结点
func Tree2PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2PostOrder(root.Left)
	res = append(res, Tree2PostOrder(root.Right)...)
	res = append(res, root.Val)

	return res
}
