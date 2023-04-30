package queue

import (
	"errors"
	"guia2/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue[T any] struct {
	items []T
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue[T]) Enqueue(v T) {
	q.items = append(q.items, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue[T]) Dequeue() (T, error) {
	if (*q).IsEmpty() {
		var t T
		return t, errors.New("queue is empty")
	}

	head := (q.items)[0]
	q.items = (q.items)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue[T]) Front() (any, error) {
	if (*q).IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return (q.items)[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Queue implementa una cola genérica.
type QueueS[T any] struct {
	//implementación a partir de dos pilas (stacks).
	itemsHead stack.Stack[T]
	itemsTail stack.Stack[T]
}

// Encola un elemento v. O(1)
func (q *QueueS[T]) Enqueue(v T) {

	if q.itemsTail.IsEmpty() { // O(1)
		q.itemsTail.Push(v) // O(1)
	} else {
		q.itemsHead.Push(v) // O(1)
	}
}

// Desencola un elemento. // O(1)
func (q *QueueS[T]) Dequeue() (any, error) {
	if !q.itemsHead.IsEmpty() { // O(1)
		return q.itemsHead.Pop() // O(1)
	} else {
		return q.itemsTail.Pop() // O(1)
	}
}

// Devuelve si la cola está vacía. // O(1)
func (q *QueueS[T]) IsEmpty() bool {
	return q.itemsTail.IsEmpty() // O(1)
}

// Devuelve el primer elemento de la cola. // O(1)
func (q *QueueS[T]) Front() (T, error) {

	if q.itemsHead.IsEmpty() { // O(1)
		return q.itemsTail.Top() // O(1)
	} else {
		return q.itemsHead.Top() // O(1)
	}
}
