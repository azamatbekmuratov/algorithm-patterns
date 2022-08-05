// Evaluate Reverse Polish Notation
// Evaluate the value of an arithmetic expression in Reverse Polish Notation.
// Valid operators are +, -, *, and /. Each operand may be an integer or another expression.
package main

import "strconv"

type Stack []int

// IsEmpty: check if is stack empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func evalRPN(tokens []string) int {
	stack := Stack{}

	for _, c := range tokens {
		if c == "+" {
			stack.Push(stack.Pop() + stack.Pop())
		}
		if c == "-" {
			a, b := stack.Pop(), stack.Pop()
			stack.Push(b - a)
		}
		if c == "*" {
			stack.Push(stack.Pop() * stack.Pop())
		}
		if c == "/" {
			a, b := stack.Pop(), stack.Pop()
			stack.Push(b / a)
		} else {
			digit, err := strconv.Atoi(c)
			if err != nil {
				stack.Push(digit)
			}
		}
	}
	return stack[0]
}
