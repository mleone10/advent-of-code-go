package queue

type Queue[T any] struct {
	queue []T
}

func (q Queue[T]) Length() int {
	return len(q.queue)
}

func (q *Queue[T]) Push(r T) {
	q.queue = append(q.queue, r)
}

func (q *Queue[T]) Pop() T {
	r := q.queue[0]
	q.queue = q.queue[1:]
	return r
}
