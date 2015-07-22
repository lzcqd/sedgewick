package symboltable

type StringToIntST interface {
	Put(key string, val int)
	Get(key string) (int, bool)
	Delete(key string)
}

type UnorderedLinkedList struct {
	key   string
	value int
	next  *UnorderedLinkedList
}

func (l UnorderedLinkedList) Put(key string, val int) {
	p := &l
	for p != nil {
		p = (*p).next
	}

	(*p).next = &UnorderedLinkedList{key, val, nil}
}

func (l UnorderedLinkedList) Get(key string) (int, bool) {
	p := &l
	for p != nil {
		if (*p).key == key {
			return (*p).value, true
		}
		p = (*p).next
	}
	return 0, false
}

func (l *UnorderedLinkedList) Delete(key string) {
	if (*l).key == key {
		l = l.next
		return
	}

	prev, curr := l, (*l).next
	for curr != nil {
		if (*curr).key == key {
			(*prev).next = (*curr).next
			return
		}
		prev, curr = curr, (*curr).next
	}
}
