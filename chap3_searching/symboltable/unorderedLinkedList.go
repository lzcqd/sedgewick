package symboltable

type Element struct {
	key   string
	value int

	next *Element
}

type List struct {
	root Element
	len  int
}

func (l *List) Init() *List {
	l.root.next = nil
	l.len = 0
	return l
}

func (l *List) Put(key string, val int) {
	if l.len == 0 {
		l.root.next = &Element{key, val, nil}
		l.len++
		return
	}

	e := &l.root
	for e.next != nil {
		if e.key == key {
			e.value = val
			return
		}
		e = e.next
	}

	e.next = &Element{key, val, nil}
	l.len++
}

func (l *List) Get(key string) (int, bool) {
	e := l.root.next
	for e != nil {
		if e.key == key {
			return e.value, true
		}
		e = e.next
	}
	return 0, false
}

func (l *List) Delete(key string) {
	if l.len == 0 {
		return
	}

	prev, curr := &l.root, l.root.next

	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			l.len--
			return
		}
		prev, curr = curr, curr.next
	}
}
