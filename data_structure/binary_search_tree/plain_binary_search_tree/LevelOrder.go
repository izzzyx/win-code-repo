package plain_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指offer 题32-1 从上到下打印二叉树
// 解题思路：回溯&二叉树可以访问左右子树&先进先出（队列）
// 测试：
//	* 功能测试：满二叉树打印
//  * 特殊输入：
//		* 0值：完全二叉树/左子树nil/右子树nil
//		* 1值：只有根节点
func LevelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	prints := make([]int, 0)
	stageNodeList := make([]*TreeNode, 0)

	var current = root
	prints = append(prints, current.Val)
	for current != nil {
		if current.Left != nil {
			stageNodeList = append(stageNodeList, current.Left)
			prints = append(prints, current.Left.Val)
		}
		if current.Right != nil {
			stageNodeList = append(stageNodeList, current.Right)
			prints = append(prints, current.Right.Val)
		}
		if len(stageNodeList) > 0 {
			current = stageNodeList[0]
			stageNodeList = stageNodeList[1:]
		} else {
			current = nil
		}
	}

	return prints
}
