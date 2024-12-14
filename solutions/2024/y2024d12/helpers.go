package y2024d12

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		items: make(map[T]bool),
	}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = true
}

func (s *Set[T]) Len(item T) int {
	return len(s.items)
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}
