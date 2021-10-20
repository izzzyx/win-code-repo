package linkedlist

import "fmt"

// 剑指offer 题35 实现一个复制复杂链表的函数，复杂链表的定义如下``
// 思路：该题有两种解法
//  * 解法1：
//		** 第一步：复制原始链表（包括Val和SiblingNode）；同时建立一个“原链表节点地址”To“复制链表节点地址”的map
//		** 第二步：遍历一次复制链表，对每个node，从map中取出siblingNode（原链表地址）对应的复制链表地址，写回复制链表节点中
//	* 解法2：
//		** 与解法1相比无需额外的map，而是把复制链表node放在原始链表的node后面，
//		** 遍历时奇数=原始链表node，偶数=复制链表node，构思巧妙
// --------------------------------------
// -------------- debug记录 -------------
//	* 在第二步“遍历复制链表”的时候，没有移动curNode指向的位置 => 对链表遍历不熟悉
//  * 一开始对map的作用没想清楚，存的是原始链表的地址，后又改成new一个对象（相当于哪个链表都不在的地址）
//  	** 一直报“siblingNode的地址不在复制链表中”的错误。其实复制siblingNode的问题就在于
//  	** 第一次遍历的时候不知道siblingNode在复制链表中的地址，因此只要给两个链表的节点地址做一个映射就好了。
type ComplexLinkedListNode struct {
	Val         int
	NextNode    *ComplexLinkedListNode
	SiblingNode *ComplexLinkedListNode
}

func ExecCopyComplexLinkedList1() {
	node2 := &ComplexLinkedListNode{
		Val: 2,
	}
	node1 := &ComplexLinkedListNode{
		Val:         1,
		NextNode:    node2,
		SiblingNode: node2,
	}
	head := &ComplexLinkedListNode{
		Val:         0,
		NextNode:    node1,
		SiblingNode: nil,
	}
	node2.SiblingNode = head

	CopyComplexLinkedList1(head)
}

// 解法1
func CopyComplexLinkedList1(oriHead *ComplexLinkedListNode) (copyedHead *ComplexLinkedListNode) {
	if oriHead == nil {
		return nil
	}

	// 原链表节点和复制链表节点地址的映射
	// 先暂存siblingNode指向原链表节点地址，再从map里面拿到对应的复制链表节点的地址
	oriNodeToCopyNodeMap := make(map[*ComplexLinkedListNode]*ComplexLinkedListNode)

	curNode := oriHead
	copyNode := &ComplexLinkedListNode{
		Val:      oriHead.Val,
		NextNode: &ComplexLinkedListNode{},
	}
	copyNode.SiblingNode = curNode.SiblingNode
	copyedHead = copyNode

	oriNodeToCopyNodeMap[curNode] = copyNode

	for curNode.NextNode != nil {
		curNode = curNode.NextNode
		copyNode = copyNode.NextNode
		copyNode.Val = curNode.Val
		copyNode.SiblingNode = curNode.SiblingNode
		copyNode.NextNode = &ComplexLinkedListNode{}

		oriNodeToCopyNodeMap[curNode] = copyNode
	}
	copyNode.NextNode = nil

	// 遍历copyNode，再链siblingNode
	copyNode = copyedHead
	for copyNode != nil {
		if copyNode.SiblingNode != nil {
			copyNode.SiblingNode = oriNodeToCopyNodeMap[copyNode.SiblingNode]
		}
		copyNode = copyNode.NextNode
	}

	fmt.Printf("%+v", copyedHead)
	return copyedHead
}
