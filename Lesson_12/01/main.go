package main

import (
	"01/stack"
	"fmt"
)

func main() {
	s := stack.NewStack()

	s.Push(10)
	s.Push(20)
	s.Push(30)

	top, _ := s.Peek()
	fmt.Println("Top element:", top) // Should print "Top element: 30"

	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println("Popped:", val)
	}
}
