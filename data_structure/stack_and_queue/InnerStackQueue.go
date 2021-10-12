package stackandqueue

type Stack struct {
	innerSlice []interface{}
}

type InnerStackQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func NewStack() *Stack {
	return &Stack{
		innerSlice: make([]interface{}, 0),
	}
}

// 把slice的尾部假装成是栈的头部
func (s *Stack) push(v interface{}) {
	s.innerSlice = append(s.innerSlice, v)
}

func (s *Stack) pop() interface{} {
	// debug1
	if len(s.innerSlice) == 0 {
		return nil
	}
	v := s.innerSlice[len(s.innerSlice)-1]
	s.innerSlice = s.innerSlice[:len(s.innerSlice)-1]
	return v
}

// 题9：用两个栈实现一个队列，实现队列的两个函数appendTail和deleteHead
// 分别完成在尾部插入节点和在头部删除节点的功能
// 思路：用stack1保持入队顺序，stack2保持出队顺序&在stack1和stack2之间把元素倒来倒去，想删掉stack1最先入栈的元素，只需要把stack1的元素都pop出来再push到stack2，
//		就像倒水一样，stack2自然是stack1的倒序
// 边界条件：1. 往空队列里添加、删除元素 2. 连续删除元素直到队列为空
// 功能测试：往非空队列里添加、删除元素
// cost: 55min
// --------------------------------------------
// ---------------- debug 记录 ----------------
//  * 没有测试边界条件：从空队列里删除元素，通不过这个测试用例（创建queue -> deleteHead）
//  * 从stack2把队列头部拿出来之后，直接就return了，保存原来的入队顺序的stack1变成了空的，
// 		* 导致除非一直是deleteHead操作，否则addTail再deleteHead之后得到的不会是队列头部的
//		* 元素而是刚入队的元素
//		* 解决：从stack2 pop出来之后再塞回stack1
func (*InnerStackQueue) run() {
	q := &InnerStackQueue{
		stack1: NewStack(),
		stack2: NewStack(),
	}
	println(q.deleteHead())
	q.appendTail(1)
	q.appendTail(2)
	println(q.deleteHead().(int))
	q.appendTail(3)
	println(q.deleteHead().(int))
}

func (q *InnerStackQueue) appendTail(v interface{}) {
	q.stack1.push(v)
}

func (q *InnerStackQueue) deleteHead() interface{} {
	// debug1
	if len(q.stack1.innerSlice) == 0 {
		return nil
	}
	for len(q.stack1.innerSlice) > 0 {
		v := q.stack1.pop()
		q.stack2.push(v)
	}

	// debug2
	// return q.stack2.pop()
	data := q.stack2.pop()
	for len(q.stack2.innerSlice) > 0 {
		v := q.stack2.pop()
		q.stack1.push(v)
	}

	return data
}
