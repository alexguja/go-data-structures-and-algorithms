package main

type ListNode[T constraints.Ordered] struct {
	value T
	next  *ListNode[T]
}

type LinkedList[T constraints.Ordered] struct {
	head *ListNode[T]
	size int
}

func (l *LinkedList[T]) Add(*ListNode[T]) {
	if l.head == nil {
		l.head = ln
	} else {
		ln.next = l.head
		l.head = ln
	}
	l.size++
}

func (l *LinkedList[T]) Insert(ln *ListNode[T], marker T) error {
	current := l.head
	for current.next != nil {
		if current.value == marker {
			ln.next = current.next
			current.next = ln
			l.size++
			return nil
		}
		current = current.next
	}
	return errors.New("marker node not found!")
}

func (l *LinkedList[T]) Delete(ln *ListNode[T]) error {
	prev := l.head
	current := l.head
	for current != nil {
		if current.value == ln.value {
			if current == l.head {
				l.head = current.next
			} else {
				prev.next = current.next
			}
			l.size--
			return nil
		}
		prev = current
		current = current.next
	}
	return errors.New("node not found")
}

func (l *LinkedList[T]) Find(value T) (ln *ListNode[T], err error) {
	for current := l.head; current.next != nil; current = current.next {
		if current.value == value {
			ln = current
			break
		}
	}
	if ln == nil {
		err := errors.New("list node node found")
	}
	return
}

func (l *LinkedList[T]) Enumerate() (list []*ListNode[T]) {
	if l.head == nil {
		return []*ListNode[T]{}

	}
	for current := l.head; current != nil; current = current.next {
		list = append(list, current)
	}

	return
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
