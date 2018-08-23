package tree

type node struct {
	Value int
	Left  *node
	Right *node
}

func (n *node) insert(value int) error {
	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &node{Value: value}
			return nil
		}
		return n.Left.insert(value)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &node{Value: value}
			return nil
		}
		return n.Right.insert(value)
	}
	return nil
}

func (n *node) find(v int) bool {
	if n == nil {
		return false
	}
	switch {
	case v == n.Value:
		return true
	case v < n.Value:
		return n.Left.find(v)
	default:
		return n.Right.find(v)
	}
}

func (n *node) replaceNode(parent *node, replacementNode *node) error {
	if n == parent.Left {
		parent.Left = replacementNode
		return nil
	}
	parent.Right = replacementNode
	return nil
}

func (n *node) findMax(parent *node) (*node, *node) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *node) delete(v int, parent *node) error {
	switch {
	case v < n.Value:
		return n.Left.delete(v, n)
	case v > n.Value:
		return n.Right.delete(v, n)
	default:
		if n.Left == nil && n.Right == nil {
			n.replaceNode(parent, nil)
			return nil
		}
		if n.Left == nil {
			n.replaceNode(parent, n.Right)
			return nil
		}
		if n.Right == nil {
			n.replaceNode(parent, n.Left)
		}
		replacement, replParent := n.Left.findMax(n)
		n.Value = replacement.Value
		return replacement.delete(replacement.Value, replParent)
	}
}
