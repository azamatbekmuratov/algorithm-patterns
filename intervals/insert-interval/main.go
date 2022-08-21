// Insert Interval. Solution in Golang language
package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	for i := range intervals {
		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			newInterval = []int{min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])}
		}
		fmt.Println(newInterval)
	}
	res = append(res, newInterval)
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
}
