package list

type dlist struct {
	head *delement
	size int
}

type delement struct {
	value    int
	next     *delement
	previous *delement
}

func (d *dlist) pushFront(v int) error {
	if d.size == 0 {
		d.head = &delement{v, nil, nil}
		return nil
	}
	newelement := &delement{v, d.head, nil}
	d.head.previous = newelement
	d.head = newelement
	d.size++
	return nil
}

func (d *dlist) popFront() int {
	if d.size == 0 {
		return -1
	}
	popElement := d.head.value
	d.head = d.head.next
	d.size--
	return popElement
}

func (d *dlist) pushBack(v int) error {
	if d.size == 0 {
		d.head = &delement{v, nil, nil}
		return nil
	}
	tmp := d.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &delement{v, tmp.previous, nil}
	d.size++
	return nil

}

func (d *dlist) popBack() int {
	if d.size == 0 {
		return -1
	}
	tmp := d.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	d.size--
	return tmp.value
}
