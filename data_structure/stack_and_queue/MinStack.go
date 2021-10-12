package stackandqueue

type MinStack struct {
	innerSlice []int
	assitStack *MinStack
}

// 题30：实现一个MinStack，包含push、pop、min方法且时间复杂度都是O(1)
// 思路：使用另一个辅助栈B&在往栈A中push的同时把最小数push到栈B
// 		min函数：最小数想要获得栈中最小的数，首先想到用一个变量暂存最小值
//		但如果这个最小值被弹出，要如何更新变量就是一个问题，除了再遍历一次栈别无他法，但这样就不能达到O(1)；
//		所以，不仅需要暂存当前栈中的最小值，还需要暂存次最小值、次次最小值…… => 每次push时都把当前最小值同时push到栈B
// 功能测试：1. 压入一串数字，不断弹出它们的最小值，直到栈为空（测试min函数无论执行多少次都能得到正确的结果）
//			2. push一个数字，再push一个比它大的数字，调用min；push一个数字，再push一个比它小的数字，调用min
//			3. 弹出的数字不是最小值：push 1、3、2，调用pop，再调用min
//			4. 弹出的数字是最小值：push 3、2、1，调用pop，再调用min
// 边界测试：1. nil：从空栈中弹出数字
// --------------------------------------
// ----------- debug记录 ----------------
// * 辅助栈也用了MinStack定义的方法，没有考虑辅助栈的assitStack我没有给它辅助，导致push和pop的时候操作assitStack的逻辑总是panic
func (m *MinStack) Run() {
	// 功能
	m.push(1)
	m.push(2)
	m.push(3)
	println(m.min())
	println(m.min())
	println(m.min())
	// nil测试
	m.pop()
	m.pop()
	m.pop()
	println(m.min())
	// 功能
	m.push(1)
	m.push(2)
	println(m.min())
	m.pop()
	m.pop()
	m.push(3)
	m.push(2)
	println(m.min())

	m.push(1)
	m.push(3)
	m.push(2)
	m.pop()
	println(m.min())

	m.push(3)
	m.push(2)
	m.push(1)
	m.pop()
	println(m.min())
}

func NewMinStack() *MinStack {
	return &MinStack{
		innerSlice: make([]int, 0),
		assitStack: &MinStack{
			innerSlice: make([]int, 0),
		},
	}
}

// 把slice的尾部假装成是栈的头部
func (s *MinStack) push(v int) {
	s.innerSlice = append(s.innerSlice, v)
	// 把push的值和辅助栈的栈顶比较
	if s.assitStack == nil {
		return
	}
	if len(s.assitStack.innerSlice) == 0 {
		s.assitStack.innerSlice = append(s.assitStack.innerSlice, v)
	} else {
		currentMin := s.assitStack.innerSlice[len(s.assitStack.innerSlice)-1]
		if v <= currentMin {
			s.assitStack.push(v)
		}
	}
}

func (s *MinStack) pop() int {
	if len(s.innerSlice) == 0 {
		return -1
	}
	v := s.innerSlice[len(s.innerSlice)-1]
	s.innerSlice = s.innerSlice[:len(s.innerSlice)-1]

	// 如果最小值弹出了，则将辅助栈的栈顶也弹出
	if s.assitStack != nil {
		cuurentMin := s.assitStack.innerSlice[len(s.assitStack.innerSlice)-1]
		if cuurentMin == v {
			s.assitStack.innerSlice = s.assitStack.innerSlice[:len(s.assitStack.innerSlice)-1]
		}
	}

	return v
}

func (s *MinStack) min() int {
	if len(s.assitStack.innerSlice) > 0 {
		return s.assitStack.innerSlice[len(s.assitStack.innerSlice)-1]
	} else {
		return -1
	}
}
