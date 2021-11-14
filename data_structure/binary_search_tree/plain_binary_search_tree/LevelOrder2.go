package plain_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 剑指offer 题32-2 从上到下打印二叉树，每一层打印到一行
// 关键字tag：二叉树；回溯
// 思路：用两个标量表示还没有遍历的节点数（作为终止标记）和下一层的节点数（作为终止标记的暂存）
//	* 功能测试：满二叉树打印
//  * 特殊输入：
//		* 0值：完全二叉树/左子树nil/右子树nil
//		* 1值：只有根节点
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{
			{
				root.Val,
			},
		}
	}

	prints := make([][]int, 0)
	stageNodeList := make([]*TreeNode, 0)

	var current = root
	prints = append(prints, []int{current.Val})
	n := 1
	m := 0
	levelPrints := make([]int, 0)
	for current != nil {
		if current.Left != nil {
			stageNodeList = append(stageNodeList, current.Left)
			levelPrints = append(levelPrints, current.Left.Val)
			m++
		}
		if current.Right != nil {
			stageNodeList = append(stageNodeList, current.Right)
			levelPrints = append(levelPrints, current.Right.Val)
			m++
		}
		n--
		if n == 0 {
			prints = append(prints, levelPrints)
			n = m
			m = 0
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
