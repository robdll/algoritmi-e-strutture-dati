package main

import (
	"fmt"
	"stack_project/stack"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := stack.NewStack()
	sequence := "2 5 3 - *"
	CalcExpression(s, strings.Split(sequence, " "))
}


func CalcExpression (s *stack.Stack, seq []string) {
	for _, token := range  seq {
		fmt.Println("Token:", token)
		digit := rune(token[0])
		if unicode.IsDigit(digit) {
			num, _ := strconv.Atoi(token)
			s.Push(num)
			fmt.Println("Token added:", num)
		} else {
			value2, _ := s.Pop()
			value1, _ := s.Pop()
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