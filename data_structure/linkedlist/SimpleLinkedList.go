package linkedlist

import (
	stackandqueue "data_structure/stack_and_queue"
	"fmt"
)

type SimpleLinkedlistNode struct {
	Val      int
	NextNode *SimpleLinkedlistNode
}

type SimpleLinkedlist struct {
	TailNode *SimpleLinkedlistNode
	HeadNode *SimpleLinkedlistNode
}

func (l *SimpleLinkedlist) Add(val int) {
	if l.HeadNode == nil {
		return
	}
	node := &SimpleLinkedlistNode{
		Val: val,
		NextNode: new(SimpleLinkedlistNode),
	}
	// 用tailNode保存最后一个node的地址
	if l.HeadNode.NextNode == nil {
		l.HeadNode.NextNode = node
		l.TailNode = l.HeadNode
	}
	curTail := l.TailNode
	if curTail != nil {
		curTail.NextNode = node
		l.TailNode = node
	} else {
		fmt.Println("tailNode is nil, ")
	}
}

func NewSimpleLinkedList(val int) *SimpleLinkedlist{
	return &SimpleLinkedlist{
		HeadNode: &SimpleLinkedlistNode{
			Val: val,
			NextNode: new(SimpleLinkedlistNode),
		},
	}
}

func (l *SimpleLinkedlist) PrintFromTail() {
	stack := stackandqueue.NewStack()
	iterNode := l.HeadNode
	for iterNode != nil {
		stack.Push(iterNode.Val)
		iterNode = iterNode.NextNode
	}
	for i := 0; i < stack.Len(); i++ {
		fmt.Printf("%d ", stack.Pop())
	}
}