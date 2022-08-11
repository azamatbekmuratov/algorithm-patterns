// Daily Temperatures. Solution in Go language
// Given an array of integers temperatures represents the daily temperatures,
// return an array answer such that answer[i] is the number of days you have to wait after
// the ith day to get a warmer temperature. If there is no future day for which this is possible,
// keep answer[i] == 0 instead.
package main

import "fmt"

type Stack []int

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

// Returns the value of peek element in stack
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	} else {
		return (*s)[len(*s)-1]
	}
}

func dailyTemperatures(t []int) []int {
	res := []int{}
	n := len(t)

	var daysOfWait []int = make([]int, n)

	s := Stack{}

	for i := 0; i < n; i++ {
		// Check if current index is the
		// next warmer temperature of
		// any previous indexes
		for !s.IsEmpty() && t[s.Peek()] < t[i] {
			daysOfWait[s.Peek()] = i - s.Peek()

			// Pop element
			s.Pop()
		}

		// Push the current index
		s.Push(i)
	}

	// Print waiting days
	for i := 0; i < n; i++ {
		fmt.Print(daysOfWait[i], " ")
		res = append(res, daysOfWait[i])
	}
	return res
}

func main() {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}

	dailyTemperatures(arr)
}
