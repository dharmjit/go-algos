package list

type list struct {
	head *element
}

type element struct {
	value int
	next  *element
}

func (l *list) pushFront(v int) error {
	l.head = &element{v, l.head}
	return nil
}

func (l *list) popFront() int {
	e := l.head
	tmp := e
	e = nil
	l.head = tmp.next
	return tmp.value

}
func (l *list) pushBack(v int) error {
	return nil
}

func (l *list) popBack() int {
	return 0
}

func (l *list) size() int {
	return 0
}

func (l *list) find(v int) bool {
	return false
}

func (l *list) remove(v int) bool {
	return false
}
