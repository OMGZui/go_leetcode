/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/8/16
 * Time: 20:40
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
	p := Ints2TreeNode([]int{1, 2, 3, 4, 5})
	q := Ints2TreeNode([]int{1, 2, 3, 4, 50})
	fmt.Println(isSameTree(p, q))
	fmt.Println(Tree2PreOrder(p))
	fmt.Println(Tree2InOrder(p))
	fmt.Println(Tree2PostOrder(p))
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

