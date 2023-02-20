package stack

type node struct {
	value interface{}
	next  *node
}

type Stack struct {
	top    *node
	bottom *node
	length int
}

// push
func (stack *Stack) Push(v interface{}) {
	new := node{value: v}
	if stack.length == 0 {
		stack.top = &new
		stack.bottom = &new
		stack.length++
		return
	}
	new.next = stack.top
	stack.top = &new
	stack.length++
}

// pop
func (stack *Stack) Pop() interface{} {
	if stack.length == 0 {
		return nil
	}
	tmp := stack.top.value
	if stack.length == 1 {
		stack.top = nil
		stack.bottom = nil
		stack.length--
		return tmp
	}
	stack.top = stack.top.next
	stack.length--
	return tmp
}

// peek
func (stack *Stack) Peek() interface{} {
	return stack.top.value
}

// Queue Using Stacks
type QueueUsingStack struct {
	queue Stack
}

func (queue *QueueUsingStack) Enqueue(v interface{}) {
	tmp_stack := Stack{}
	for i := 0; i < queue.queue.length; {
		last := queue.queue.Pop()
		tmp_stack.Push(last)
	}
	tmp_stack.Push(v)
	queue.queue = Stack{}
	for i := 0; i < tmp_stack.length; {
		last := tmp_stack.Pop()
		queue.queue.Push(last)
	}
}

func (queue *QueueUsingStack) Dequeue() (v interface{}) {
	return queue.queue.Pop()
}

func (queue *QueueUsingStack) Peek() (v interface{}) {
	return queue.queue.Peek()
}
