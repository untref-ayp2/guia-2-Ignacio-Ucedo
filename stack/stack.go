package stack

import "errors"

// Stack implementa una pila genérica sobre un arreglo dinámico.
type Stack[T any] struct {
	items []T
}

// Push agrega un elemento a la pila. O(1)
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop saca un elemento de la pila. O(1)
func (s *Stack[T]) Pop() (any, error) {
	if (*s).IsEmpty() {
		return nil, errors.New("la pila esta vacia")
	}
	v := (s.items)[len(s.items)-1]
	s.items = (s.items)[:len(s.items)-1] //acorta s.itmes en sacándole el último valor
	return v, nil
}

// Top devuelve el elemento del tope de la pila. O(1)
func (s *Stack[T]) Top() (T, error) {
	if (*s).IsEmpty() {
		var t T
		return t, errors.New("la pila esta vacia")
	}
	v := (s.items)[len(s.items)-1] //guardo el ultimo valor de s.items en v
	return v, nil
}

// IsEmpty verifica si la pila esta vacia. O(1)
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
