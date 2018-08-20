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

// 测试数据
var null = 1000

// 利用数组生成二叉树
func CreateTreeNode(num []int) *TreeNode {
	n := len(num)
	if n == 0 {
		return nil
	}

	// 根结点
	root := &TreeNode{Val: num[0]}
	// 结点容器
	tree := make([]*TreeNode, 1, n*2)
	// 将root存入
	tree[0] = root

	i := 1
	for i < n {
		// 构造根结点
		node := tree[0]
		// 每次都从新的根结点开始
		tree = tree[1:]

		// 构造左结点
		if i < n && num[i] != null {
			node.Left = &TreeNode{Val: num[i]}
			tree = append(tree, node.Left)
		}
		i++
		// 构造右结点
		if i < n && num[i] != null {
			node.Right = &TreeNode{Val: num[i]}
			tree = append(tree, node.Right)
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

// 树的最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + Max(MaxDepth(root.Left), MaxDepth(root.Right))
}

// 树的最小深度
func MinDepth(root *TreeNode) int {
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + MinDepth(root.Right)
	case root.Right == nil:
		return 1 + MinDepth(root.Left)
	default:
		return 1 + Min(MinDepth(root.Left), MinDepth(root.Right))
	}
}
