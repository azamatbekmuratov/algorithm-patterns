package main

import "fmt"

type void struct{}

var member void

func singleNumber(nums []int) int {
	set := make(map[int]void)
	set[2] = member

	for k := range set {
		fmt.Println(k)
	}
}

func main() {

}
