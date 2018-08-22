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
