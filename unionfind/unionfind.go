package unionfind

// Node of a set
type Node struct {
	data   int
	parent *Node
}

// UnionFind DataStructure
type UnionFind struct {
	Set []Node
}

// Make taken a slice of int representing value of nodes
// and return a slice of single node disjoint sets
func Make(nodes []int) *UnionFind {
	uf := UnionFind{make([]Node, len(nodes))}
	for i, n := range nodes {
		uf.Set[i] = Node{data: n}
		uf.Set[i].parent = &uf.Set[i]
	}
	return &uf
}

// String to implement stringer interface
func (uf *UnionFind) String() string {
	//TODO
	return ""
}

// Find will find the representative node of the set or will return -1
func (uf *UnionFind) Find(node int) int {
	for _, u := range uf.Set {
		for u.parent != nil {
			if u.data == node {
				return u.data
			}
		}
	}
	return -1
}
