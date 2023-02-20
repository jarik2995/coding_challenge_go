package queue

type node struct {
	value interface{}
	next  *node
}

type Queue struct {
	first  *node
	last   *node
	length int
}

// enqueue
func (queue *Queue) Enqueue(v interface{}) {
	new := node{value: v}
	if queue.length == 0 {
		queue.first = &new
		queue.last = &new
		queue.length++
		return
	}
	queue.last.next = &new
	queue.last = &new
	queue.length++
}

// dequeue
func (queue *Queue) Dequeue() (v interface{}) {
	if queue.length == 0 {
		return nil
	}
	tmp := queue.first.value
	if queue.length == 1 {
		queue.first = nil
		queue.last = nil
		queue.length--
		return tmp
	}
	queue.first = queue.first.next
	queue.length--
	return tmp
}

// peek
func (queue *Queue) Peek() interface{} {
	return queue.first.value
}
