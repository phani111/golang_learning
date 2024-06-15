package list

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Len  int
}

func New[T any]() *List[T] {
	return &List[T]{
		Head: nil,
		Len:  0,
	}
}

func (l *List[T]) Add(value T) {
	newNode := &Node[T]{Value: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Len++
}
