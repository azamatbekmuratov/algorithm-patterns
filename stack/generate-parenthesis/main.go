package main

import "fmt"

func generateParenthesis(n int) []string {
	// only add open parenthesis if open < n
	// only add closing parenthesis if closed < open
	// valid open == closed == n
	if n == 0 {
		return []string{}
	}
	res := []string{}
	backtrack(n, n, "", &res)
	return res
}

func backtrack(openN int, closedN int, str string, res *[]string) {
	if openN == 0 && closedN == 0 {
		*res = append(*res, str)
		return
	}
	if openN > 0 {
		backtrack(openN-1, closedN, str+"(", res)
	}
	if closedN > 0 && openN < closedN {
		backtrack(openN, closedN-1, str+")", res)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
}
