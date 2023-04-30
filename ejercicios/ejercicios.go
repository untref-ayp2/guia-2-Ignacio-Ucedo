package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

// recibe una cadena de caracteres y devuelve la cadena invertida. O(n)
func InvertirCadena(s string) string {
	var pila stack.Stack[string]

	for i := 0; i < len(s); i++ { // O(len(s)) = O(n)
		pila.Push(string([]rune(s)[i])) //voy poniendo caracter por caracter en la pila
	}

	s = "" //vacio el string

	for !pila.IsEmpty() { // O(len(s)) = O(n)
		char, _ := pila.Top()
		pila.Pop()
		s = s[:] + char //lo voy llenando //suposición: O(1)
	}

	//esta función es 2O(n), o sea, O(n)

	return s
}

// devuelve si un string es palíndromo (es decir que una cadena es igual a su inversa). O(n)
func Palindromo(s string) bool {
	return s == InvertirCadena(s)
	// lo único que hace es llamar a un O(n), por lo que es O(n)
}

// devuelve si una cadena de paréntesis, corchetes y llaves está bien balanceada. O(n)
func Balanceada(s string) bool {

	var balanceada bool = true

	var pila stack.Stack[string]
	var parentesis queue.Queue[string]
	var corchetes queue.Queue[string]
	var llaves queue.Queue[string]

	s = InvertirCadena(s) // O(n)

	for i := 0; i < len(s); i++ { // O(n)
		pila.Push(string([]rune(s)[i])) //voy poniendo caracter por caracter en la pila
	}

	for !pila.IsEmpty() && balanceada { // O(n) (todo lo de adentro del while es O(1))
		top, _ := pila.Top()
		pila.Pop()
		if top == "{" || top == "}" || top == "[" || top == "]" || top == "(" || top == ")" {
			if top == "{" || top == "}" {
				if (top == "}" && (llaves.IsEmpty() || !corchetes.IsEmpty())) || (top == "{" && !llaves.IsEmpty()) {
					balanceada = false
				} else {
					if llaves.IsEmpty() {
						llaves.Enqueue(top)
					} else {
						llaves.Enqueue(top)
						llaves.Dequeue()
						llaves.Dequeue()
					}
				}
			} else if top == "[" || top == "]" {
				if (top == "]" && (corchetes.IsEmpty() || !parentesis.IsEmpty())) || (top == "[" && !corchetes.IsEmpty()) {
					balanceada = false
				} else {
					if corchetes.IsEmpty() {
						corchetes.Enqueue(top)
					} else {
						corchetes.Enqueue(top)
						corchetes.Dequeue()
						corchetes.Dequeue()
					}
				}
			} else {
				if (top == ")" && parentesis.IsEmpty()) || (top == "(" && !parentesis.IsEmpty()) {
					balanceada = false
				} else {
					if parentesis.IsEmpty() {
						parentesis.Enqueue(top)
					} else {
						parentesis.Enqueue(top)
						parentesis.Dequeue()
						parentesis.Dequeue()
					}
				}
			}
		} else {
			balanceada = false
		}
	}

	if !llaves.IsEmpty() || !corchetes.IsEmpty() {
		balanceada = false
	}

	return balanceada

	//la función termina siendo 3o(len(s)), o sea O(n)
}

// Une dos colas. O(n)
func UnirColas(q1, q2 queue.Queue[any]) queue.Queue[any] {
	q3 := q1

	for !q2.IsEmpty() { // O(largo de q2) => O(n)
		front, _ := q2.Front()
		q2.Dequeue()
		q3.Enqueue(front)
	}

	return q3
}
