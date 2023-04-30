package main

import (
	"fmt"
	"guia2/queue"
)

func main() {
	var q queue.Queue[any]
	q.Enqueue(1)
	q.Enqueue("hola")
	q.Enqueue("Mundo")

	v, err := q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	v, err = q.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}

	_, err = q.Dequeue()
	fmt.Println(err)

}
