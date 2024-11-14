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
	sequenceInfix := "2 * ( 5 - 3 )"
	convertedToPostfix := converti(sequenceInfix)	
	fmt.Println("sequenceInfix in Postfixed:", convertedToPostfix)
	valuta(sequence)
}


func converti(seq string) string {
	s := stack.NewStack()
	result := ""
	for _, token := range  strings.Split(seq, " ") {
		digit := rune(token[0])
		if unicode.IsDigit(digit) {
			result += token + " "
		} else {
			switch token {
			case "(":
				break
			case ")":
				operand, _ := s.Pop()
				result += operand.(string) + " "
			default:  
				s.Push(token)
			}
		}
	}
	for s.Len() > 0 {
		operand, _ := s.Pop()
		result += operand.(string) + " "
	}
	return result
}

func valuta (seq string) int{
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
		res, _ := s.Pop()
		result := res.(int)
		fmt.Println("Postfixed Expression result:", result)
		return result
	} else {
		fmt.Println("Error: insufficient values in the expression")
		panic("Error: insufficient values in the expression")
	}
}