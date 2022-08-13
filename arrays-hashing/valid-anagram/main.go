// Valid Anagram solution in golang
// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS := make(map[byte]int)
	countT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, found := countS[s[i]]; found {
			countS[s[i]] = 1 + countS[s[i]]
		} else {
			countS[s[i]] = 0
		}
		if _, found := countT[t[i]]; found {
			countT[t[i]] = 1 + countT[t[i]]
		} else {
			countT[t[i]] = 0
		}
	}

	for k, v := range countS {
		if w, ok := countT[k]; !ok || v != w {
			return false
		}
	}

	return true

	// or can use reflect package and use DeepEqual function
	// return reflect.DeepEqual(countS, countT)

}

func main() {
	fmt.Println(isAnagram("anagram", "naagram"))
}
