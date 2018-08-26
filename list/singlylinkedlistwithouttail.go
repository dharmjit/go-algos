package list

type list struct {
	head *element
	size int
}

type element struct {
	value int
	next  *element
}

func (l *list) pushFront(v int) error {
	l.head = &element{v, l.head}
	l.size++
	return nil
}

func (l *list) popFront() int {
	e := l.head
	tmp := e
	e = nil
	l.head = tmp.next
	l.size--
	return tmp.value

}
func (l *list) pushBack(v int) error {
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &element{v, nil}
	l.size++
	return nil
}

func (l *list) popBack() int {
	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	last := tmp.value
	l.size--
	return last
}

func (l *list) len() int {
	return l.size
}

func (l *list) find(v int) bool {
	return false
}

func (l *list) remove(v int) bool {
	return false
}
