package tree

type node struct {
	value int
	left  *node
	right *node
}

type BSTree struct {
	root *node
}

func (tree *BSTree) Root() *node {
	return tree.root
}

func (tree *BSTree) Lookup(v int) *node {
	cur := tree.root
	for cur != nil {
		if v > cur.value {
			cur = cur.right
			continue
		}
		if v < cur.value {
			cur = cur.left
			continue
		}
		return cur
	}
	return nil
}

// Insert
func (tree *BSTree) Insert(v int) {
	new := node{value: v}
	cur := tree.root
	if tree.root == nil {
		tree.root = &new
	}
	for cur != nil {
		if v >= cur.value {
			// right
			if cur.right != nil {
				cur = cur.right
				continue
			}
			cur.right = &new
		} else {
			// left
			if cur.left != nil {
				cur = cur.left
				continue
			}
			cur.left = &new
		}
		break
	}
}

// Delete
func (tree *BSTree) Delete(v int) {
	cur := tree.root
	var prnt *node
	// Find Node To Delete
	for cur != nil {
		if v == cur.value {
			break
		}
		prnt = cur
		if v > cur.value {
			cur = cur.right
			continue
		}
		if v < cur.value {
			cur = cur.left
			continue
		}
	}
	// If removed node do not have right node
	if cur.right == nil {
		if prnt == nil {
			tree.root = cur.left
		}
		if prnt.right == cur {
			prnt.right = cur.left
		} else {
			prnt.left = cur.left
		}
		return
	}
	// If removed node has right node
	tarPrnt := cur
	tar := tarPrnt.right
	// Find most left node
	for tar.left != nil {
		tarPrnt = tar
		tar = tar.left
	}
	if prnt == nil {
		tree.root = tar
	}
	if prnt != nil && prnt.right == cur {
		prnt.right = tar
	}
	if prnt != nil && prnt.left == cur {
		prnt.left = tar
	}
	if tar != cur.right {
		tar.right = cur.right
	}
	tar.left = cur.left
	tarPrnt.left = nil

}

func (node *node) Left() *node {
	return node.left
}

func (node *node) Right() *node {
	return node.right
}

// Get Node Value
func (node *node) Value() interface{} {
	return node.value
}
