package linked_list

type node struct {
	Value interface{}
	Next  *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

// Append
func (ll *LinkedList) Append(v interface{}) {
	new := node{Value: v}
	if ll.Head == nil {
		ll.Head = &new
	} else {
		ll.Tail.Next = &new
	}
	ll.Tail = &new
	ll.Length++
}

// Prepend
func (ll *LinkedList) Prepend(v interface{}) {
	new := node{Value: v}
	new.Next = ll.Head
	ll.Head = &new
	ll.Length++
}

// Insert
func (ll *LinkedList) InsertAfterNode(v interface{}, n *node) {
	new := node{Value: v}
	new.Next = n.Next
	n.Next = &new
	ll.Length++
}

// Delete
func (ll *LinkedList) DeleteNode(n *node) {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		if cur.Next == n {
			cur.Next = n.Next
			ll.Length--
			break
		}
	}
}

// Search
func (ll *LinkedList) GetNodeByValue(v interface{}) (n *node) {
	for cur := ll.Head; cur != nil; cur = cur.Next {
		if cur.Value == v {
			n = cur
		}
	}
	return
}

func (ll *LinkedList) GetNodeByIndex(index int) (n *node) {
	if index > ll.Length {
		return
	}
	cur := ll.Head
	for i := 0; i <= index; i++ {
		if i == index {
			n = cur
			break
		}
		cur = cur.Next
	}
	return
}

// Reverse
func (ll *LinkedList) Reverse() {
	if ll.Head == ll.Tail {
		return
	}
	cur := ll.Head
	next := cur.Next
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	ll.Head, ll.Tail = ll.Tail, ll.Head
	ll.Tail.Next = nil
}
