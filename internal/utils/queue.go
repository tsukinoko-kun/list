package utils

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Push(data T) {
	q.data = append(q.data, data)
}

func (q *Queue[T]) Pop() *T {
	if len(q.data) == 0 {
		return nil
	}
	data := q.data[0]
	q.data = q.data[1:]
	return &data
}

func (q *Queue[T]) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.data)
}

func (q *Queue[T]) Peek() *T {
	return &q.data[0]
}
