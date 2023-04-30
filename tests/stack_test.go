package tests

import (
	"guia2/stack"
	"testing"
)

func TestPush(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, _ := s.Top()

	if v != 3 {
		t.Error("Error en Push")
	}
}

func TestPop(t *testing.T) {
	var s stack.Stack[int]

	s.Push(1)
	s.Push(2)
	s.Push(3)

	v, err := s.Pop()

	if err != nil || v != 3 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 2 {
		t.Error("Error en Pop")
	}

	v, err = s.Pop()

	if err != nil || v != 1 {
		t.Error("Error en Pop")
	}

	_, err = s.Pop()

	if err == nil {
		t.Error("Error en Pop")
	}
}
