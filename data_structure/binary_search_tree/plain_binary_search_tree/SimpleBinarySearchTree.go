package main

import "fmt"

type BinaryTreeNode struct {
	m_nValue int
	m_pLeft  *BinaryTreeNode
	m_pRight *BinaryTreeNode
}

func main() {
	RebuildBinaryTree()
}

// 剑指offer 题7：根据前序遍历和中序遍历的结果重建二叉树
// !!!!前提：前序遍历和中序遍历的结果中不含重复的数字非常重要，为了找分割左右子树的点！！（见32行）（如果包含了重复的数字？）
// 根据前序遍历的第一个节点一定能在中序遍历中找到分开左右子树的点，即：得到左右子树的序列
// （因为只是同一颗树的不同遍历方式，推及子树也一样）
// 当某个子序列中只存在3个结点时，一定能推导出树，这颗树同时是上一个层级的左子树或者右子树；（！！！更正：不一定是三个，因为不一定是满二叉树）
// 分别对左右节点依次从下至上推导树，当某棵树的上一个根节点等于nil时，树构造完成
// 递归的方式：从根节点往下一直切分子序列，直到某个子序列能够推导出子树（如上所述）
// 		返回树的结构，和根节点，由于返回来的子树结构已经构建好了，所以只需要把子树的根节点
// 		挂到当前递归层级的树的相应位置（左子树or右子树）
// 考虑异常情况：1. 非平衡二叉树：左子树/右子树的层级比右子树/左子树深，导致len(左/右子树)在某一次递归及之后是0
//			2. 考虑数组越界： 3. 考虑访问nil值
// -------------------------------------------
// ------------- debug错误记录 ----------------
// 	1. 第二次递归时传进左子树的root是nil，而preOrder和midOrder都有值
//	   *看了一下题解，递归的func没传root，传root有必要吗？root的作用是在递归中挂上左右子树，不传能挂吗？
//	   *不传也是可以挂的，通过返回值层层传上去，这样的话还能用指针吗？岂不是会出现悬挂指针
//	   *不用指针定义不了结构体，报错：illegal cycle in definition struct
//  2. 递归边界没找全，不只是len(preOrder) == 1，还有len==0（如果倒数第二层的序列刚好是奇数的话）
func RebuildBinaryTree() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	midOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	binaryTree := RebuildFunc(preOrder, midOrder)
	printTreeHorizontal(binaryTree)
}

func RebuildFunc(preOrder []int, midOrder []int) (aboveRoot *BinaryTreeNode) {
	root := &BinaryTreeNode{}
	if len(preOrder) == 1 {
		root.m_nValue = preOrder[0]
		return root
	}
	if len(preOrder) == 0 {
		return nil
	}

	// 找出子树的前序遍历序列和中序遍历序列，给下层递归用
	currentRoot := preOrder[0]
	i := 0
	for ; i < len(midOrder); i++ {
		if currentRoot == midOrder[i] {
			break
		}
	}
	// 前序遍历序列根据已经找到的左子树的序列长度找就行了
	leftPreOrder := preOrder[1 : i+1]
	rightPreOrder := preOrder[i+1:]

	root.m_nValue = currentRoot
	root.m_pLeft = RebuildFunc(leftPreOrder, midOrder[0:i])
	root.m_pRight = RebuildFunc(rightPreOrder, midOrder[i+1:])

	return root
}

// 剑指offer 题32 从上到下依次打印每一层的节点（广度优先遍历）（层序遍历问题）
// 相似题：leetcode 102
// 通过一个list暂存每一层的节点&思路是只push子节点不push当前节点（避免重复）
// 开始遍历下一层时通过list里面暂存的节点继续遍历
// 先进先出，所以是个队列。
// ！！！为什么一定要暂存？？因为存在一个基础逻辑：程序是顺序执行的&树不能随机访问
//    因此，通过list暂存节点来达成回溯
func printTreeHorizontal(root *BinaryTreeNode) {
	stageNodesQueue := make([]*BinaryTreeNode, 0)
	stageNodesQueue = append(stageNodesQueue, root)
	// 结束条件：遍历完最后一层，即当前层的nodes没有一个子节点，即len(stageNodesQueue) == 0
	for {
		if len(stageNodesQueue) == 0 {
			break
		}
		node := stageNodesQueue[0]
		stageNodesQueue = stageNodesQueue[1:]
		fmt.Printf("%d  ", node.m_nValue)
		if node.m_pLeft != nil {
			stageNodesQueue = append(stageNodesQueue, node.m_pLeft)
		}
		if node.m_pRight != nil {
			stageNodesQueue = append(stageNodesQueue, node.m_pRight)
		}
	}
}