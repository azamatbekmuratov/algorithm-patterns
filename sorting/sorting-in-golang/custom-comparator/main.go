package main

import (
	"fmt"
	"sort"
)

func main() {
	family := []struct {
		Name string
		Age  int
	}{
		{"Alice", 23},
		{"David", 2},
		{"Eve", 2},
		{"Bob", 25},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[i].Age
	})
	fmt.Println(family)
}
