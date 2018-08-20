package stack

type pstack []int

func (s *pstack) push(v int) {
	*s = append(*s, v)
}

func (s *pstack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
