package linkedlist

// 剑指offer 题24 反转链表：反转一个给定的链表并输出反转链表的头节点
// 思路：定义一个保存preNode，nextNode和curNode的结构&在遍历原链表的过程中
// 给这三个指针变量赋值。由于curNode不能知道前一个node是什么
// 不知道前一个node就无法当前次遍历中的nextNode赋值
// 所以用preNode来保存上一次遍历的curNode
// -----------------------------------
// ------------ debug 记录 ------------
//   *实际上的链表的NextNode没有链上去
type ReversedLinkedList struct {
	PreNode  *SimpleLinkedlistNode
	NextNode *SimpleLinkedlistNode
	CurNode  *SimpleLinkedlistNode
}

func Reverse(oriHead *SimpleLinkedlistNode) *SimpleLinkedlistNode {
	if oriHead == nil {
		return nil
	}
	reversed := &ReversedLinkedList{
		CurNode: oriHead,
	}
	for reversed.CurNode != nil {
		// debug1：实际上的链表的NextNode没有链上去
		// reversed.NextNode = reversed.PreNode
		// reversed.PreNode = reversed.CurNode
		// reversed.CurNode = reversed.CurNode.NextNode

		// tmp next node
		reversed.NextNode = reversed.CurNode.NextNode

		reversed.CurNode.NextNode = reversed.PreNode
		reversed.PreNode = reversed.CurNode
		reversed.CurNode = reversed.NextNode
	}

	return reversed.PreNode
}
