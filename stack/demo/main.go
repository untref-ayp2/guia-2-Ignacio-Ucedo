package main

import (
	"fmt"
	"guia2/stack"
)

func main() {
	var s stack.Stack[string]
	cadena := "Hola Mundo"
	for _, v := range cadena {
		s.Push(string(v))
	}

	_, err := s.Pop()
	for err == nil {
		_, err = s.Pop()
		fmt.Printf("%v\n", err)
	}

}
