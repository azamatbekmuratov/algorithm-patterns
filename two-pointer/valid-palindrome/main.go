// Valid Palindrome solution in Golang language using two pointer technique
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(str string) bool {
	l, r := 0, len(str)-1
	for l < r {
		if !is_alphanum(string(str[l])) {
			l++
			continue
		} else if !is_alphanum(string(str[r])) {
			r--
			continue
		} else if !strings.EqualFold(string(str[l]), string(str[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func is_alphanum(word string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(word)
}

func main() {
	fmt.Println(isPalindrome("aka"))

}
