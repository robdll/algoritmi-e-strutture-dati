package main

import (
	"fmt"
	"stack_project/stack"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	sequence := "2 5 3 - *"
	showExpression(sequence)	
	valuta(sequence)
}


func showExpression(seq string) {
	s := stack.NewStack()
	for _, token := range  strings.Split(seq, " ") {
		digit := rune(token[0])
		if unicode.IsDigit(digit) {
			num, _ := strconv.Atoi(token)
			s.Push(num)
		} else {
			value2, _ := s.Pop()
			value1, _ := s.Pop()
			s.Push(value1)
			s.Push(value2)
		}
	}
	fmt.Println("Postfixed Expression:", s)
}

func valuta (seq string) {
	s := stack.NewStack()
	for _, token := range  strings.Split(seq, " ") {
		fmt.Println("Token:", token)
		digit := rune(token[0])
		if unicode.IsDigit(digit) {
			num, _ := strconv.Atoi(token)
			s.Push(num)
			fmt.Println("Token added:", num)
		} else {
			val2, _ := s.Pop()
			val1, _ := s.Pop()
			value2 := val2.(int)
			value1 := val1.(int)
			switch token {
			case "+":
				fmt.Println("Adding:", value1, value2)
				s.Push(value1 + value2)
			case "-":
				fmt.Println("Subtracting:", value1, value2)
				s.Push(value1 - value2)
			case "*":
				fmt.Println("Multiplying:", value1, value2)
				s.Push(value1 * value2)
			case "/":
				fmt.Println("Dividing:", value1, value2)
				s.Push(value1 / value2)
			}
			newValue, _ := s.Peek()
			fmt.Println("Added value to stack:", newValue)
		}
	}

	if s.Len() == 1 {
		result, _ := s.Pop()
		fmt.Println("Postfixed Expression result:", result)
	} else {
		fmt.Println("Error in expression")
	}
}