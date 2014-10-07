package linkedlist

type List struct {
	head   *Node
	tail   *Node
	Length int
}

type Node struct {
	value float64
	next  *Node
	prev  *Node
}

func (n *Node) Del() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

func New() *List {
	return &List{}
}

func (l *List) Add(value float64) bool {
	n := &Node{value: value}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.Length = 1
		return false
	} else {
		current := l.head
		for ; current.next != nil; current = current.next {
			if current.next.value > n.value {
				break
			} else if current.next.value == n.value {
				return true
			}
		}
		n.next = current.next
		if current.next != nil {
			current.next.prev = n
		} else {
			// this would mean we're at the end of the list
			l.tail = n
		}
		current.next = n
		n.prev = current
		l.Length += 1
	}
	return false
}

func (l *List) GetIndex(value float64) int {
	current := l.head
	if current == nil {
		return -1
	}
	for index := 0; current != nil; index += 1 {
		if current.value == value {
			return index
		}
		current = current.next
	}
	return -1
}

func (l *List) GetAtIndex(i int) (float64, bool) {
	current := l.head
	if current == nil {
		return 0, false
	}
	for index := 0; current != nil; index += 1 {
		if index == i {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func (l *List) Del(value float64) bool {
	if l.head == nil {
		return false
	}
	deleted := false
	current := l.head
	for ; current.next != nil; current = current.next {
		if current.value == value {
			current.Del()
			deleted = true
			l.Length -= 1
		} else if current.value > value {
			return false
		}
	}
	return deleted
}

func (l *List) DelAtIndex(i int) bool {
	current := l.head
	if current == nil {
		return false
	}
	for index := 0; current != nil; index += 1 {
		if index == i {
			current.Del()
			l.Length -= 1
			return true
		}
		current = current.next
	}
	return false
}

func (l *List) GetAll() []float64 {
	results := []float64{}
	current := l.head
	if current == nil {
		return nil
	}
	for ; current != nil; current = current.next {
		results = append(results, current.value)
	}
	return results
}

func (l *List) GetAllRev() []float64 {
	results := []float64{}
	current := l.tail
	if current == nil {
		return nil
	}
	for ; current != nil; current = current.prev {
		results = append(results, current.value)
	}
	return results
}
